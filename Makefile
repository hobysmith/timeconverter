BINARY_NAME=timeconverter
SHORT_DATE := $(shell date +%Y-%m-%d\ %H:%M:%S)
LONG_DATE := $(shell date)
FLAGS := -X 'github.com/hobysmith/timeconverter/cmd.AppShortBuildTime=${SHORT_DATE}' -X 'github.com/hobysmith/timeconverter/cmd.AppLongBuildTime=${LONG_DATE}'
.DEFAULT_GOAL := build

build:
	$(info Building default platform target...)
	go build -o ${BINARY_NAME} -ldflags="${FLAGS}"

install: build
	@echo
	@echo Installing Timeconverter...
	go install

mac-arm64:
	$(info Building Mac ARM64 target...)
	GOOS=darwin GOARCH=arm64 go build -o ${BINARY_NAME}-mac-arm64 -ldflags="${FLAGS}"

mac-amd64:
	$(info Building Mac AMD64 target...)
	GOOS=darwin GOARCH=amd64 go build -o ${BINARY_NAME}-mac-amd64 -ldflags="${FLAGS}"

win-arm64:
	$(info Building Windows ARM64 target...)
	GOOS=windows GOARCH=arm64 go build -o ${BINARY_NAME}-arm64.exe -ldflags="${FLAGS}"

win-amd64:
	$(info Building Windows AMD64 target...)
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}-amd64.exe -ldflags="${FLAGS}"

linux-arm64:
	$(info Building Linux ARM64 target...)
	GOOS=linux GOARCH=arm64 go build -o ${BINARY_NAME}-linux-arm64 -ldflags="${FLAGS}"

linux-amd64:
	$(info Building Linux AMD64 target...)
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}-linux-amd64 -ldflags="${FLAGS}"

all: clean
	@echo
	@echo Building all targets...
	go build -o ${BINARY_NAME} -ldflags="${FLAGS}"
	@echo
	GOOS=darwin GOARCH=arm64 go build -o ${BINARY_NAME}-mac-arm64 -ldflags="${FLAGS}"
	@echo
	GOOS=darwin GOARCH=amd64 go build -o ${BINARY_NAME}-mac-amd64 -ldflags="${FLAGS}"
	@echo
	GOOS=windows GOARCH=arm64 go build -o ${BINARY_NAME}-arm64.exe -ldflags="${FLAGS}"
	@echo
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}-amd64.exe -ldflags="${FLAGS}"
	@echo
	GOOS=linux GOARCH=arm64 go build -o ${BINARY_NAME}-linux-arm64 -ldflags="${FLAGS}"
	@echo
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}-linux-amd64 -ldflags="${FLAGS}"

list:
	$(info Available targets)
	$(info =====================)
	$(info build (or just 'make') - Builds the current platform)
	$(info mac-arm64)
	$(info mac-amd64)
	$(info win-arm64)
	$(info win-amd64)
	$(info linux-arm64)
	$(info linux-amd64)
	$(info all - builds all targets)
	$(info clean - removes all targets that have been built)


clean:
	@echo Cleaning GO build artifacts and removing targets...
	go clean
	-rm ${BINARY_NAME}-mac-arm64
	-rm ${BINARY_NAME}-mac-amd64
	-rm ${BINARY_NAME}-arm64.exe
	-rm ${BINARY_NAME}-amd64.exe
	-rm ${BINARY_NAME}-linux-arm64
	-rm ${BINARY_NAME}-linux-amd64
	-rm ${BINARY_NAME}-mac-arm64.tar
	-rm ${BINARY_NAME}-mac-amd64.tar
	-rm ${BINARY_NAME}-arm64.exe.zip
	-rm ${BINARY_NAME}-amd64.exe.zip
	-rm ${BINARY_NAME}-linux-arm64.tar
	-rm ${BINARY_NAME}-linux-amd64.tar

test:
	$(info Running all go tests...)
	go test ./... -count=1

prep: all
	@echo
	@echo Compressing all targets...
	zip -q timeconverter-arm64.exe.zip timeconverter-arm64.exe
	zip -q timeconverter-amd64.exe.zip timeconverter-amd64.exe
	tar -cf timeconverter-mac-arm64.tar timeconverter-mac-arm64
	tar -cf timeconverter-mac-amd64.tar timeconverter-mac-amd64
	tar -cf timeconverter-linux-arm64.tar timeconverter-linux-arm64
	tar -cf timeconverter-linux-amd64.tar timeconverter-linux-amd64
