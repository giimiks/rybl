OUT_DIR := ./releases
APP := rybl

.PHONY: all build clean run \
        build-linux build-mac build-windows \
        build-arm build-arm64

all: build-linux build-mac build-windows build-arm build-arm64

build:
	go build -o $(OUT_DIR)/$(APP).exe

run: build
	$(OUT_DIR)/$(APP).exe

clean:
	rm -rf $(OUT_DIR)

$(OUT_DIR):
	mkdir -p $(OUT_DIR)

build-windows: $(OUT_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/$(APP)-win64.exe
	GOOS=windows GOARCH=386 go build -o $(OUT_DIR)/$(APP)-win32.exe

build-linux: $(OUT_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/$(APP)-linux64
	GOOS=linux GOARCH=386 go build -o $(OUT_DIR)/$(APP)-linux32

build-mac: $(OUT_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/$(APP)-mac64
	GOOS=darwin GOARCH=arm64 go build -o $(OUT_DIR)/$(APP)-mac-arm64

build-arm: $(OUT_DIR)
	GOOS=linux GOARCH=arm GOARM=6 go build -o $(OUT_DIR)/$(APP)-armv6
	GOOS=linux GOARCH=arm GOARM=7 go build -o $(OUT_DIR)/$(APP)-armv7

build-arm64: $(OUT_DIR)
	GOOS=linux GOARCH=arm64 go build -o $(OUT_DIR)/$(APP)-arm64
