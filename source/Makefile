root := $(ROOT_DIR)
output := $(OUTPUT_DIR)

flags := -fPIC -O3
source := $(root)/source
hotspot := $(source)/hotspot

$(output)/libcircuit.a: $(output)/libhotspot.a $(output)/circuit.o
	cp $(output)/libhotspot.a $@
	ar -q $@ $(output)/circuit.o

$(output)/libhotspot.a: $(hotspot)/Makefile
	$(MAKE) -C $(hotspot) -j$(NUM_JOBS) EXTRAFLAGS='$(flags)'
	cp $(hotspot)/libhotspot.a $@
	$(MAKE) -C $(hotspot) clean

$(output)/circuit.o: $(source)/circuit.c
	$(CC) $(flags) -I$(hotspot) -I$(source) -c $< -o $@

$(hotspot)/Makefile:
	git submodule update --init

clean:
	cd $(hotspot) && (git checkout . && git clean -df)

.PHONY: clean
