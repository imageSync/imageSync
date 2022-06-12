NAME := imageSync
VERSION := 1.0.9
OUTPUT_PATH := ./build/${VERSION}/

build-mac: baoName := ${NAME}-${VERSION}-mac-amd64
build-linux: baoName := ${NAME}-${VERSION}-linux-amd64
build-windows: baoName := ${NAME}-${VERSION}-windows-amd64





.PHONY: clean start build-mac
all: clean start build-mac


clean:
	rm -rf ./build/*
	rm -rf ./*.log
	rm -rf ./*.tar.gz
	rm -rf ./.imageSync


start:
	go mod tidy


build-linux:
	mkdir -p ${OUTPUT_PATH}${baoName}/
	GOARCH=amd64 \
	GOOS=linux \
	CGO_ENABLED=1 \
	CGO_LDFLAGS="-static" \
	CC=x86_64-linux-musl-gcc \
	CXX=x86_64-linux-musl-g++ \
	go build -o "${OUTPUT_PATH}${baoName}/${NAME}"
	tar -zcvf ${OUTPUT_PATH}${baoName}.tar.gz -C ${OUTPUT_PATH} ${baoName}


build-windows:
	mkdir -p ${OUTPUT_PATH}${baoName}/
	GOARCH=amd64 \
	GOOS=windows \
	CGO_ENABLED=1 \
	CGO_CFLAGS="-g -O2 -Wno-return-local-addr" \
	CC=x86_64-w64-mingw32-gcc \
	CXX=x86_64-w64-mingw32-g++ \
	go build -o "${OUTPUT_PATH}${baoName}/${NAME}"
	tar -zcvf ${OUTPUT_PATH}${baoName}.tar.gz -C ${OUTPUT_PATH} ${baoName}


build-mac:
	mkdir -p ${OUTPUT_PATH}${baoName}/
	GOARCH=amd64 GOOS=darwin CGO_ENABLED=1 \
	go build -o "${OUTPUT_PATH}${baoName}/${NAME}"
	tar -zcvf ${OUTPUT_PATH}${baoName}.tar.gz -C ${OUTPUT_PATH} ${baoName}
