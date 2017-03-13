# Go template toolkit

[![Build Status](https://travis-ci.org/go-qbit/template.svg?branch=master)](https://travis-ci.org/go-qbit/template)

## Description
The `ttgen` generates Go code from a template description. 
The template syntax looks like Perl Template-Toolkit (http://template-toolkit.org/), 
but not the same because Go is compiled language.  

## Performance
It's faster than native Go templates in about 7 times:
```
BenchmarkProcessTest-8              2000            879969 ns/op           48194 B/op       6008 allocs/op
BenchmarkCoreTemplate-8              200           5831531 ns/op         1042839 B/op      31035 allocs/op
```

## Installation
    go get github.com/go-qbit/template/ttgen

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
            [% <CONTENT> %]
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
        package template
        
        import (
            "context"
            "github.com/go-qbit/template/filter"
            "io"
        )
        
        func Wrapperpage(ctx context.Context, w io.Writer, tplClbF func(), caption string) {
            io.WriteString(w, "<!DOCTYPE html>\n<html>\n    <head>\n        <title>")
            io.WriteString(w, filter.HTML(caption))
            io.WriteString(w, "</title>\n    </head>\n    <body>")
            tplClbF()
            io.WriteString(w, "</<body>\n</html>")
        }
        ```
    1. `template/index.gtt` -> `template/index.gtt.go`:
        ```go
        package templates
        
        import (
            "context"
            "fmt"
            "github.com/go-qbit/template/filter"
            "io"
        )
        
        func ProcessTest(ctx context.Context, w io.Writer, header string, users []User) {
            Wrapperpage(ctx, w, func() {
                io.WriteString(w, "<h1>")
                io.WriteString(w, filter.HTML(header))
                io.WriteString(w, "</h1>\n    <hr>")
                for _, user := range users {
                    io.WriteString(w, "<p>")
                    ProcessUserName(ctx, w, user)
                    io.WriteString(w, ":")
                    io.WriteString(w, fmt.Sprint(user.Age))
                    io.WriteString(w, "(")
                    if user.IsMan {
                        io.WriteString(w, "Man")
                    } else {
                        io.WriteString(w, "Woman")
                    }
                    io.WriteString(w, ")\n        </p>")
                }
            }, header)
        }
        
        func ProcessUserName(ctx context.Context, w io.Writer, user User) {
            io.WriteString(w, filter.HTML(user.Name))
            io.WriteString(w, " ")
            io.WriteString(w, filter.HTML(user.Lastname))
        }
        ```