package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/getlantern/systray"

	"github.com/vandancd/wifiquality/util"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "/tmp/wifiquality.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
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
			systray.SetTitle("SNR: " + strconv.Itoa(w.Quality) + " dB")
			if w.Quality >= 40 {
				systray.SetIcon(util.WiFiGreen)
			} else if w.Quality <= 15 {
				systray.SetIcon(util.WiFiRed)
			} else {
				systray.SetIcon(util.WiFiOrange)
			}
			systray.SetTooltip("Strength: " + strconv.Itoa(w.Strength) + " | Noise: " + strconv.Itoa(w.Noise))
			mStrength.SetTitle("Strength: " + strconv.Itoa(w.Strength) + " dBm")
			mNoise.SetTitle("Noise: " + strconv.Itoa(w.Noise) + " dBm")
			log.Printf("Network: %s; BSSID: %s; SNR: %v; Strength: %v; Noise:%v", w.NetworkName, w.BSSID, w.Quality, w.Strength, w.Noise)
			time.Sleep(60 * time.Second)
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
