BINARY_NAME=timeconverter
SHORT_DATE := $(shell date +%Y-%m-%d\ %H:%M:%S)
LONG_DATE := $(shell date)
FLAGS := -X 'github.com/hobysmith/timeconverter/cmd.AppShortBuildTime=${SHORT_DATE}' -X 'github.com/hobysmith/timeconverter/cmd.AppLongBuildTime=${LONG_DATE}'
.DEFAULT_GOAL := build

.SILENT:

build:
	$(info Building default platform target...)
	go build -o ${BINARY_NAME} -ldflags="${FLAGS}"

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

bsd:
	$(info Building BSD target - Intel support only...)
	GOOS=freebsd GOARCH=amd64 go build -o ${BINARY_NAME}-bsd -ldflags="${FLAGS}"

all:
	$(info Building all targets...)

	$(info Building default platform target...)
	go build -o ${BINARY_NAME} -ldflags="${FLAGS}"

	$(info Building Mac ARM64 target...)
	GOOS=darwin GOARCH=arm64 go build -o ${BINARY_NAME}-mac-arm64 -ldflags="${FLAGS}"

	$(info Building Mac AMD64 target...)
	GOOS=darwin GOARCH=amd64 go build -o ${BINARY_NAME}-mac-amd64 -ldflags="${FLAGS}"

	$(info Building Windows ARM64 target...)
	GOOS=windows GOARCH=arm64 go build -o ${BINARY_NAME}-arm64.exe -ldflags="${FLAGS}"

	$(info Building Windows AMD64 target...)
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}-amd64.exe -ldflags="${FLAGS}"

	$(info Building Linux ARM64 target...)
	GOOS=linux GOARCH=arm64 go build -o ${BINARY_NAME}-linux-arm64 -ldflags="${FLAGS}"

	$(info Building Linux AMD64 target...)
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}-linux-amd64 -ldflags="${FLAGS}"

	$(info Building BSD target - Intel support only...)
	GOOS=freebsd GOARCH=amd64 go build -o ${BINARY_NAME}-bsd -ldflags="${FLAGS}"

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
	$(info bsd)
	$(info all - builds all targets)
	$(info clean - removes all targets that have been built)


clean:
	$(info Cleaning GO build artifacts and removing targets...)
	go clean
	-rm ${BINARY_NAME}-mac-arm64
	-rm ${BINARY_NAME}-mac-amd64
	-rm ${BINARY_NAME}-arm64.exe
	-rm ${BINARY_NAME}-amd64.exe
	-rm ${BINARY_NAME}-linux-arm64
	-rm ${BINARY_NAME}-linux-amd64
	-rm ${BINARY_NAME}-bsd

test:
	$(info Running all go tests)
	go test ./... -count=1
