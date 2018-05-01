package view

import (
	"fmt"
	"github.com/skiptomyliu/gomoku/model"
	"html/template"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"
	"time"
)

var Params = struct {
	Title         string
	Width, Height *int
	RunId         int64
	ShowFreezeBtn bool
}{AppTitle, &ViewWidth, &ViewHeight, time.Now().Unix(), false}

// Template of the play html page
var playTempl = template.Must(template.New("t").Parse(play_html))

// The client's (browser's) view position
// var Pos image.Point

// init registers the http handlers.
func init() {
	http.HandleFunc("/", playHtmlHandle)
	http.HandleFunc("/runid", runIdHandle)
	http.HandleFunc("/img", imgHandle)
	http.HandleFunc("/clicked", clickedHandle)
	http.HandleFunc("/new", newGameHandle)
	// http.HandleFunc("/help", helpHtmlHandle)
}

// InitNew initializes a new game.
func InitNew() {
	// Pos = image.Point{}
}

// playHtmlHandle serves the html page where the user can play.
func playHtmlHandle(w http.ResponseWriter, r *http.Request) {
	playTempl.Execute(w, Params)
}

// runidHandle serves the running app id which changes if app is restarted
// (so browser clients can detect if app was restarted).
func runIdHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "runid: %d", Params.RunId)
}

// imgHandle serves images of the player's view.

var quality int // is this right?
func imgHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in imghandle")
	fmt.Printf("w: %v h: %v", ViewWidth, ViewHeight)
	quality = 100

	rect := image.Rect(0, 0, ViewWidth, ViewHeight).Add(image.Pt(10, 10))
	model.Mutex.Lock()
	jpeg.Encode(w, model.LabImg.SubImage(rect), &jpeg.Options{quality})
	model.Mutex.Unlock()
	// Store the new view's position:
	// Pos = rect.Min
}

// clickedHandle receives mouse click (mouse button pressed) events with mouse coordinates.
func clickedHandle(w http.ResponseWriter, r *http.Request) {

	x, err := strconv.Atoi(r.FormValue("x"))
	if err != nil {
		return
	}

	y, err := strconv.Atoi(r.FormValue("y"))
	if err != nil {
		return
	}

	fmt.Println("Clicked Handle")
	fmt.Printf("Clicked: %s, %s", x, y)

	// x, y are in the coordinate system of the client's view.
	// Translate them to the Labyrinth's coordinate system:
	// select {
	// case model.ClickCh <- model.Click{Pos.X + x, Pos.Y + y, btn}:
	// default:
	// }
}


// // newGameHandle signals to start a newgame.
func newGameHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new game handle")
	// Use non-blocking send
	// select {
	// case model.NewGameCh <- 1:
	// default:
	// }
}