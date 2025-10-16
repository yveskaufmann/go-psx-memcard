package main

import "com.yv35.memcard/internal/ui"

func main() {
	if err := ui.Start(); err != nil {
		panic(err)
	}
}
