# Paths to tools needed in dependencies
GO := $(shell which go)
GIT := $(shell which git)

# Build flags
BUILD_MODULE := $(shell go list -m)
BUILD_FLAGS = -ldflags "-s -w" 

# Paths to locations, etc
BUILD_DIR := build
CMD_DIR := $(wildcard cmd/*)

# Targets
all: clean cmd

cmd: $(wildcard cmd/*)

$(CMD_DIR): dependencies mkdir
	@echo Build cmd $(notdir $@)
	@${GO} build ${BUILD_FLAGS} -o ${BUILD_DIR}/$(notdir $@) ./$@

FORCE:

dependencies:
	@test -f "${GO}" && test -x "${GO}"  || (echo "Missing go binary" && exit 1)
	@test -f "${GIT}" && test -x "${GIT}"  || (echo "Missing git binary" && exit 1)

mkdir:
	@echo Mkdir ${BUILD_DIR}
	@install -d ${BUILD_DIR}

clean:
	@echo Clean
	@rm -fr $(BUILD_DIR)
	@${GO} mod tidy
	@${GO} clean
