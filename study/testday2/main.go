package main

import "fmt"

//func f1() int {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//
//func f2() (x int) {
//	defer func() {
//		x++
//	}()
//	return 5
//}
//
//func f3() (y int) {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//func f4() (x int) {
//	defer func(x int) {
//		x++
//	}(x)
//	return 5
//}
//func main() {
//	fmt.Println(f1())
//	fmt.Println(f2())
//	fmt.Println(f3())
//	fmt.Println(f4())
//}

const (
	name string = "wuHu"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("name是常量，无法修改")
		}
	}()



	x := 1
	y := 2

	// A 10 20 30
	// AA 10 30 40
	// B 10 20 30
	// BB 10 30 40
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

}