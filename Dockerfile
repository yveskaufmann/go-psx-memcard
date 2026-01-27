FROM golang:1.24 AS builder

WORKDIR /app

# Enable multi-arch support
RUN dpkg --add-architecture arm64

# Add GoReleaser PPA
RUN echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | tee /etc/apt/sources.list.d/goreleaser.list

# Install all dependencies and LLVM-MinGW for Windows ARM64 cross-compilation
RUN apt-get update && apt-get install --no-install-recommends -y \
   wget \
   xz-utils \
   build-essential \
   pkg-config \
   yq \
   # Cross-compilation toolchains \
   mingw-w64 \
   gcc-aarch64-linux-gnu \
   libc6-dev-arm64-cross \
   # Native amd64 libraries \
   libgl1-mesa-dev \
   libwayland-dev \
   libx11-dev \
   libxkbcommon-dev \
   libxrandr-dev \
   libxcursor-dev \
   libxi-dev \
   libxinerama-dev \
   libxxf86vm-dev \
   xorg-dev \
   xvfb \
   # ARM64 cross-arch libraries \
   libgl1-mesa-dev:arm64 \
   libwayland-dev:arm64 \
   libx11-dev:arm64 \
   libxrandr-dev:arm64 \
   libxcursor-dev:arm64 \
   libxi-dev:arm64 \
   libxinerama-dev:arm64 \
   libxxf86vm-dev:arm64 \
   libxkbcommon-dev:arm64 \
   goreleaser \
   && wget -q https://github.com/mstorsjo/llvm-mingw/releases/download/20240619/llvm-mingw-20240619-ucrt-ubuntu-20.04-x86_64.tar.xz \
   && tar -xf llvm-mingw-20240619-ucrt-ubuntu-20.04-x86_64.tar.xz -C /opt \
   && rm llvm-mingw-20240619-ucrt-ubuntu-20.04-x86_64.tar.xz \
   && ln -s /opt/llvm-mingw-20240619-ucrt-ubuntu-20.04-x86_64 /opt/llvm-mingw

# Install yq v4.47.1
RUN wget https://github.com/mikefarah/yq/releases/download/v4.47.1/yq_linux_amd64.tar.gz -O - |\
  tar xz && mv yq_linux_amd64 /usr/bin/yq


ENV PATH="/opt/llvm-mingw/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Enable CGO for Fyne/OpenGL builds
ENV CGO_ENABLED=1

# Set default build target to linux/amd64
ARG GOOS=linux
ARG GOARCH=amd64

# Disable nfpms for non-linux builds - It works around of the missing npmfs "if" which is exclusive to the pro version.
# See: https://goreleaser.com/customization/nfpm/
RUN set -ex && \
    if [ "${GOOS}" != "linux" ]; then \
        yq 'del(.nfpms)' .goreleaser.yml > .goreleaser.yml.tmp && \
        mv .goreleaser.yml.tmp .goreleaser.yml; \
    fi && \
    GOOS=${GOOS} GOARCH=${GOARCH} \
    yq '.builds[0].goos = [env(GOOS)] , .builds[0].goarch = [env(GOARCH)]' \
       .goreleaser.yml > .goreleaser.yml.tmp && \
    mv .goreleaser.yml.tmp .goreleaser.yml

# Run GoReleaser to build artifacts
RUN goreleaser release --snapshot --clean --config .goreleaser.yml

FROM scratch
COPY --from=builder /app/dist /
