root := $(shell pwd)
source := $(root)/source
target := $(root)/target

clibrary := libcircuit.a
glibrary := main.syso

export OUTPUT_DIR := $(target)

all: $(glibrary)

install: $(glibrary)
	go install

$(glibrary): $(target)/$(clibrary)
	mkdir -p $(target)/objects
	cd $(target)/objects && ar x $<
	ld -r -o $@ $(target)/objects/*.o

$(target)/$(clibrary):
	mkdir -p $(target)
	$(MAKE) -C $(source)

clean:
	rm -rf $(target) $(glibrary)
	$(MAKE) -C $(source) clean

.PHONY: all install clean
