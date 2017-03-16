# Go template toolkit

[![Build Status](https://travis-ci.org/go-qbit/template.svg?branch=master)](https://travis-ci.org/go-qbit/template)

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
        
        var (
            s2ca1056cc2919b1c85f85e38e3a0041a = []byte{0x22, 0x3c, 0x21, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x22}
            s0d0afb06f292309a0fc79bab304b4b12 = []byte{0x22, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x22}
            s3c9300d2d2a49a76aa8a0ce68342b092 = []byte{0x22, 0x3c, 0x2f, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x22}
        )
        
        func Wrapperpage(ctx context.Context, w io.Writer, tplClbF func(), caption string) {
            w.Write(s2ca1056cc2919b1c85f85e38e3a0041a)
            io.WriteString(w, filter.HTML(caption))
            w.Write(s0d0afb06f292309a0fc79bab304b4b12)
            tplClbF()
            w.Write(s3c9300d2d2a49a76aa8a0ce68342b092)
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
            s96fe34f46084699530cb8cbd35895dda = []byte{0x22, 0x3c, 0x70, 0x3e, 0x22}
            s7b0f3af983749abb730d6180d0a43586 = []byte{0x22, 0x28, 0x22}
            sb653000c3bf94af208d7453ae8971fe5 = []byte{0x22, 0x4d, 0x61, 0x6e, 0x22}
            scd6c3f76caa0c7939d754b9b331e1c49 = []byte{0x22, 0x29, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x70, 0x3e, 0x22}
            sf3f1196886a8ba27a941f322228e4439 = []byte{0x22, 0x3c, 0x2f, 0x68, 0x31, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x72, 0x3e, 0x22}
            s5d70ad090a95c7bb49b92212141809dc = []byte{0x22, 0x3a, 0x22}
            sd5b84c0a6a6a38e2acb82de2283f56b9 = []byte{0x22, 0x57, 0x6f, 0x6d, 0x61, 0x6e, 0x22}
            sfcc3d7489d15ef49dbbf735234234cf7 = []byte{0x22, 0x20, 0x22}
            sa5f6581eccd129593b2a978d9e4fc581 = []byte{0x22, 0x3c, 0x68, 0x31, 0x3e, 0x22}
        )
        
        func ProcessTest(ctx context.Context, w io.Writer, header string, users []User) {
            Wrapperpage(ctx, w, func() {
                w.Write(sa5f6581eccd129593b2a978d9e4fc581)
                io.WriteString(w, filter.HTML(header))
                w.Write(sf3f1196886a8ba27a941f322228e4439)
                for _, user := range users {
                    w.Write(s96fe34f46084699530cb8cbd35895dda)
                    ProcessUserName(ctx, w, user)
                    w.Write(s5d70ad090a95c7bb49b92212141809dc)
                    io.WriteString(w, utils.ToString(user.Age))
                    w.Write(s7b0f3af983749abb730d6180d0a43586)
                    if user.IsMan {
                        w.Write(sb653000c3bf94af208d7453ae8971fe5)
                    } else {
                        w.Write(sd5b84c0a6a6a38e2acb82de2283f56b9)
                    }
                    w.Write(scd6c3f76caa0c7939d754b9b331e1c49)
                }
            }, header)
        }
        
        func ProcessUserName(ctx context.Context, w io.Writer, user User) {
            io.WriteString(w, filter.HTML(user.Name))
            w.Write(sfcc3d7489d15ef49dbbf735234234cf7)
            io.WriteString(w, filter.HTML(user.Lastname))
        }
        ```