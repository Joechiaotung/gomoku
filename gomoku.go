
package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"github.com/skiptomyliu/gomoku/view"

)


var port int


func processFlags() error {
	// General flags
	flag.IntVar(&port, "port", 1234, "Port to start the UI web server on; valid range: 0..65535")
	// flag.BoolVar(&autoOpen, "autoOpen", true, "Auto-opens the UI web page in the default browser")

	// View package flags
	flag.IntVar(&view.ViewWidth, "viewWidth", 700, "width of the view image in pixels in the UI web page; valid range: 150..2000")
	flag.IntVar(&view.ViewHeight, "viewHeight", 700, "height of the view image in pixels in the UI web page; valid range: 150..2000")

	flag.Parse()

	if port < 0 || port > 65535 {
		return fmt.Errorf("port %d is outside of valid range", port)
	}

	if view.ViewWidth < 150 || view.ViewWidth > 2000 {
		return fmt.Errorf("viewWidth %d is outside of valid range", view.ViewWidth)
	}

	if view.ViewHeight < 150 || view.ViewHeight > 2000 {
		return fmt.Errorf("viewHeight %d is outside of valid range", view.ViewHeight)
	}

	return nil
}

func main() {

	if err := processFlags(); err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}

	port := 3000

	// ctrl.StartEngine()
	fmt.Printf("Starting GoLab webserver on port %d...\n", port)
	url := fmt.Sprintf("http://localhost:%d/", port)

	fmt.Printf("Opening %s...\n", url)
	if err := open(url); err != nil {
		fmt.Println("Auto-open failed:", err)
		fmt.Printf("Open %s in your browser.\n", url)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}


func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)

	return exec.Command(cmd, args...).Start()
}