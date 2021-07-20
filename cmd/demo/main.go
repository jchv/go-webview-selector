package main

import (
	"log"

	"github.com/jchv/go-webview-selector"
)

func main() {
	debug := true
	w := webview.New(debug)
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintFixed)
	w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	w.Run()
}
