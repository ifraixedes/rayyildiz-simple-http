package main

import (
	"fmt"

	"github.com/rayyildiz/simple-http/icons"
	"runtime"
	tr "github.com/cratonica/trayhost"
)

func main() {
	// EnterLoop must be called on the OS's main thread
	runtime.LockOSThread()

	go func() {
		// Run your application/server code in here. Most likely you will
		// want to start an HTTP server that the user can hit with a browser
		// by clicking the tray icon.

		tr.SetUrl("http://localhost")

		// Be sure to call this to link the tray icon to the target url
		tr.SetUrl("http://localhost:8080")
	}()

	// Enter the host system's event loop
	tr.EnterLoop("Simple Http", icons.IconData)

	// This is only reached once the user chooses the Exit menu item
	fmt.Println("Exiting")
}
