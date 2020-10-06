package Test5

import "fmt"

func main() {
	Test5()
	a := [...]int{
		1, 2, 3, 5, 8, 8, 65, 4, 4,
	}

	for i, li := 0, len(a); i < li; i++ {
		fmt.Println(a[i])
	}

	b := []int{
		1, 5, 5, 74, 5, 7454, 54,
	}
}
