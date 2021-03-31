package main

import "fmt"

type test1 struct {
	Name string
}

//var tMap map[int]*test1
//var t *test1
//var tt []int

func main() {
	//m := 5
	//x := &m
	//y := x
	//*y = 8
	//n := *x
	//n = 9
	//fmt.Println(*x, m, n)

	//tMap = make(Map[int]*test1)
	////t=&test1{Name: "snow"}
	////
	////tMap[0]=t
	//tMap[0] = &test1{Name: "snow"}
	////fmt.Println(t.Name)
	////t.Name="555"
	////fmt.Println(tMap[0].Name)
	//
	//tMap[1] = &test1{Name: "hello"}
	//tMap[2] = &test1{Name: "world"}
	//for _, v := range tMap {
	//	fmt.Println(*v)
	//	v.Name = "No." + v.Name
	//}
	//for _, v := range tMap {
	//	fmt.Println(*v)
	//}

	//tt:=[]int{1,2,3}
	//xx:=[]int{4,5,6}
	//tt=append(tt,xx[:]...)
	//fmt.Println(tt)

	//x:=[]int{5,7,6}
	//var y []int
	//y=x
	//fmt.Println(y)

	str := "hello world"
	for i := range str {
		fmt.Println(i, str[i])
		fmt.Println(fmt.Sprintf("%c", str[i]))
	}
}
