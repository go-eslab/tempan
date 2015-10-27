source := $(shell pwd)
output := $(OUTPUT_DIR)
hotspot := $(source)/hotspot

flags := -fPIC -O3

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
	cd $(hotspot) && git clean -df

.PHONY: clean
