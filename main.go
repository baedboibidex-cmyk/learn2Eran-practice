package main 

import (
	"fmt"
)

func Isaboy(name string)bool{
	return name == "boy"
}

func main() {
	fmt.Println(Isaboy("boy"))
	fmt.Println(Isagirl("girl"))
}

func Isagirl(name string)bool{
	return name == "girl"
}
