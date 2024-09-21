# Makefile for building and running the project
OUTPUT_NAME := real_time_audio

BUILD_FLAGS = -ldflags "-s -w" 

# out put dir neetd to be created
TARGETS_DIR := sherpa_onnx

# Targets
all: build

build:
	@echo "Building to main.go with CGO enabled..."
	@echo "Building to main.go into ${BUILD_DIR}"
	@${GO} build ${BUILD_FLAGS} -o ${BUILD_DIR}/${TARGETS_DIR}/${OUTPUT_NAME} main.go

clean:
	@echo "Cleaning build directory..."
	@rm -rf ${BUILD_DIR}/${TARGETS_DIR}

mkdir:
	@echo "Creating build directory..."
	@mkdir -p ${BUILD_DIR}/${TARGETS_DIR}
