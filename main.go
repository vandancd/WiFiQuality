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
	//systray.SetIcon(getAppIcon("assets/icon.icns"))
	go func() {
		w := util.NewWireless()
		for {
			systray.SetTitle("👌🏽 Quality: " + strconv.Itoa(w.Quality))
			systray.SetTooltip("Strength: " + strconv.Itoa(w.Strength) + " | Noise: " + strconv.Itoa(w.Noise))
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
