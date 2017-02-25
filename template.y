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
%type	<iAstNode>  top template header imports import_list vars var_list var body body_stmt expr loop condition
%type   <astFilter> filter

%%

top:        template                    { yylex.(*exprLex).result = $$  }

template:   header body                 { $$ = &astTemplate{$1, $2} }

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
        |   IDENTIFIER var_type         { $$ = &astVariableDef{$1, $2} }

var_type:   IDENTIFIER                  { $$ = $1 }
        |   '[' ']' IDENTIFIER          { $$ = "[]" + $3 }

body:       body_stmt                   { $$ = &astList{[]iAstNode{$1}} }
        |   body ';' body_stmt          { $$.(*astList).children = append($$.(*astList).children, $3) }

body_stmt:                              { $$ = nil }
        |   STRING                      { $$ = &astWriteString{&astString{$1}} }
        |   IDENTIFIER                  { $$ = &astWriteValue{&astValue{$1}} }
        |   IDENTIFIER '|' filter       { $3.value = &astValue{$1}; $$ = &astWriteString{$3} }
        |   STRING '|' filter           { $3.value = &astString{$1}; $$ = &astWriteString{$3} }
        |   loop                        { $$ = $1 }
        |   condition                   { $$ = $1 }


expr:       IDENTIFIER                  { $$ = &astValue{$1} }
        |   STRING                      { $$ = &astString{$1} }

filter:     IDENTIFIER                  { $$ = &astFilter{name: $1} }
        |   IDENTIFIER '(' ')'          { $$ = &astFilter{name: $1} }


loop:       FOR IDENTIFIER IN expr ';' body ';' END
                                        { $$ = &astLoop{$2, $4, $6} }

condition:  IF expr ';' body ';' END    { $$ = &astCondition{$2, $4, nil} }
        |   IF expr ';' body ';' ELSE ';' body ';' END
                                        { $$ = &astCondition{$2, $4, $8} }

%%