package main

import (
	"fmt"
	"sync"
	"time"
)

type student struct {
	name string
	age  int
}

func (s *student) poolName () {
	fmt.Println("hello world")
	return
}

func newStudent(name string, age int) *student {
	return &student{
		name: name,
		age: age,
	}
}

type Classroom struct {
	name string
	classSess string
	intEd int
	student
}

func (c *Classroom) changeName (name string) {
	c.name = name
	return
}

func (s *student) changeName (name string) {
	s.name = name
	return
}

type NewJson struct {
	name, address string
	age int
}

type JsonTest struct {
	Name string `json:"name"`
	Age int 	`json:"age"`
	Address string 	`json:"address"`
}

type NotSlice struct {
	name string
	notSlice []string
}

func (n *NotSlice) ChooseSlice (notSlice []string) {
	n.notSlice = make([]string, len(notSlice))
	copy(n.notSlice, notSlice)

	//n.notSlice = notSlice
	return
}

type teacher struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Age int `json:"age"`
}

type teacher2 struct {
	Name string
	Students []*student
}

type Studented struct {
	ID     int
	Gender string
	Name   string
}

type Classed struct {
	Title    string `json:"title"`
	Students []*Studented `json:"students"`
}

func GoMax1 () {
	for i := 0; i < 10; i++ {
		fmt.Println("GoMax1", i)
	}
}
func GoMax2 () {
	for i := 0; i < 10; i++ {
		fmt.Println("GoMax2", i)
	}
}

func DeadGoroutine1 (ch chan int64) {
	for i := 0; i < len(ch); i++ {
		x := <- ch
		fmt.Println(x)
	}
}

