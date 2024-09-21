package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// var (
// 	outputDevice *portaudio.DeviceInfo
// 	inputDevice  *portaudio.DeviceInfo
// )

// func init() {
// 	err := portaudio.Initialize()
// 	if err != nil {
// 		panic(err)
// 		// log.Fatalf("Unable to initialize portaudio: %v\n", err)
// 	}
// 	defer portaudio.Terminate()

// 	outputDevice, err = portaudio.DefaultOutputDevice()
// 	if err != nil {
// 		panic(err)
// 		// log.Fatalf("Failed to get devices: %v\n", err)
// 	}
// 	inputDevice, err = portaudio.DefaultInputDevice()
// 	if err != nil {
// 		panic(err)
// 		// log.Fatalf("Failed to get devices: %v\n", err)
// 	}

// }

func wirteText(textChan chan string) {
	textChan <- "Hello, World!"
	for {
		textChan <- fmt.Sprintf("now is: %v", time.Now())
		time.Sleep(1 * time.Second)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("TMTChat")

	// set height and width of the window
	w.Resize(fyne.NewSize(600, 500))

	// create a channel to receive text
	textChan := make(chan string)
	defer close(textChan)

	go wirteText(textChan)

	// save history of the text

	history := widget.NewMultiLineEntry()
	// history.Resize(fyne.NewSize(400, 100))

	infoLabel := widget.NewLabel("info")
	// infoLabel.Resize(fyne.NewSize(400, 0))

	history.Wrapping = fyne.TextWrapWord
	history.SetMinRowsVisible(20)

	w.SetContent(container.NewVBox(history, infoLabel))

	go func() {
		for text := range textChan {
			infoLabel.SetText(text)
			history.SetText(history.Text + "\n" + text)
			lines := len(history.Text)
			history.CursorRow = lines - 1
			history.Refresh()
		}
	}()

	// go func() {
	// 	textChan <- fmt.Sprintf("%v\n", outputDevice.Name)
	// 	textChan <- fmt.Sprintf("%v\n", inputDevice.Name)
	// }()

	w.ShowAndRun()
}
