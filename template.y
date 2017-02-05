%{
package template;

import (
	__yyfmt__ "fmt"
)
%}

%union {
    string      string
	iAstNode    iAstNode
}

%token  IDENTIFIER
%token  STRING
%token  NUMBER
%token  FOR
%token  IN
%token  IF
%token  ELSE
%token  END

%type   <string>    STRING IDENTIFIER
%type	<iAstNode>  NUMBER top expr_list expr loop

%%

top:        expr_list           { yylex.(*exprLex).result = $$  }

expr_list:  expr                { $$ = &astList{[]iAstNode{$1}} }
        |   expr_list ';' expr  { $$.(*astList).children = append($$.(*astList).children, $3) }

expr:                           { $$ = nil }
        |   STRING              { $$ = &astString{$1} }
        |   IDENTIFIER          { $$ = &astString{$1} }
        |   loop                { $$ = $1 }

loop:   FOR IDENTIFIER IN IDENTIFIER ';' expr_list ';' END
                                { $$ = &astLoop{$2, $4, $6} }
%%