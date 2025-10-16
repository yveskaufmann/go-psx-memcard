package main

import "com.yvka.memcard/pkg/ui"

func main() {
	if err := ui.Start(); err != nil {
		panic(err)
	}
}
