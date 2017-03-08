%{
package template;

import (
	__yyfmt__ "fmt"
)
%}

%union {
    string          string
	iAstNode        iAstNode
	astList         *astList
	astUseWrapper   *astUseWrapper
	astFilter       *astFilter
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
%token  TEMPLATE
%token  IMPORT
%token  WRAPPER
%token  USE
%token  CONTENT_MARKER
%token  PROCESS
%token  EQ NE GE LE OR AND NOT

%type   <string>        STRING IDENTIFIER NUMBER var_type var_value
%type	<iAstNode>      top file header macros_stmt var body_stmt expr loop condition
%type   <astList>       imports import_list macroses param_list var_list body
%type   <astUseWrapper> use_wrapper
%type   <astFilter>     filter

%left '+' '-' '*' '/' '%' '>' '<' EQ NE GE LE OR AND
%left NOT

%%

top:        file                        { yylex.(*exprLex).result = $$  }

file:       header macroses             { $$ = &astFile{$1, $2} }

header:                                 { $$ = nil }
        |   imports                     { $$ = &astHeader{$1} }

imports:    IMPORT '(' import_list ')' ';'
                                        { $$ = $3 }

import_list: STRING                     { $$ = &astList{[]iAstNode{&astImport{$1}}} }
        |   import_list ',' STRING      { $$.Add(&astImport{$3}) }

macroses:   macros_stmt                 { $$ = &astList{[]iAstNode{$1}} }
        |   macroses ';' macros_stmt    { $$.Add($3) }

macros_stmt:                            { $$ = nil }
        |   TEMPLATE IDENTIFIER '(' var_list ')' use_wrapper ';' body ';' END
                                        { $$ = &astTemplate{$2, $4, $6, $8} }
        |   WRAPPER IDENTIFIER '(' var_list ')' ';' body ';' END
                                        { $$ = &astWrapper{$2, $4, $7} }
        |   STRING                      { $$ = nil }

use_wrapper:                            { $$ = nil }
        |   USE WRAPPER IDENTIFIER '(' param_list ')'
                                        { $$ = &astUseWrapper{$3, $5} }

param_list:                             { $$ = nil }
        |   expr                        { $$ = &astList{[]iAstNode{$1}} }
        |   param_list ',' expr         { $$.Add($3) }

var_list:   var                         { $$ = &astList{[]iAstNode{$1}} }
        |   var_list ',' var            { $$.Add($3) }

var:                                    { $$ = nil }
        |   IDENTIFIER var_type         { $$ = &astVariableDef{$1, $2} }

var_type:   IDENTIFIER                  { $$ = $1 }
        |   '[' ']' IDENTIFIER          { $$ = "[]" + $3 }
        |   '*' IDENTIFIER              { $$ = "*" + $2 }
        |   IDENTIFIER '.' IDENTIFIER   { $$ = $1 + "." + $3 }
        |   '*' IDENTIFIER '.' IDENTIFIER
                                        { $$ = "*" + $2 + "." + $4 }

body:       body_stmt                   { $$ = &astList{[]iAstNode{$1}} }
        |   body ';' body_stmt          { $$.Add($3) }

body_stmt:                              { $$ = nil }
        |   expr                        { $$ = &astWriteValue{$1} }
        |   loop                        { $$ = $1 }
        |   condition                   { $$ = $1 }
        |   CONTENT_MARKER              { $$ = &astWriteContent{} }
        |   PROCESS IDENTIFIER '(' param_list ')'
                                        { $$ = &astProcessTemplate{$2, $4} }
        | expr '|' filter               { $3.value = $1; $$ = &astWriteString{$3} }


expr:       expr '>' expr               { $$ = &astExpr{">", $1, $3} }
        |   expr '<' expr               { $$ = &astExpr{"<", $1, $3} }
        |   expr '+' expr               { $$ = &astExpr{"+", $1, $3} }
        |   expr '-' expr               { $$ = &astExpr{"-", $1, $3} }
        |   expr '*' expr               { $$ = &astExpr{"*", $1, $3} }
        |   expr '/' expr               { $$ = &astExpr{"/", $1, $3} }
        |   expr '%' expr               { $$ = &astExpr{"%", $1, $3} }
        |   expr EQ expr                { $$ = &astExpr{"==", $1, $3} }
        |   expr NE expr                { $$ = &astExpr{"!=", $1, $3} }
        |   expr GE expr                { $$ = &astExpr{">=", $1, $3} }
        |   expr LE expr                { $$ = &astExpr{"<=", $1, $3} }
        |   expr OR expr                { $$ = &astExpr{"||", $1, $3} }
        |   expr AND expr               { $$ = &astExpr{"&&", $1, $3} }
        |   NOT expr                    { $$ = &astExpr{"!", nil, $2} }
        |   '(' expr ')'                { $$ = &astParenthesis{$2} }
        |   IDENTIFIER '(' param_list ')'
                                        { $$ = &astFunc{$1, $3} }
        |   var_value                   { $$ = &astValue{$1} }
        |   STRING                      { $$ = &astString{$1} }
        |   NUMBER                      { $$ = &astValue{$1} }

var_value:  var_value '.' IDENTIFIER    { $$ = $1 + "." + $3 }
        |   var_value '[' NUMBER ']'    { $$ = $1 + "[" + $3 + "]" }
        |   var_value '[' STRING ']'    { $$ = $1 + "[" + $3 + "]" }
        |   '*' '(' var_value ')'       { $$ = "*(" + $3 + ")" }
        //ToDo:|   '*' var_value               { $$ = "*" + $2 }
        |   IDENTIFIER                  { $$ = $1 }


filter:     IDENTIFIER                  { $$ = &astFilter{name: $1} }
        |   IDENTIFIER '(' ')'          { $$ = &astFilter{name: $1} }


loop:       FOR IDENTIFIER IN expr ';' body ';' END
                                        { $$ = &astLoop{"_", $2, $4, $6} }
        |   FOR IDENTIFIER ',' IDENTIFIER IN expr ';' body ';' END
                                        { $$ = &astLoop{$2, $4, $6, $8} }

condition:  IF expr ';' body ';' END    { $$ = &astCondition{$2, $4, nil} }
        |   IF expr ';' body ';' ELSE ';' body ';' END
                                        { $$ = &astCondition{$2, $4, $8} }

%%