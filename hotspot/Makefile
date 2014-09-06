archive := HotSpot-5.02.tar.gz
url := http://lava.cs.virginia.edu/HotSpot/grab/$(archive)

build := build
syso := main.syso

all: $(syso)

install: $(syso)
	go install

$(syso): $(build)/libhotspot.a
	mkdir -p $(build)/$@
	cd $(build)/$@ && ar x ../libhotspot.a
	ld -r -o $@ $(build)/$@/*.o

$(build)/libhotspot.a: $(build)
	$(MAKE) -C $(build)

$(build): $(archive)
	mkdir $@
	tar -xzf $(archive) --strip=1 -C $@

$(archive):
	curl $(url) -o $@

clean:
	rm -rf $(archive) $(build) $(syso)

.PHONY: all install clean
