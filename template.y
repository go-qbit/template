%{
package template;

import (
	__yyfmt__ "fmt"
)
%}

%union {
    string      string
	iAstNode    iAstNode
	astFilter   *astFilter
	astOutput   *astOutput
}

%token  IDENTIFIER
%token  STRING
%token  NUMBER
%token  FOR
%token  IN
%token  IF
%token  ELSE
%token  END
%token  VARS
%token  IMPORT

%type   <string>    STRING IDENTIFIER NUMBER var_type
%type	<iAstNode>  top template header imports import_list vars var_list var expr_list expr output loop
%type   <astFilter> filter

%%

top:        template                    { yylex.(*exprLex).result = $$  }

template:   header expr_list            { $$ = &astTemplate{$1, $2} }

header:     imports vars                { $$ = &astHeader{$1, $2} }

imports:                                { $$ = nil }
        | IMPORT '(' import_list ')' ';'{ $$ = $3 }

import_list: STRING                     { $$ = &astList{[]iAstNode{&astImport{$1}}} }
        | import_list ',' STRING        { $$.(*astList).children = append($$.(*astList).children, &astImport{$3}) }

vars:                                   { $$ = nil }
        | VARS '(' var_list ')' ';'     { $$ = $3 }

var_list:   var                         { $$ = &astList{[]iAstNode{$1}} }
        |   var_list ',' var            { $$.(*astList).children = append($$.(*astList).children, $3) }

var:                                    { $$ = nil }
        | IDENTIFIER var_type           { $$ = &astVariableDef{$1, $2} }

var_type: IDENTIFIER                    { $$ = $1 }
        | '[' ']' IDENTIFIER            { $$ = "[]" + $3 }

expr_list:  expr                        { $$ = &astList{[]iAstNode{$1}} }
        |   expr_list ';' expr          { $$.(*astList).children = append($$.(*astList).children, $3) }

expr:                                   { $$ = nil }
        |   STRING                      { $$ = &astWriteString{$1} }
        |   output                      { $$ = &astWriteValue{$1} }
        |   loop                        { $$ = $1 }

output:     IDENTIFIER                  { $$ = &astValue{$1} }
        |   STRING '|' filter           { $3.value = &astString{$1}; $$ = $3 }
        |   output '|' filter           { $3.value = $1; $$ = $3 }

filter:     IDENTIFIER                  { $$ = &astFilter{name: $1} }
        |   IDENTIFIER '(' ')'          { $$ = &astFilter{name: $1} }


loop:   FOR IDENTIFIER IN IDENTIFIER ';' expr_list ';' END
                                { $$ = &astLoop{$2, &astValue{$4}, $6} }
%%