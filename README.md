# Persian Validator

Persian text validator in Golang.

## Installation

Use go get:
```
go get github.com/barahouei/persianvalidator
```

## How To Use

```
package main

import (
	"fmt"

	"github.com/barahouei/persianvalidator"
)

func main() {
	str := "سلام دنیا"

	isValid := persianvalidator.IsValid(str)

	fmt.Println(isValid)
}
```
