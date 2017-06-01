package main

import "fmt"

type myInt int

func (i *myInt) add (another int) myInt{
	*i = *i + myInt(another)
	return *i
}

func main() {
	i1 := myInt(1)
	i2 := (&i1).add(2)
	fmt.Println(i1, i2)
}
