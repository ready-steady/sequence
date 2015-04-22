build := sobol
syso := main.syso

all: $(syso)

install: $(syso)
	go install

$(syso):
	$(MAKE) -C $(build)
	$(LD) -r -o $@ $(build)/*.o

clean:
	$(RM) -rf $(syso)
	$(MAKE) -C $(build) clean

.PHONY: all install clean
