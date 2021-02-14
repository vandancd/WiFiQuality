#!/bin/bash
echo "`date`: Building wifiquality . . . "
go mod download
GOMAXPROCS=1 go build -o wifiquality 2>err.log
if [ -s "err.log" ];then
  echo -e "`date`: Build failed, check err.log "
  exit 2
fi
mkdir -p WiFiQuality.app/Contents/MacOS
cp wifiquality WiFiQuality.app/Contents/MacOS

mkdir -p WiFiQuality.app/Contents/MacOS/assets
cp assets/wifiquality.ico WiFiQuality.app/Contents/MacOS/assets

mkdir -p WiFiQuality.app/Contents/Resources
cp assets/icon.icns WiFiQuality.app/Contents/Resources
cp assets/WiFiQuality.png WiFiQuality.app/Contents/Resources

cat << EOF > WiFiQuality.app/Contents/Info.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
        <key>NSHighResolutionCapable</key>
        <string>True</string>
        <!-- avoid showing the app on the Dock -->
        <key>LSUIElement</key>
        <string>1</string>
</dict>
</plist>
EOF
echo -e "`date`: App build successfully." 