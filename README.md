# Typos Resolver

This is a simple typos converter(resolver) for keyboard layout mismatch, support two keyboards layouts,
russian and english and helps to map mismatched chars from one layout to another.

## Installation

```
go get github.com/Zyqsempai/typos-resolver
```

## Usage

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
}
```