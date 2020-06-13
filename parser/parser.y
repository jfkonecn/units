%{
package units

// const UNIT := "UNIT"
// const NAMES := "NAMES"

func setResult(l yyLexer, v Result) {
  l.(*lex).result = v
}
%}


%union{
    str string
}

%token UNIT NAMES IS
%token <str> measurementName unitName

%type <str> measurementType unitList unitType

%start main

%%

main:measurementType
        {
            setResult(yylex, 0)
        }

measurementType : UNIT measurementName NAMES unitList
                    {
                        $$ = 1
                    }

unitList : unitType
            {
                $$ = 1
            }
         | unitType unitType
            {
                $$ = 1
            }

unitType : unitName IS
            {
                $$ = 1 
            }
%%