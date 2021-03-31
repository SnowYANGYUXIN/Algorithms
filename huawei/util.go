package huawei

import (
	"fmt"
	"strconv"
	"strings"
)

func Judge(line string) {
	line = strings.Replace(line, " ", "", -1)
	line = strings.Trim(line, "()")
	str := strings.Split(line, ",")
	if str[0] == "add" {
		id, _ := strconv.Atoi(str[2])
		AddVisualMachine(str[1], id)
	} else if str[0] == "del" {
		id, _ := strconv.Atoi(str[1])
		DeleteVisualMachine(id)
	}
}

func ResetConfig() {
	Purchase = 0
	Record = []string{}
	Migration = 0
	PurchaseSer = make(map[string]int)
	MigrationNum = make(map[int]int)
}

//打印购买操作
func ToStringRecord() {
	fmt.Printf("(purchase, %d)\n", Purchase)
	for i, v := range PurchaseSer {
		fmt.Printf("(%s, %d)\n", i, v)
	}

	fmt.Printf("(migration, %d)\n", Migration)
	for k, v := range MigrationNum {
		fmt.Printf("(%d, %d)\n", k, v)
	}

	for _, v := range Record {
		fmt.Println(v)
	}
	for _, v := range SurplusSeres {
		if v.NodeACPUHaveNum != 0 || v.NodeAMemoryHaveNum != 0 || v.NodeBCPUHaveNum != 0 || v.NodeBMemoryHaveNum != 0 {
			RMB += v.MoneyDay
		}
	}
}
