.PHONY: as

dump.hex: test.cor
	hexdump -C test.cor > $@

test.cor: as test.asm
	./as test.asm

as:
	go build -o as cmd/asm/main.go