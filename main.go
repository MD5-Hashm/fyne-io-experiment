package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type players struct {
	Now int `json:"now"`
}

type mcApi struct {
	Status string `json:"status"`
	Online bool `json:"online"`
	Motd string `json:"motd"`
	Playercount players `json:"players"`
}

func getResp(url string) mcApi {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response mcApi
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response
}

func getPlayers(ip string) string {
	resp := getResp("https://mcapi.us/server/status?ip="+ip)
	players := strconv.Itoa(resp.Playercount.Now)
	return players
}

func newLabel(i string) fyne.CanvasObject {
	return widget.NewLabel(i)
}

func main() {
	myApp := app.New()
	win := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input.SetPlaceHolder("Server ip")

	win.Resize(fyne.NewSize(500,250))
	content := container.NewVBox(input, widget.NewButton("Search", func() {
		win.SetContent(container.NewVBox(newLabel("Players online:"),newLabel(getPlayers(input.Text))))
		win.Resize(fyne.NewSize(222,150))
	}))
	win.SetContent(content)
	win.ShowAndRun()
}
