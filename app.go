package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Timer and Stopwatch")

	myWindow.SetFixedSize(true)

	timerLabel := widget.NewLabel("Timer: 0:00")
	timerEntry := widget.NewEntry()
	timerEntry.SetPlaceHolder("Enter seconds")
	startTimerButton := widget.NewButton("Start Timer", nil)
	resetTimerButton := widget.NewButton("Reset Timer", nil)

	stopwatchLabel := widget.NewLabel("Stopwatch: 0:00")
	startStopwatchButton := widget.NewButton("Start", nil)
	pauseStopwatchButton := widget.NewButton("Pause", nil)
	resetStopwatchButton := widget.NewButton("Reset", nil)

	var timerDuration int
	var timerTicker *time.Ticker
	var timerRunning bool
	startTimerButton.OnTapped = func() {
		if timerRunning {
			return
		}
		timerDuration, _ = strconv.Atoi(timerEntry.Text)
		timerTicker = time.NewTicker(1 * time.Second)
		timerRunning = true
		go func() {
			for timerRunning {
				time.Sleep(1 * time.Second)
				timerDuration--
				if timerDuration <= 0 {
					timerLabel.SetText("Timer: Time's up!")
					timerTicker.Stop()
					timerRunning = false
				} else {
					minutes := timerDuration / 60
					seconds := timerDuration % 60
					timerLabel.SetText(fmt.Sprintf("Timer: %d:%02d", minutes, seconds))
				}
			}
		}()
	}
	resetTimerButton.OnTapped = func() {
		if timerTicker != nil {
			timerTicker.Stop()
		}
		timerRunning = false
		timerLabel.SetText("Timer: 0:00")
		timerEntry.SetText("")
	}

	var stopwatchTicker *time.Ticker
	var stopwatchDuration int
	var stopwatchRunning bool
	startStopwatchButton.OnTapped = func() {
		if stopwatchRunning {
			return
		}
		stopwatchRunning = true
		stopwatchTicker = time.NewTicker(1 * time.Second)
		go func() {
			for stopwatchRunning {
				time.Sleep(1 * time.Second)
				stopwatchDuration++
				minutes := stopwatchDuration / 60
				seconds := stopwatchDuration % 60
				stopwatchLabel.SetText(fmt.Sprintf("Stopwatch: %d:%02d", minutes, seconds))
			}
		}()
	}
	pauseStopwatchButton.OnTapped = func() {
		if stopwatchTicker != nil {
			stopwatchTicker.Stop()
		}
		stopwatchRunning = false
	}
	resetStopwatchButton.OnTapped = func() {
		if stopwatchTicker != nil {
			stopwatchTicker.Stop()
		}
		stopwatchRunning = false
		stopwatchDuration = 0
		stopwatchLabel.SetText("Stopwatch: 0:00")
	}

	content := container.NewVBox(
		widget.NewLabel("Timer"),
		timerEntry,
		timerLabel,
		container.NewHBox(startTimerButton, resetTimerButton),
		widget.NewSeparator(),
		widget.NewLabel("Stopwatch"),
		stopwatchLabel,
		container.NewCenter(
			container.NewHBox(
				startStopwatchButton,
				pauseStopwatchButton,
				resetStopwatchButton,
			),
		),
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
