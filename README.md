# Welcome to bava-helper 👋

[![Version](https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000)](https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000) [![pipeline status](https://gitlab.com/bavatech/architecture/software/libs/go-modules/bava-helper/badges/main/pipeline.svg)](https://gitlab.com/bavatech/architecture/software/libs/go-modules/bava-helper/-/commits/main) [![coverage report](https://gitlab.com/bavatech/architecture/software/libs/go-modules/bava-helper/badges/main/coverage.svg)](https://gitlab.com/bavatech/architecture/software/libs/go-modules/bava-helper/-/commits/main)

> Custom methods that can make developers' work easier.

## Avaliable Methods

### FindIndex

Returns the index position of an element in a slice using a callback function.

```golang
package main

import (
  "fmt"
	bavahelper "gitlab.com/bavatech/architecture/software/libs/go-modules/bava-helper.git"
)

func findApple(fruit string) bool {
  return fruit == "Apple"
}

func main() {
  fruits := []string{"Banana", "Grape", "Apple"}

  fmt.Println(bavahelper.FindIndex(fruits, findApple)) // will print "2"
}
```

### Find

Returns an element of a slice using a callback function.

```golang
package main

import (
  "fmt"
	bavahelper "gitlab.com/bavatech/architecture/software/libs/go-modules/bava-helper.git"
)

type fruit struct {
  name string
  color string
}

func findLime(fruit fruit) bool {
  return fruit.name == "Lime"
}

func main() {
  fruits := []fruit{
		{name: "Banana", color: "yellow"},
		{name: "Lime", color: "green"},
		{name: "Apple", color: "red"},
	}

  fmt.Println(bavahelper.Find(fruits, findLime)) // will print "{Lime green}"
}
```

### Map

Applies the callback to the elements of the given slice.

```golang
package main

import (
  "fmt"
	bavahelper "gitlab.com/bavatech/architecture/software/libs/go-modules/bava-helper.git"
)

func cube(number int) int {
  return (number * number * number)
}

func main() {
  numbers := []int{1, 2, 3, 4, 5}

  fmt.Println(bavahelper.Map(numbers, cube)) // will print "[1 8 27 64 125]"
}
```

### Filter

Filters elements of a slice using a callback function.

```golang
package main

import (
  "fmt"
	bavahelper "gitlab.com/bavatech/architecture/software/libs/go-modules/bava-helper.git"
)

func odd(number int) bool {
  return (number % 2) == 1
}

func main() {
  numbers := []int{1, 2, 3, 4, 5}

  fmt.Println(bavahelper.Filter(numbers, odd)) // will print "[1 3 5]"
}
```

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_