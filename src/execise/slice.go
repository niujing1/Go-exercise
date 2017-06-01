package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s)) //len计算长度 4
	fmt.Println("cap:", cap(s)) //cap计算容量 6
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s)) // 6
	fmt.Println("cap:", cap(s)) // 6

	s2 := s[2:3:4]
	fmt.Println("s2:", s2)
	fmt.Println("len:", len(s2))
	fmt.Println("cap:", cap(s2))

	fmt.Printf("%v\n", append(s, s2...))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("s11", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3: ", l)


	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3] //创建一个新的切片，长度为2,容量为4
	newSlice = append(newSlice, 60) //使用原来的容量来分配一个新的元素
	newSlice = append(newSlice, 0) //使用原来的容量来分配一个新的元素
	fmt.Println("3d:", newSlice)
	fmt.Printf("%d\n", len(newSlice)) // 5
	fmt.Printf("%d\n", cap(newSlice)) // 8 容量的增长可能是2倍，超过1000是1.25.。。
	fmt.Printf("%d\n", slice[3]) // 8 容量的增长可能是2倍，超过1000是1.25.。。

	s1 := []int{1, 2}
	sr2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, sr2...))
}
