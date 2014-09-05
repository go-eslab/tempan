library := libhotspot.a
archive := HotSpot-5.02.tar.gz
url := http://lava.cs.virginia.edu/HotSpot/grab/$(archive)

build := build

all: $(library)

$(library): $(build)
	$(MAKE) -C $(build)
	mv $(build)/$(library) $@

$(build): $(archive)
	mkdir $@
	tar -xzf $(archive) --strip=1 -C $@

$(archive):
	curl $(url) -o $@

install: $(library)
	go install

clean:
	rm -rf $(archive) $(build) $(library)

.PHONY: all install clean
