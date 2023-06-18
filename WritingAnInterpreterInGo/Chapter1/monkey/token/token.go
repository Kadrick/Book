package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // wrong token
	EOF     = "EOF"     // end-of-file

	/*================================ 식별자 + 리터럴 ==============================*/
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 1234567890

	/*================================ 연산자 ==============================*/
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	NOT      = "!"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="
	LTEQ   = "<="
	GTEQ   = ">="

	/*================================ 구분자 ==============================*/
	COMMA     = ","
	SEMICOLON = ";"
	QUOT      = "\""
	APOS      = "'"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	/*================================ 예약어 ==============================*/
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