func DeadGoroutine2 (ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func DeadGoroutine3 (ch1 chan int, ch2 chan int) {
	for {
		v, ok := <- ch1
		if !ok {
			break
		}
		ch2 <- v*v
	}
	close(ch2)
}

func InChannel1 (ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func InChannel2 (ch1 <-chan int, ch2 chan int) {
	for {
		v, ok := <- ch1
		if !ok {
			fmt.Println("Error <- ch1")
			break
		}
		ch2 <- v*v
	}

	close(ch2)
}

func worker (id int, jobs <-chan int, results chan<- int) {
	for j:= range jobs {
	fmt.Printf("start jobs %d value %d\n", id, j)
	time.Sleep(time.Second)
	fmt.Printf("end jobs %d value %d\n", id, j)
	results <- j*j
	}
}

var noX int = 10
var wg sync.WaitGroup
var lock sync.Mutex
var RwLock sync.RWMutex

func nox1()  {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		noX += i
		lock.Unlock()
	}
	wg.Done()
}

func nox2()  {
	for i := 0; i < 5000; i++ {
		RwLock.RLock()
		noX += i
		lock.Unlock()
	}
	wg.Done()
}

func readLock () {
	RwLock.RLock()
	//fmt.Println(noX)
	time.Sleep(3 * time.Second)
	RwLock.RUnlock()
	wg.Done()
}

func writeLock (){

	RwLock.Lock()
	noX += 1
	time.Sleep(3 * time.Millisecond)
	RwLock.Unlock()

	wg.Done()
}

var n int = 0
var once sync.Once

func func111 () {
	fmt.Println(n)
	//once.Do(func() {
	//	n += 1
	//})
	n += 1
	wg.Done()
}
func func222 () {
	fmt.Println(1)
	wg.Done()
}

var SyncMap sync.Map

func main() {

	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			SyncMap.Store(i, "int")

		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			SyncMap.Delete(i)

		}
		wg.Done()
	}()

	wg.Wait()

	SyncMap.Range(func(key, value interface{}) bool {
		fmt.Println(value)
		return true
	})

	//a := 97
	//fmt.Printf(strconv.Itoa(a))
	//fmt.Printf(string(a))


	//for i := 0; i <10; i++ {
	//	wg.Add(2)
	//	//once.Do(func1)
	//	go func111()
	//	go func222()
	//}
	//
	//wg.Wait()
	//
	//fmt.Println("N:", n)

	//start := time.Now()
	//
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go writeLock()
	//}
	//
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go readLock()
	//}
	//
	//end := time.Now()
	//fmt.Println(end.Sub(start))

	//wg.Add(4)
	//go readLock()
	//go readLock()
	//go writeLock()
	//go writeLock()
	//wg.Wait()
	//
	//fmt.Println(noX)

	//ch := make(chan int, 10)
	//
	//for i := 0; i < 10; i++ {
	//	select {
	//	case x := <- ch:
	//		fmt.Println(x)
	//	case ch <- i:
	//
	//	}
	//}

	/*
		worker pool
	 */
	//jobs := make(chan int, 100)
	//results := make(chan int, 100)
	//
	//for w := 0; w < 3; w++ {
	//	go worker(w, jobs, results)
	//}
	//
	//for j := 0; j < 5; j++ {
	//	jobs <- j
	//}
	//
	//close(jobs)
	//
	//for a := 1; a <= 5; a++ {
	//	<-results
	//}

	//now := time.Now()
	//secondS := now.Second()
	////fmt.Printf("%s:%d", now.Format("2006-01-02 15:04"), secondS)
	//times := fmt.Sprintf("%s:%d", now.Format("2006-01-02 15:04"), secondS)
	//
	//fmt.Println(reflect.TypeOf(times))

	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//
	//go InChannel1(ch1)
	//go InChannel2(ch1, ch2)
	//
	//for v := range ch2 {
	//	fmt.Println(v)
	//}

	//go DeadGoroutine2(ch1)
	//go DeadGoroutine3(ch1, ch2)
	//
	//for s := range ch2 {
	//	fmt.Println(s)
	//}

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch1 <- i
	//	}
	//	close(ch1)
	//}()
	//
	//go func() {
	//	for {
	//		v, ok := <-ch1
	//		if !ok {
	//			break
	//		}
	//		ch2 <- v*v
	//	}
	//	close(ch2)
	//}()
	//
	//for v := range ch2 {
	//	fmt.Println(v)
	//}

	//ch := make(chan int,1)
	 //ch <- 10
	 //
	 //x := <- ch
	 //fmt.Println(x)

	//ch := make(chan int64, 20)
	//
	//go DeadGoroutine1(ch)
	//
	////ch <- 10
	//for i := 0; i < 10; i++ {
	//	ch <- int64(i)
	//}
	//
	//time.Sleep(time.Second)

	//runtime.GOMAXPROCS(2)
	//go GoMax1()
	//go GoMax2()
	//
	//time.Sleep(time.Second)

	//str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"}]}`
	//c1 := &Classed{}
	//err := json.Unmarshal([]byte(str), c1)
	//if err != nil {
	//	fmt.Println("json unmarshal failed!")
	//	return
	//}
	//fmt.Printf("%#s\n", c1)

	//jsTest := `{
	//	"name":"stt",
	//	"address":"peiking",
	//	"age":18
	//}`

	//jsStu := `{"Name":"qqs","Students":[{"Name":"huangtian1","Age":19},{"Name":"huangtian2","Age":21},{"Name":"huangtian3","Age":19},{"Name":"huangtian6","Age":18}]}`
	//
	////c1 := &teacher{}
	//c1 := &teacher2{}
	//
	//err := json.Unmarshal([]byte(jsStu), c1)
	//if err != nil {
	//	fmt.Println("反json序列化失败")
	//	return
	//}
	//
	//fmt.Printf("%#v", c1)

	//tea := teacher{
	//	Name: "teacher",
	//	Address: "peiking",
	//	Age: 29,
	//}
	//
	//classes := []teacher{}
	//
	//for i := 0; i < 10; i ++ {
	//	classes = append(classes, tea)
	//}
	//
	//data, err := json.Marshal(classes)
	//if err != nil {
	//	fmt.Println("json序列化错误")
	//}
	//
	//fmt.Printf("%s", data)

	//p1 := NotSlice{
	//	name: "qqi",
	//}
	//data1 := []string{
	//	"name",
	//	"query",
	//	"anything",
	//}
	//
	//p1.ChooseSlice(data1)
	//
	//fmt.Println(p1)
	//data1[0] = "NONAME"
	//fmt.Println(p1)
	//fmt.Println(data1)

	//js := JsonTest{
	//	Name: "qwer",
	//	Age: 19,
	//	Address: "北京",
	//}
	//
	//data, err := json.Marshal(js)
	//if err != nil {
	//	fmt.Println("json marshal failed!")
	//	return
	//}
	//
	//fmt.Printf("%s", data)

	////jsonList := []NewJson{}
	//var jsonList = make(map[int]NewJson)
	//
	//
	//
	//for i := 1; i < 11; i++ {
	//	//jsonList = append(jsonList, NewJson{
	//	//	name: fmt.Sprintf("黄%d", i),
	//	//	address: fmt.Sprintf("北京%d路", i),
	//	//	age: 12 + i,
	//	//})
	//	jsonList[i] = NewJson{
	//		name: fmt.Sprintf("黄%d", i),
	//		address: fmt.Sprintf("北京%d路", i),
	//		age: 12 + i,
	//	}
	//}
	//
	//fmt.Println(jsonList,"\n")
	//fmt.Println(jsonList[1].name)

	////as := student{
	////	name: "stt",
	////	age: 16,
	////}
	//
	//clack := &Classroom{
	//	name: "qqa",
	//	classSess: "JOJO",
	//	intEd: 18,
	//	student: student{
	//		name: "stt",
	//		age: 16,
	//	},
	//}
	//
	////clack.student.name = "tts"
	//clack.changeName(clack.student.name)
	//clack.student.changeName("qqa")
	////clack.poolName()
	//
	//fmt.Println(clack)
	//fmt.Println(clack.student.name)

	//testStudent := newStudent("newStt", 18)
	//fmt.Println(testStudent)

	//m := make(map[string]*student)
	//stus := []student{
	//	{name: "小王子", age: 18},
	//	{name: "娜扎", age: 23},
	//	{name: "大王八", age: 9000},
	//}
	//
	//for _, stu := range stus {
	//
	//	m[stu.name] = &stu
	//}
	//
	////fmt.Println(m)
	//for k, v := range m {
	//	fmt.Println(k, "=>", v.name)
	//}
}
