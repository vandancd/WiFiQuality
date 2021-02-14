build:
	go build -o wifiquality
	go build -o macmaker github.com/vandancd/wifiquality/macos-app
run:
	./macmaker -bin wifiquality -name "WiFiQuality" -icon macos-app/WiFiQuality.png -identifier com.vandan.wifi
