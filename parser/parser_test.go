package parser

import (
	"../tokenizer"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	code := "ILLGÅNG FUNKÖN SKOGSFIBBLA hey SKOGSFIBBLA SKOGSFIBBLA nik zebi SKOGSFIBBLA ÄPPLARÖ FJÄLLBO"
	ast := Parse(tokenizer.Tokenize(code))

	excepted := []Node{
		{
			Type: CallExpression,
			Value: "ILLGÅNG",
			Params: []Node{
				{
					Type: String,
					Value: "hey",
				},
				{
					Type: String,
					Value: "nik zebi",
				},
			},
		},
		{
			Type: SemiColon,
			Value: tokenizer.SemiColon,
		},
	}

	if !reflect.DeepEqual(ast, excepted) {
		t.Errorf("Parse() failed, excepted %s got %s.", excepted, ast)
	}
}
