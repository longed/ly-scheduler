EXECUTABLE		:= ly-scheduler
PKG_DIR			:= $(EXECUTABLE).$(shell date +"%F"-"%H"-"%M"-"%S"-"%Z")
TAR_GZ			:= $(PKG_DIR).tar.gz
TARGET_DIR		:= targets
BIN_DIR			:= $(PKG_DIR)/bin
CONF_DIR		:= $(PKG_DIR)/conf
DOC_DIR			:= $(PKG_DIR)/doc
LIB_DIR			:= $(PKG_DIR)/lib
CONFIG_FILE		:= config.toml

#--------------------------------------------------------
# build

GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

#GO              := CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GO111MODULE=on go
GO              := GO111MODULE=on go
GOBUILD         := $(GO) build -gcflags="all=-N -l" -o $(EXECUTABLE) .

.PHONY: integrate build clean test

integrate: build
	# create directories by script
	bash assemly.sh mkdir_loop $(TARGET_DIR) $(TARGET_DIR)/$(BIN_DIR) $(TARGET_DIR)/$(CONF_DIR) $(TARGET_DIR)/$(DOC_DIR) $(TARGET_DIR)/$(LIB_DIR)

	# copy files | config-file scripts README.md docs/*
	# yes | cp -t $(TARGET_DIR)/$(CONF_DIR) $(CONFIG_FILE)
	yes | cp -r -t $(TARGET_DIR)/$(PKG_DIR) misc/bin
	yes | cp -r -t $(TARGET_DIR)/$(PKG_DIR) misc/doc
	yes | cp -r -t $(TARGET_DIR)/$(CONF_DIR) src/$(CONFIG_FILE)
	
	# move files 
	yes | mv -t $(TARGET_DIR)/$(LIB_DIR) $(EXECUTABLE)
	
	# compress targets
	cd $(TARGET_DIR) &&	tar cvzf $(TAR_GZ) * && cd -

build: clean
	cd src;	$(GOBUILD);	mv $(EXECUTABLE) ../; cd -

clean:
	-rm -rf $(TARGET_DIR)
	-rm -rf $(EXECUTABLE)

test: