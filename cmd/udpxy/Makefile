# Description: Makefile for building udpxy-go for OpenWrt
BUILD_FLAGS = -ldflags "-s -w -extldflags '-static'" 
OUTPUT_NAME := udpxy-go

# out put dir neetd to be created
# {BUILD_DIR}/udpxy
UDPXY_DIR := udpxy

# Targets
all: build build-openwrt-amd64 build-openwrt-arm build-openwrt-mips

build: 
	@echo "Building to main.go with CGO enabled..."
	@CGO_ENABLED=0 ${GO} build ${BUILD_FLAGS} -tags netgo -o ${BUILD_DIR}/${UDPXY_DIR}/${OUTPUT_NAME} main.go

build-openwrt-amd64: 
	@echo "Building main.go for OpenWrt (amd64)..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build ${BUILD_FLAGS} -tags netgo -o ${BUILD_DIR}/${UDPXY_DIR}/openwrt_amd64_${OUTPUT_NAME} main.go

build-openwrt-arm: 
	@echo "Building main.go for OpenWrt (arm)..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm ${GO} build ${BUILD_FLAGS} -tags netgo -o ${BUILD_DIR}/${UDPXY_DIR}/openwrt_arm_${OUTPUT_NAME} main.go

build-openwrt-mips:
	@echo "Building main.go for OpenWrt (mips)..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=mips ${GO} build ${BUILD_FLAGS} -tags netgo -o ${BUILD_DIR}/${UDPXY_DIR}/openwrt_mips_${OUTPUT_NAME} main.go