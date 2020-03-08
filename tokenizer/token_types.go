package tokenizer

type TokenType string

const (
	Str           TokenType = "STRING"
	Name                    = "CHAMP_NAME"
	Comment                 = "CHAMP_COMMENT"
	Instr                   = "INSTRUCTION"
	Space                   = "SPACE"
	Label                   = "LABEL"
	Separator               = "SEPARATOR"
	Register                = "REGISTER"
	Direct                  = "DIRECT"
	DirectLabel             = "DIRECT_LABEL"
	Indirect                = "INDIRECT"
	IndirectLabel           = "INDIRECT_LABEL"
	LineBreak               = "LINE_BREAK"
	Sum                     = "SUM"
	Sub                     = "SUB"
	EOF                     = "EOF"
)
