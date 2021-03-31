package huawei

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Huawei() {
	Init()
	wg.Add(2)
	b, err := ioutil.ReadFile("./huawei/training-1.txt")
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\n")
	serTypeNum, _ := strconv.Atoi(lineSlice[0])
	go initServerType(lineSlice[1 : serTypeNum+1])
	visualTypeNum, _ := strconv.Atoi(lineSlice[0+1+serTypeNum])
	go initVisualMachine(lineSlice[serTypeNum+2 : serTypeNum+2+visualTypeNum])
	//Day, _ = strconv.Atoi(lineSlice[serTypeNum+2+visualTypeNum])
	wg.Wait()
	//fmt.Println(Day)
	for i := serTypeNum + 2 + visualTypeNum + 1 + 1; i <= len(lineSlice)-1; i++ {
		//fmt.Println(lineSlice[i])
		if _, err := strconv.ParseFloat(lineSlice[i], 64); err == nil || i == len(lineSlice)-1 {
			ToStringRecord()
			ResetConfig()
			RemoveVisualMachine()
			//for _, v := range SurplusSeres {
			//	fmt.Printf("%+v\n", *v)
			//}
			continue
		}
		Judge(lineSlice[i])
	}
	fmt.Println("RMB is ", RMB)
}

