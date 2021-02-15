# WiFiQuality

macOS has a utility to check your wifi connectivity - `airport`.  Running this will show important metrics that you need to understand the quality of your WiFi Network:
```
$> /System/Library/PrivateFrameworks/Apple*.framework/Versions/Current/Resources/airport -I
     agrCtlRSSI: -40
     agrExtRSSI: 0
    agrCtlNoise: -93
    agrExtNoise: 0
          state: running
        op mode: station
     lastTxRate: 234
        maxRate: 867
lastAssocStatus: 0
    802.11 auth: open
      link auth: wpa2-psk
          BSSID: 1x:1x:1:1x:11:1
           SSID: XXXXXXXXXX
            MCS: 5
        channel: 48,80
```

Two numbers are most important here. `agrCtlRSSI` (Received Signal Strength Indicator) is the power of the received signal in the wireless network. It uses a logarithmic scale expressed in decibels (db) and typically ranges from 0 to -100. The close this number is to 0 the better quality of signal. 

The second is Noise or `agrCtlNoise`; is the impact of unwanted interfering signal sources, such as distortion and radio frequency interference. This is also measured in decibels (db) from 0 to -120. The lower the value i.e closer to  -120 means little to no noise in the wireless network. 

Once you have these two values, you can now measure the Signal to Noise Margin (SNR Margin) with the simple formula `agrCtlRSSI - agrCtlNoise`.

## About WiFiQuality
`WiFiQuality` will run showing SNR Margin updating every 15 seconds. You can change this in `main.go` 
`time.Sleep(15 * time.Second)`

TODO: Add a config file / preferences.

The `Makefile build` has two options:
1. Buld the WiFiQuality 
2. Build a macOS app maker

You can run `make build` to create these two items.

The `Makefile run` builds the `macOS` app. It uses [mholt/macapp](https://gist.github.com/mholt/11008646c95d787c30806d3f24b2c844) that allows me to build a macOS application.

## Icon for WiFiQuality 
Right now it uses the 👌🏽 emoji. 
TODO: Use an Icon Data file to generate a relevant icon to show up in the system tray. 

![WiFiQuality Systray](https://optimisticallyskeptical.files.wordpress.com/2021/02/wifiqulity.png?w=349)
