GO=go

all: as corewar

.PHONY: as
as:
	$(GO) build -o as cmd/asm/main.go

.PHONY: corewar
corewar:
	$(GO) build -o corewar cmd/corewar/main.go


EXAMPLES_DIR=./examples
ASMEXAMFILES = $(wildcard $(EXAMPLES_DIR)/*.s)
EXAMPCORFILES = $(ASMEXAMFILES:.s=.cor)

.PHONY: examples
examples: $(EXAMPCORFILES)
	@echo $(EXAMPCORFILES)

#tests
TEST_VM_DIR=pkg/corewar/testdata
ASMTESTFILES = $(wildcard $(TEST_VM_DIR)/*.s)
CORFILES = $(ASMTESTFILES:.s=.cor)

.PHONY: corewar-test
corewar-test: as $(CORFILES)
	@echo $(CORFILES)
	go test ./pkg/corewar
%.cor: %.s
	./as $<

