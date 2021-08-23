build-tgbot-for-pi:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/tgbot github.com/nazarov-pro/nazarovsh/cmd/tgbot

build-localtunnel-for-pi:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/localtunnel github.com/nazarov-pro/nazarovsh/cmd/localtunnel

send-pi-via-scp:
	scp -r bin/ pi:/home/pi/bin

build-and-send-pi: build-tgbot-for-pi build-localtunnel-for-pi send-pi-via-scp
