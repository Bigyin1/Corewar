package parser

import (
	"corewar/pkg/asm/tokenizer"
	"corewar/pkg/consts"
	"fmt"
)

type Parser struct {
	lexer        tokenizer.Tokenizer
	currentToken tokenizer.Token
}

func (p *Parser) eatLineBreaksSpaces() {
	for p.currentToken.Type == tokenizer.LineBreak || p.currentToken.Type == tokenizer.Space {
		_ = p.eatToken(p.currentToken.Type)
	}
}

func (p *Parser) eatLineBreaks() {
	for p.currentToken.Type == tokenizer.LineBreak {
		_ = p.eatToken(tokenizer.LineBreak)
	}
}

func (p *Parser) eatSpaces() {
	for p.currentToken.Type == tokenizer.Space {
		_ = p.eatToken(tokenizer.Space)
	}
}

func (p *Parser) eatToken(tokenType tokenizer.TokenType) error {
	if p.currentToken.Type == tokenizer.EOF {
		return tokenizer.EOFErr
	}
	if p.currentToken.Type == tokenType {
		p.currentToken = p.lexer.GetNextToken()
		return nil
	}
	return fmt.Errorf("unexpected token at line: %d, col: %d",
		p.currentToken.PosLine, p.currentToken.PosColumn)
}

func (p *Parser) argument(instrName consts.InstructionName) InstructionArgument {
	argType := p.currentToken.Type.GetArgType()
	arg := InstructionArgument{Token: p.currentToken, Type: argType}
	if p.currentToken.Type.IsDirectArgType() {
		arg.Type.Size = consts.InstructionsConfig[instrName].TDirSize // set alternative dir size
	}
	_ = p.eatToken(p.currentToken.Type)
	return arg
}

func (p *Parser) instruction() (InstructionNode, error) {
	var instrNode InstructionNode
	instrNode.Name = p.currentToken.Value.(consts.InstructionName)
	instrNode.Token = p.currentToken
	instrNode.Meta = consts.InstructionsConfig[instrNode.Name]
	err := p.eatToken(tokenizer.Instr)
	if err != nil {
		return InstructionNode{}, err
	}
	if p.currentToken.Type.IsOfArgType() {
		arg := p.argument(instrNode.Name)
		instrNode.Args = append(instrNode.Args, arg)
	} else {
		return InstructionNode{}, fmt.Errorf("got instruction w/o arguments; line: %d, col: %d",
			p.currentToken.PosLine, p.currentToken.PosColumn)
	}
	for p.currentToken.Type == tokenizer.Separator {
		err := p.eatToken(tokenizer.Separator)
		if err != nil {
			return InstructionNode{}, err
		}
		if p.currentToken.Type.IsOfArgType() {
			arg := p.argument(instrNode.Name)
			instrNode.Args = append(instrNode.Args, arg)
		} else {
			return InstructionNode{}, fmt.Errorf("got separator w/o instruction; line: %d, col: %d",
				p.currentToken.PosLine, p.currentToken.PosColumn)
		}
	}
	return instrNode, nil
}

func (p *Parser) command() (CommandNode, error) {
	cmdNode := CommandNode{}
	for p.currentToken.Type == tokenizer.Label {
		label := LabelNode{Name: p.currentToken.Value.(string), Token: p.currentToken}
		cmdNode.Labels = append(cmdNode.Labels, label)
		_ = p.eatToken(tokenizer.Label)
		p.eatLineBreaks()
	}
	p.eatLineBreaks()
	// for the case of label at the end of code block
	if p.currentToken.Type == tokenizer.Instr {
		instNode, err := p.instruction()
		if err != nil {
			return CommandNode{}, err
		}
		cmdNode.Instruction = &instNode
		err = p.eatToken(tokenizer.LineBreak)
		if err == tokenizer.EOFErr {
			return cmdNode, nil
		}
		if err != nil {
			return CommandNode{}, err
		}
	}
	return cmdNode, nil
}

func (p *Parser) codeBlock() (CodeNode, error) {
	codeNode := CodeNode{}
	if p.currentToken.Type == tokenizer.Label || p.currentToken.Type == tokenizer.Instr {
		cmdNode, err := p.command()
		if err != nil {
			return CodeNode{}, err
		}
		codeNode.Commands = append(codeNode.Commands, cmdNode)
	} else {
		return CodeNode{}, fmt.Errorf("no instructions provided")
	}
	p.eatLineBreaks()
	for p.currentToken.Type == tokenizer.Label || p.currentToken.Type == tokenizer.Instr {
		cmdNode, err := p.command()
		if err != nil {
			return CodeNode{}, err
		}
		codeNode.Commands = append(codeNode.Commands, cmdNode)
		p.eatLineBreaks()
	}
	return codeNode, nil
}

func (p *Parser) comment() (string, error) {
	err := p.eatToken(tokenizer.ChampComment)
	if err != nil {
		return "", err
	}
	commentVal := p.currentToken.Value
	err = p.eatToken(tokenizer.Str)
	if err != nil {
		return "", err
	}
	comment, ok := commentVal.(string)
	if !ok {
		return "", fmt.Errorf("unexpected error, could not get val for comment")
	}
	err = p.eatToken(tokenizer.LineBreak)
	if err != nil {
		return "", err
	}
	return comment, nil
}

func (p *Parser) name() (string, error) {
	err := p.eatToken(tokenizer.ChampName)
	if err != nil {
		return "", err
	}
	nameVal := p.currentToken.Value
	err = p.eatToken(tokenizer.Str)
	if err != nil {
		return "", err
	}
	name, ok := nameVal.(string)
	if !ok {
		return "", fmt.Errorf("unexpected error, could not get val for name")
	}
	err = p.eatToken(tokenizer.LineBreak)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (p *Parser) program() (ProgramNode, error) {
	champName := ""
	champComment := ""
	if p.currentToken.Type == tokenizer.ChampName {
		n, err := p.name()
		if err != nil {
			return ProgramNode{}, err
		}
		champName = n

		c, err := p.comment()
		if err != nil {
			return ProgramNode{}, err
		}
		champComment = c
	} else if p.currentToken.Type == tokenizer.ChampComment {
		c, err := p.comment()
		if err != nil {
			return ProgramNode{}, err
		}
		champComment = c
		n, err := p.name()
		if err != nil {
			return ProgramNode{}, err
		}
		champName = n
	}
	p.eatLineBreaks()
	codeNode, err := p.codeBlock()
	if err != nil {
		return ProgramNode{}, err
	}
	progNode := ProgramNode{
		Code:         codeNode,
		ChampName:    champName,
		ChampComment: champComment,
	}

	return progNode, nil

}

func (p *Parser) Parse() (ProgramNode, error) {
	p.currentToken = p.lexer.GetNextToken()
	ast, err := p.program()
	if err != nil {
		return ProgramNode{}, err
	}
	return ast, nil
}

func NewParser(lex tokenizer.Tokenizer) Parser {
	return Parser{lexer: lex}
}
