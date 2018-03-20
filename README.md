[![GoDoc](https://godoc.org/github.com/Zyqsempai/typos-resolver?status.svg)](https://godoc.org/github.com/Zyqsempai/typos-resolver)

# Typos Resolver

This is a simple typos converter(resolver) for keyboard layout mismatch, support two keyboards layouts,
russian and english and helps to map mismatched chars from one layout to another.

## Installation

```
go get github.com/Zyqsempai/typos-resolver
```

## Usage

Check [GoDoc](https://godoc.org/github.com/Zyqsempai/typos-resolver)

```go
package main

import (
    "fmt"
    "github.com/Zyqsempai/typos-resolver"
)

func main(){
    txt := "руддщ"
    fixed := typos.MakeReverse(txt, "en")
    fmt.Println(fixed)
    // print: hello

    txt2 := "cok"
    fixed2 := typos.SubstituteTypos(txt2, "ru")
    fmt.Println(fixed2)
    // print: сок
}
```