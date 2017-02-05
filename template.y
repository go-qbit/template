%{
package template;

import (
	__yyfmt__ "fmt"
)
%}

%union {
	string	  string
}

%token  IDENTIFIER
%token  STRING
%token  NUMBER
%token  FOR
%token  IN
%token  IF
%token  ELSE
%token  END

//%type	<string>	IDENTIFIER STRING NUMBER

%%

expr_list:  expr
        |   expr_list ';' expr

expr:
        |   STRING
        |   NUMBER
        |   IDENTIFIER
        |   loop

loop:   FOR IDENTIFIER IN IDENTIFIER ';' expr_list ';' END
%%