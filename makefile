.PHONE: build bin

build:
	go build -o hnh main.go
	chmod +x hnh

bin: build
	sudo cp hnh /usr/local/bin
