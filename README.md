# Welcome to helpers 👋

> Custom methods that can make developer's work easier.

## Avaliable Methods

### FindIndex

Returns the index position of an element in a slice using a callback function.

```golang
package main

import (
  "fmt"
	helpers "github.com/libercapital/helpers-go"
)

func findApple(fruit string) bool {
  return fruit == "Apple"
}

func main() {
  fruits := []string{"Banana", "Grape", "Apple"}

  fmt.Println(helpers.FindIndex(fruits, findApple)) // will print "2"
}
```

### Find

Returns an element of a slice using a callback function.

```golang
package main

import (
  "fmt"
	helpers "github.com/libercapital/helpers-go"
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

  fmt.Println(helpers.Find(fruits, findLime)) // will print "{Lime green}"
}
```

### Map

Applies the callback to the elements of the given slice.

```golang
package main

import (
  "fmt"
	helpers "github.com/libercapital/helpers-go"
)

func cube(number int) int {
  return (number * number * number)
}

func main() {
  numbers := []int{1, 2, 3, 4, 5}

  fmt.Println(helpers.Map(numbers, cube)) // will print "[1 8 27 64 125]"
}
```

### Filter

Filters elements of a slice using a callback function.

```golang
package main

import (
  "fmt"
	helpers "github.com/libercapital/helpers-go"
)

func odd(number int) bool {
  return (number % 2) == 1
}

func main() {
  numbers := []int{1, 2, 3, 4, 5}

  fmt.Println(helpers.Filter(numbers, odd)) // will print "[1 3 5]"
}
```

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_