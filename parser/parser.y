%{
package units

// const UNIT := "UNIT"
// const NAMES := "NAMES"

func setResult(l yyLexer, v Result) {
  l.(*lex).result = v
}
%}


%union{
}

%token UNIT NAMES

%start main

%%

main:unitType
        {
            setResult(yylex, 0)
        }

unitType: UNIT NAMES


%%