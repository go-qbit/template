# Go template toolkit

[![Build Status](https://travis-ci.org/go-qbit/template.svg?branch=master)](https://travis-ci.org/go-qbit/template)

* [Description](#description)
* [Performance](#performance)
* [Installation](#installation)
* [Usage](#usage)
* [Syntax](#syntax)
  + [Intro](#intro)
* [Directives](#directives)
  + [IMPORT](#import)
  + [TEMPLATE](#template)
  + [WRAPPER](#wrapper)
  + [PROCESS](#process)
  + [IF / ELSE](#if---else)
  + [FOR](#for)
* [Filters](#filters)
  + [HTML](#html)
* [Operators](#operators)
  + [Comparison](#comparison)
  + [Mathematical](#mathematical)
  + [Boolean](#boolean)
  + [Unary](#unary)

## Description
The `ttgen` generates Go code from a template description. 
The template syntax looks like Perl Template-Toolkit (http://template-toolkit.org/), 
but not the same because Go is compiled language.  

## Performance
It's faster than native Go templates in more than 10 times:
```
BenchmarkGoCoreTemplate-4           3000           5306959 ns/op         1041697 B/op      31035 allocs/op
BenchmarkQBitTemplate-4            30000            407632 ns/op           36096 B/op       4006 allocs/op
```

## Installation
    go get -u github.com/go-qbit/template
    go install github.com/go-qbit/template/ttgen

## Usage

1. Write a new wrapper for pages in `template/wrapper.gtt``:
    
    ```
    [% WRAPPER page(caption string) %]
        <!DOCTYPE html>
        <html>
        <head>
            <title>[% caption | HTML %]</title>
        </head>
        <body>
            [% CONTENT %]
        </<body>
        </html>
    [% END %]
    ```
1. Write a new template in `template/index.gtt``:
    
    ```
    [% TEMPLATE Test(header string, users []User) USE WRAPPER page(header) %]
        <h1>[% header | HTML %]</h1>
        <hr>
        [% FOR user IN users %]
            <p>
                [% PROCESS UserName(user) %]:
                [% user.Age %] ([% IF user.IsMan %]Man[% ELSE %]Woman[% END %])
            </p>
        [% END %]
    [% END %]
    
    [% TEMPLATE UserName(user User) %]
        [% user.Name | HTML +%] [% user.Lastname | HTML %]
    [% END %]
    ```
1. Define a `User` type in `template/user.go`:
    
    ```go
    type User struct {
        Name     string
        Lastname string
        Age      uint8
        IsMan    bool
    }
    ```
1. Write a go package that will process template:
    
    ```go
    //go:generate ttgen template/*.gtt
    package main
    
    import (
        "context"
        "os"
        "template"
    )
    
    func main()  {
         templates.ProcessTest(context.Background(), os.Stdout, "<Header>", []templates.User{
                {"Ivan", "Sidorov", 20, true},
                {"Petr", "Ivanov", 30, true},
            },
         )
    }
    ```
1. Generate Go code:
    `go generate ./...`
1. Check new Go files in the folder `template`, templates will be converted to:
    1. `template/wrapper.gtt` -> `template/wrapper.gtt.go`:
        ```go
        package templates
        
        import (
            "context"
            "github.com/go-qbit/template/filter"
            "io"
        )
        
        var (
            sc9e1fc75ad64f59917915544d65154cf = []byte{0x3c, 0x21, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e}
            sa25bcc91ec51d6fea14a629de184814c = []byte{0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e}
            s33e48fc8e2ea8a4250506db7e0be31fd = []byte{0x3c, 0x2f, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e}
        )
        
        func Wrapperpage(ctx context.Context, w io.Writer, tplClbF func(), caption string) {
            w.Write(sc9e1fc75ad64f59917915544d65154cf)
            io.WriteString(w, filter.HTML(caption))
            w.Write(sa25bcc91ec51d6fea14a629de184814c)
            tplClbF()
            w.Write(s33e48fc8e2ea8a4250506db7e0be31fd)
        }
        ```
    1. `template/index.gtt` -> `template/index.gtt.go`:
        ```go
        package template
        
        import (
            "context"
            "github.com/go-qbit/template/filter"
            "github.com/go-qbit/template/utils"
            "io"
        )
        
        var (
            s4da1a46ec20cf93ee5c846a51e04f0ed = []byte{0x3c, 0x70, 0x3e}
            s627661c621eab1b7b298abc47d1a250d = []byte{0x4d, 0x61, 0x6e}
            s5276343758bbfb9e0a374c9dd2fa6ce7 = []byte{0x57, 0x6f, 0x6d, 0x61, 0x6e}
            s921c704d576fc50eafc928f28d287019 = []byte{0x29, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x70, 0x3e}
            sa5c0aa5a33c99f4179eb48443e8bb33c = []byte{0x3c, 0x68, 0x31, 0x3e}
            s2b7c41aee73401e8cd970fba4c3bd4bd = []byte{0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x72, 0x3e}
            s853ae90f0351324bd73ea615e6487517 = []byte{0x3a}
            s84c40473414caf2ed4a7b1283e48bbf4 = []byte{0x28}
            s7215ee9c7d9dc229d2921a40e899ec5f = []byte{0x20}
        )
        
        func ProcessTest(ctx context.Context, w io.Writer, header string, users []User) {
            Wrapperpage(ctx, w, func() {
                w.Write(sa5c0aa5a33c99f4179eb48443e8bb33c)
                io.WriteString(w, filter.HTML(header))
                w.Write(s2b7c41aee73401e8cd970fba4c3bd4bd)
                for _, user := range users {
                    w.Write(s4da1a46ec20cf93ee5c846a51e04f0ed)
                    ProcessUserName(ctx, w, user)
                    w.Write(s853ae90f0351324bd73ea615e6487517)
                    io.WriteString(w, utils.ToString(user.Age))
                    w.Write(s84c40473414caf2ed4a7b1283e48bbf4)
                    if user.IsMan {
                        w.Write(s627661c621eab1b7b298abc47d1a250d)
                    } else {
                        w.Write(s5276343758bbfb9e0a374c9dd2fa6ce7)
                    }
                    w.Write(s921c704d576fc50eafc928f28d287019)
                }
            }, header)
        }
        
        func ProcessUserName(ctx context.Context, w io.Writer, user User) {
            io.WriteString(w, filter.HTML(user.Name))
            w.Write(s7215ee9c7d9dc229d2921a40e899ec5f)
            io.WriteString(w, filter.HTML(user.Lastname))
        }
        ```
        
## Syntax

### Intro
Template directives are embedded between tags `[%` and `%]`.
```
[% IF a == b %]
    text
[% END %]
```
Several directives within one block must be splitted by `;`.
```
[% 
    IF a == b;
        FOR v IN array;
            v;
        END;
    END 
%]
```
By default whitespaces around tags is chomped. The code
```
Before
[% v1 %] [% v2 %]
After
```
will produce `BeforeV1V2After` if `v1 == "V1"` and `v2 == "V2"`.
To save whitespaces use tag `[%+` to save left whitespaces and `+%]` to save right whitespaces.
The code
```
Before
[%+ v1 +%] [% v2 +%]
After
```
will produce 
```
Before
V1 V2
After
```
if `v1 == "V1"` and `v2 == "V2"`.

## Directives

### IMPORT
The `IMPORT` directive declares a list of go packages that will be added to result Go file.
Packages must be quoted by `"` and splitted by ",".  It's optional, but tt must be first directive in a template file if you add it.
```
[% 
    IMPORT(
        "package1",
        "package2"
    )
%]
```

### TEMPLATE
The `TEMPLATE` directive defines a new template:
```
 [% TEMPLATE <name>(<parameters_list>) USE WRAPPER <wrapper_name>(<parameters_list>) %]
    Template content
 [% END %]
```
`USE WRAPPER` is optional, see [WRAPPER](#wrapper) for more information.

### WRAPPER
The `WRAPPER` directive defines a new template wrapper. The place for content must be declared by `CONTENT`.
You can use wrappers to define common parts of templates, for example web site header and footer.
```
[% WRAPPER page(title string) %]
<!DOCTYPE html>
<html lang="ru">
    <head>
        <title>[% title | HTML %]</title>
    </head>
    <body>
        [% CONTENT %]
    </body>
</html>
[% END %]
```

Also you can use WRAPPER for WRAPPER:
```
[% WRAPPER page(title string) USE WRAPPER outer() %]
    [% CONTENT %]
[% END %]
```

### PROCESS
The `PROCESS` directive inserts output of other template to current.
```
[% PROCESS <template_name>(parameters) %]
```

### IF / ELSE
The `IF` / `ELSE` can be used to process or ignore a block based on some run-time condition.
```
[% IF a < 10 %]
    Less
[% ELSE %]
    Greater or equal
[% END %]
```

### FOR
The `FOR` directive will iterate through the items in a list/map
```
[% FOR item IN list %]
    [% item | HTML %]
[% END %]

[% FOR i, item IN list %]
    [% i %]: [% item | HTML %]
[% END %]

[% FOR key IN map %]
    [% key | HTML %]
[% END %]

[% FOR key, value IN map %]
    [% key | HTML %]: [% value | HTML %]
[% END %]
```
or while the last expression of the three is true
```
[% FOR i:=0; i<100; i++ %]
    [% i %]
[% END %]
```

## Filters
Any expression can be processed through filter before output
```
[% varName | filterName1 %]
[% (a + b) | filterName2(10) %]
```

### HTML
The `HTML` filter escaped dangerous character 

## Operators

### Assignments
1. `=`
1. `:=`

### Comparison
1. `>`, `GT`
1. `<`, `LT` 
1. `>=`, `GE`
1. `<=`, `LE` 
1. `==`, `EQ`  

### Mathematical
1. `+`
1. `-`
1. `*`
1. `/`
1. `%`

### Boolean
1. `&&`, `AND`
1. `||`, `OR`

### Unary
1. `!`, `NOT`
1. `++`
1. `--`