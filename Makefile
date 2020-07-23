#asm
ASM_SRC_DIRS = \
	cmd/asm \
	pkg/asm \
	pkg/consts

ASM_SRC_FILES = $(foreach i,$(ASM_SRC_DIRS),$(shell find $(i) -name "*.go" -not -name "test_*.go"))

.PHONY: as
as:
	go build -o as cmd/asm/main.go


#tests
TEST_VM_DIR=pkg/corewar/testdata

ASMTESTFILES = $(wildcard $(TEST_VM_DIR)/*.asm)
CORFILES = $(ASMTESTFILES:.asm=.cor)

.PHONY: corewar-test
corewar-test: as $(CORFILES)
	go test ./pkg/corewar


%.cor: %.asm
	./as $<
