build := hotspot
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

clean:
	rm -rf $(syso) $(build)/$(syso)
	$(MAKE) -C $(build) clean

.PHONY: all install clean
