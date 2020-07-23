package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func mathLen() {
	s := "123abc你好啊"
	counts := 0;
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		//fmt.Println(len(string(runeS[i])))
		//fmt.Println(string(runeS[i]))
		if len(string(runeS[i])) == 3 || len(string(runeS[i])) == 4 {
			counts++
			fmt.Println(string(runeS[i]))
		}
	}
	fmt.Println(counts)
}

func ifThings() {
	if a := 10; a < 11 {
		fmt.Println("A")
	}
}

func testTag() {
	breaksTag:
		for i := 0; i < 10; i++ {
			for j := 1; j < 10; j++ {
				fmt.Println(j)
				if i == 1 {
					break breaksTag
				}
			}
			fmt.Println(i)
		}

	fmt.Println("BreakTrue")
}

func testArray() {
	arr1 := [3][2]string{
		{"北京", "上海"},
		{"南京", "苏州"},
		{"青岛", "济南"},
	}

	for _, a1 := range arr1{
		for _, a2 := range a1{
			fmt.Println(a2)
		}
	}

	fmt.Println(arr1)
}

func testSlice()  {
	s1 := []int{1,2,3,4,5}
	//s2 := [][]int{
	//	{1,3,4,5},
	//	{6,7,8,9},
	//}
	s3 := []int{1,2,3,4,5}

	s1 = append(s1, s3...)

	fmt.Println(s1)
}

func testMap() {
	// 用来初始化随机种子
	rand.Seed(time.Now().UnixNano())
	// 定义map
	scoreMap := make(map[string]int, 200)

	// 组成随机map
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)

		scoreMap[key] = value
	}

	// 循环添加key到keys切片
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// 排序
	sort.Strings(keys)
	// 输出
	for _, key := range keys{
		fmt.Printf("%s => %v\n", key, scoreMap[key])
	}

}

func lenEng(str string) (maps map[string]int) {
	strs := strings.Split(str, " ")

	maps = make(map[string]int, 100)
	for _, value := range strs {
		maps[value] = 0
	}

	for _, value := range strs{
		maps[value] += 1
	}

	return
}

func lenEng2(str string) (maps map[string]int) {
	strs := strings.Split(str, " ")

	maps = make(map[string]int, 100)

	for _, value := range strs {
		_, ok := maps[value]
		if !ok {
			maps[value] = 1
		} else {
			maps[value] += 1
		}
	}

	return
}

func testError() error {
	return errors.New("Error!")
}

func main()  {
	//fmt.Println(lenEng2("how do you do"))
	fmt.Println(testError())

	//testMap()

	//testSlice()

	//s := "stt"
	//fmt.Println(strings.Contains(s, "q"))
	//fmt.Println(strings.LastIndex(s, "t"))
	//fmt.Println(strings.HasPrefix(s, "s"))
	//sl := strings.Split(s, "")
	//sf := strings.Join(sl, "=")
	//fmt.Println(sf)
	//for _, skv := range sl {
	//	fmt.Println(skv)
	//}
	//
	//for i := 0; i < len(sl); i++ {
	//	fmt.Printf("%v\n", sl[i])
	//}
	//fmt.Println(sl)

	// ============================================================
	//nameA := "黄麒英"
	//runeN1 := []rune(nameA)
	//fmt.Printf("%T\n", runeN1)
	//runeN1[1] = '飞'
	//runeN1[2] = '鸿'
	//nameA = string(runeN1)
	//fmt.Println(nameA)
	//
	//
	//nameB := "huangqiying"
	//byteN1 := []byte(nameB)
	//fmt.Printf("%T\n", byteN1)
	//byteN1[0] = 'H'
	//byteN1[5] = 'Q'
	//byteN1[7] = 'Y'
	//nameB = string(byteN1)
	//fmt.Println(nameB)

	// ============================================================
	//mathLen()

	//ifThings()

	//testTag()

	//testArray()

	//var a = make([]string, 5, 10)
	//for i := 0; i < 10; i++ {
	//	a = append(a, fmt.Sprintf("%v", i))
	//}
	//fmt.Println(len(a))
}