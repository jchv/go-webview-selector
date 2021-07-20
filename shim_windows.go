// +build windows

package webview

import (
	"unsafe"

	"github.com/jchv/go-webview2"
)

type WebView = webview2.WebView

type Hint = webview2.Hint

const (
	// HintNone specifies that width and height are default size
	HintNone Hint = iota

	// HintFixed specifies that window size can not be changed by a user
	HintFixed

	// HintMin specifies that width and height are minimum bounds
	HintMin

	// HintMax specifies that width and height are maximum bounds
	HintMax
)

// New creates a new webview in a new window.
func New(debug bool) WebView {
	return webview2.New(debug)
}

// NewWindow creates a new webview using an existing window.
func NewWindow(debug bool, window unsafe.Pointer) WebView {
	return webview2.NewWindow(debug, window)
}
