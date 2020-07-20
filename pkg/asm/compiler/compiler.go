package compiler

import (
	"bytes"
	"corewar/pkg/asm/parser"
	"corewar/pkg/consts"
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
	_ = binary.Write(r, binary.BigEndian, argVal)
}

func (c *Compiler) writeMetaData(w *bytes.Buffer, codeLen uint64) {
	w.Grow(len(consts.MagicHeader) + consts.ProgNameLength + len(consts.NullSeq))

	w.WriteString(consts.MagicHeader)
	w.WriteString(c.ast.ChampName)
	for i := 0; i < consts.ProgNameLength-len(c.ast.ChampName); i++ {
		w.WriteByte(0)
	}

	w.WriteString(consts.NullSeq)

	_ = binary.Write(w, binary.BigEndian, codeLen)

	w.WriteString(c.ast.ChampComment)
	for i := 0; i < consts.CommentLength-len(c.ast.ChampComment); i++ {
		w.WriteByte(0)
	}

	w.WriteString(consts.NullSeq)
}

func (c *Compiler) GetByteCode() io.Reader {

	var code bytes.Buffer
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

	var header bytes.Buffer
	c.writeMetaData(&header, uint64(code.Len()))

	return io.MultiReader(&header, &code)
}

func (c *Compiler) PrintAnnotatedCode(w io.Writer) {
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
	_, _ = fmt.Fprint(w, annotations.String())

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
