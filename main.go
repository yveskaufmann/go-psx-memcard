package main

import "com.yv35.memcard/pkg/ui"

func main() {
	if err := ui.Start(); err != nil {
		panic(err)
	}
}
