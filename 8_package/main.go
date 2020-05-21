package main

import (
	"fmt"
	"udemy/lessons/8_package/mylib"
	"udemy/lessons/8_package/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	under.Hello()
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
	mylib.Talib()
}
