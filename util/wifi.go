package util

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const (
	osxCmd  = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
	osxArgs = "-I"
)

//Wireless Object on macOS
type Wireless struct {
	NetworkName string `json:"networkName"`
	Strength    int    `json:"strength"`
	Noise       int    `json:"noise"`
	Quality     int    `json:"quality"`
}

//NewWireless - gets the name of the wifi
func NewWireless() *Wireless {
	w := new(Wireless)

	out, err := exec.Command(osxCmd, osxArgs).Output()
	if err != nil {
		//return nil, err
		fmt.Println(err)
	}

	// TODO: Figure out  better way to get this RegEx wokring in the for loop
	regName := regexp.MustCompile(`s*SSID: (.+)s*`)
	tmpName := regName.FindAllStringSubmatch(string(out), -1)
	if len(tmpName) > 1 {
		w.NetworkName = tmpName[1][1]
	}

	wifiData := strings.Split(string(out), "\n")
	for _, con := range wifiData {
		if len(con) == 0 {
			continue
		}
		if (w.Strength >= 0) || (w.Noise >= 0) {
			// Get Wireless Strength data
			regStrength := regexp.MustCompile("agrCtlRSSI: (-?\\d+)")
			tmpStrength := regStrength.FindStringSubmatch(con)
			if len(tmpStrength) > 1 {
				if tmp, err := strconv.Atoi(tmpStrength[1]); err == nil {
					w.Strength = tmp
				}
			}
			// Get Wireless Noise data
			regNoise := regexp.MustCompile("agrCtlNoise: (-?\\d+)")
			tmpNoise := regNoise.FindStringSubmatch(con)
			if len(tmpNoise) > 1 {
				if tmp, err := strconv.Atoi(tmpNoise[1]); err == nil {
					w.Noise = tmp
				}
			}
		} else {
			// No need to parse further as we have got our Singal Strength & Noise values
			break
		}

	}
	w.Quality = w.Strength - w.Noise
	return w
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
