package compiler

import (
	"bytes"
	"corewar/consts"
	"corewar/pkg/asm/parser"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

type Compiler struct {
	ast        parser.ProgramNode
	labelTable map[string]int
}

func (c *Compiler) getArgTypeCode(args []parser.InstructionArgument) byte {
	var argTypeCode byte
	offset := 6
	for _, arg := range args {
		var currCode byte = 0
		currCode |= arg.Type.ByteCode
		currCode <<= offset
		argTypeCode |= currCode
		offset -= 2
	}
	return argTypeCode
}

func (c *Compiler) writeArgValue(r io.Writer, argVal interface{}) {
	_ = binary.Write(r, binary.LittleEndian, argVal)
}

func (c *Compiler) getChampNameBytes(n []byte) {
	for i, c := range c.ast.ChampName {
		n[i] = byte(c)
	}
}

func (c *Compiler) GetByteCode() []byte {
	var code bytes.Buffer
	nameBytes := make([]byte, consts.ChampNameLength)
	c.getChampNameBytes(nameBytes)
	code.Write(nameBytes)

	for _, cmd := range c.ast.Code.Commands {
		if cmd.Instruction == nil {
			continue
		}
		code.WriteByte(cmd.Instruction.Meta.OpCode)
		meta := cmd.Instruction.Meta
		if meta.IsArgTypeCode {
			code.WriteByte(c.getArgTypeCode(cmd.Instruction.Args))
		}
		for _, arg := range cmd.Instruction.Args {
			c.writeArgValue(&code, arg.Value)
		}
	}
	return code.Bytes()
}

func (c *Compiler) PrintAnnotatedCode() {
	annotations := &strings.Builder{}
	_, _ = fmt.Fprintf(annotations, "champion name: %s\n", c.ast.ChampName)
	for _, cmd := range c.ast.Code.Commands {
		if cmd.Instruction == nil {
			continue
		}
		_, _ = fmt.Fprintf(annotations, "%02x(%s) ",
			cmd.Instruction.Meta.OpCode, cmd.Instruction.Name)

		meta := cmd.Instruction.Meta
		if meta.IsArgTypeCode {
			_, _ = fmt.Fprintf(annotations, "%08b(argTypeCode) ",
				c.getArgTypeCode(cmd.Instruction.Args))
		}
		for _, arg := range cmd.Instruction.Args {
			_, _ = fmt.Fprintf(annotations, "%0x(%v, %v) ",
				arg.Value, arg.Token.Value, arg.Value)
		}
		_, _ = fmt.Fprintf(annotations, "\n")
	}
	fmt.Print(annotations.String())

}

func (c Compiler) Compile() error {
	err := c.validateMeta()
	if err != nil {
		return err
	}
	err = c.validateInstructions()
	if err != nil {
		return err
	}
	c.setOffsets()
	err = c.setupLabelsTable()
	if err != nil {
		return err
	}
	err = c.fillArgValues()
	if err != nil {
		return err
	}
	return nil
}

func NewCompiler(ast parser.ProgramNode) Compiler {
	return Compiler{ast: ast, labelTable: make(map[string]int)}
}
