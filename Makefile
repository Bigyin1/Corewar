.PHONY: as

test.cor: as test.asm
	./as test.asm
	 hexdump  test.cor

as:
	go build -o as cmd/asm/main.go