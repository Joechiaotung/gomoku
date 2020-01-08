
package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"os/exec"
	"os"
	"runtime"
	"./ctrl"
	"./model"
	"./view"
)

var port string

func processFlags() error {
	// General flags
	// flag.BoolVar(&autoOpen, "autoOpen", true, "Auto-opens the UI web page in the default browser")

	flag.IntVar(&model.Rows, "rows", 21, "the number of rows on Board; must be odd; valid range: 9..99")
	flag.IntVar(&model.Cols, "cols", 21, "the number of columns on Board; must be odd; valid range: 9..99")


	// View package flags
	flag.IntVar(&view.ViewWidth, "viewWidth", 670, "width of the view image in pixels in the UI web page; valid range: 150..2000")
	flag.IntVar(&view.ViewHeight, "viewHeight", 670, "height of the view image in pixels in the UI web page; valid range: 150..2000")

	flag.Parse()


	model.BoardWidth = model.Cols * model.BlockSize
	model.BoardHeight = model.Rows * model.BlockSize


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

	ctrl.StartEngine()

	port := "3000"
	fmt.Printf("Starting GoLab webserver on port %d...\n", port)
	url := fmt.Sprintf("http://localhost:%d/", port)

	fmt.Printf("Opening %s...\n", url)
	if err := open(url); err != nil {
		fmt.Println("Auto-open failed:", err)
		fmt.Printf("Open %s in your browser.\n", url)
	}

	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
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