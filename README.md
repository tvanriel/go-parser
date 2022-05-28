# go-parser

A simple parser helper written in go with the [go-lexer](https://github.com/bbuck/go-lexer/) in mind.


## Usage
```go
import(
    lex "github.com/bbuck/go-lexer"
)

func Parse(tokens []*lex.Token) (*parse.AST, error) {
	p := &parse.Parser{
		Tokens: tokens,
		Cur:    0,
		AST:    &parse.AST{},
	}
	for p.Continue() {
		ParseTextOrDirective(p)
	}
	if p.Err() != nil {
		return nil, p.Err()
	}
	return p.AST, nil
}

```