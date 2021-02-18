package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/getlantern/systray"

	"github.com/vandancd/wifiquality/util"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	// Sets the icon in the systray. This matches the App Icon.
	systray.SetIcon(util.WiFi)

	// Show Signal Strength and Noise as menu items
	mStrength := systray.AddMenuItem("Strength:", "")
	mNoise := systray.AddMenuItem("Noise:", "")

	// Since these are display only, disabling click targets
	mStrength.Disable()
	mNoise.Disable()

	systray.AddSeparator()

	go func() {
		for {
			w := util.NewWireless()
			systray.SetTitle("SNR: " + strconv.Itoa(w.Quality))
			systray.SetTooltip("Strength: " + strconv.Itoa(w.Strength) + " | Noise: " + strconv.Itoa(w.Noise))
			mStrength.SetTitle("Strength: " + strconv.Itoa(w.Strength) + " dBm")
			mNoise.SetTitle("Noise: " + strconv.Itoa(w.Noise) + " dBm")
			time.Sleep(15 * time.Second)
		}
	}()
	mQuit := systray.AddMenuItem("Quit", "Quit WiFi Quality")
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

}
func onExit() {
	// nothing much to do here
}

func getAppIcon(ico string) []byte {
	appIcon, err := ioutil.ReadFile(ico)
	if err != nil {
		fmt.Print(err)
	}
	return appIcon
}
