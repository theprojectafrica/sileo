# Intented For Just Development Purposes
# Linux/MacOS

.PHONY: default sileo clean run log

LOGFILE := ./build/sileo.log

default: sileo

sileo:
	mkdir -p build
	go build -o build/

clean:
	rm -rf ./build/*

run: clean sileo
	SILEO_LOGFILE=$(LOGFILE) ./build/sileo

log:
	> $(LOGFILE)
	tail -f $(LOGFILE)
