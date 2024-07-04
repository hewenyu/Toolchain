# Paths to tools needed in dependencies
export GO := $(shell which go)
export GIT := $(shell which git)

# Build flags
BUILD_MODULE := $(shell go list -m)
BUILD_FLAGS = -ldflags "-s -w" 

# Paths to locations, etc
export BUILD_DIR := $(abspath ./build)
CMD_DIR := $(wildcard cmd/*)

# Targets
all: clean cmd

cmd: $(wildcard cmd/*)


$(CMD_DIR): dependencies mkdir
	@echo Build cmd $(notdir $@)
	@${MAKE} -C $@

# $(CMD_DIR): dependencies mkdir
# 	@echo Build cmd $(notdir $@)
# 	@${GO} build ${BUILD_FLAGS} -o ${BUILD_DIR}/$(notdir $@) ./$@

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
