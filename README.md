# Go Webview Selector
This library is a tiny shim that provides access to a native webview for a given platform. It selects between two implementations:

* On Windows, [go-webview2](https://github.com/jchv/go-webview2) will be chosen. It embeds a copy of WebView2Loader, so it should work out of the box on any setup that has the WebView 2 runtime installed, which should include up-to-date Windows 10 installations and Windows 11. (As well, WebView 2 runtime is available going backwards to Windows 7.) This library does not require CGo and thus will build and run with `CGO_ENABLED=0`.

* On other operating systems, [webview/webview](https://github.com/webview/webview) will be used. This runtime requires CGo, but supports a wide variety of platforms.
