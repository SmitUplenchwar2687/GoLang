package main

import (
	"errors"
	"fmt"
)

type structure struct {
	index int
	prop  string
	stre  structure1
}

type structure1 struct {
	ref int
}

func (x structure) multiply() int {
	return x.index * x.stre.ref
}

func main() {

	fmt.Println("Hello World!")
	var value int = demo(3)
	fmt.Printf("The value is retuned %v\n", value)

	i := 3

	var err error = iserror(i)
	fmt.Printf("%v", err)

	if err != nil {
		var text string = "Error exist\n"
		fmt.Printf(text)
	} else {
		text := "There is no error"
		fmt.Printf(text)
	}

	var Arr []int = []int{0, 1, 2, 3, 4}
	fmt.Println(Arr[1:3])

	Arr1 := []int{11, 22, 33, 44, 55}
	fmt.Printf("The len:%v, and the cap:%v\n", len(Arr1), cap(Arr1))
	Arr1 = append(Arr1, 7)
	fmt.Printf("The len:%v, and the cap:%v\n", len(Arr1), cap(Arr1))
	fmt.Println(Arr1)

	slice1 := make([]int, 3, 8)
	fmt.Println(slice1, len(slice1), cap(slice1))

	var Map map[int]string = make(map[int]string, 3)
	Map[0] = "First"
	Map = map[int]string{1: "First1", 2: "Second"}
	fmt.Println(Map)
	var val, bool = Map[2]
	delete(Map, 2)
	fmt.Println(val, bool)

	for i, v := range Arr1 {
		fmt.Println(i, v)
	}
	for i = 0; i < 10; i++ {
		fmt.Print(i)
	}

	mystring := []rune("This is an Éxample%@")
	for i, v := range mystring {
		fmt.Println(i, v)
	}
	fmt.Print(len(mystring), "\n")

	var struct1 structure = structure{index: 2, prop: "prop", stre: structure1{ref: 4}}
	fmt.Println(struct1.index, struct1.prop, struct1.stre.ref, struct1.multiply())

	//Pointers

	var p *int = new(int)
	var x int = 2
	*p = x
	fmt.Printf("the value of the refrence is %v\n", *p)
	fmt.Println(&x)

}

func demo(k int) int {

	fmt.Printf("The func is working %v\n", k)

	return k
}

func iserror(i int) error {

	if i%2 == 0 {
		var err error = errors.New("The is just an even error\n")
		return err
	} else {
		var err error = errors.New("This is just an odd alarm\n")
		return err
	}
}
