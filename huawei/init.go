package huawei

import (
	"strconv"
	"strings"
)

func Init() {
	ServerTypeID = 0
	ServerTypeMap = []*Service{}
	VisualMachineTypeMap = make(map[string]VisualMachine)
	SerMap = []*Service{}
	MachineSer = make(map[int]*Service)
	SurplusSeres = make(map[int]*SurplusServer)
	PurchaseSer = make(map[string]int)
	MigrationNum = make(map[int]int)
	SurplusSer = &SurplusServer{
		NodeACPUHaveNum:    0,
		NodeBCPUHaveNum:    0,
		NodeAMemoryHaveNum: 0,
		NodeBMemoryHaveNum: 0,
		NodeA: ServerNode{
			CPU:    -1,
			Memory: -1,
		},
		NodeB: ServerNode{
			CPU:    -1,
			Memory: -1,
		},
		VisualMachines: []int{},
	}
	Purchase = 0
	Record = []string{}
	Migration = 0
	RMB = 0
}

func initServerType(sli []string) {
	for _, v := range sli {
		v = strings.Replace(v, " ", "", -1)
		v = strings.Trim(v, "()")
		str := strings.Split(v, ",")

		cpu, _ := strconv.Atoi(str[1])
		memory, _ := strconv.Atoi(str[2])
		money, _ := strconv.Atoi(str[3])
		moneyDay, _ := strconv.Atoi(str[4])
		ServerTypeMap = append(ServerTypeMap, &Service{
			ID:       ServerTypeID,
			Name:     str[0],
			CPU:      cpu,
			Memory:   memory,
			Money:    money,
			MoneyDay: moneyDay,
		})
		ServerTypeID++
	}
	//fmt.Println(*ServerTypeMap[0])
	wg.Done()
}

func initVisualMachine(sli []string) {
	for _, v := range sli {
		v = strings.Replace(v, " ", "", -1)
		v = strings.Trim(v, "()")
		str := strings.Split(v, ",")

		cpu, _ := strconv.Atoi(str[1])
		memory, _ := strconv.Atoi(str[2])
		isDouble, _ := strconv.Atoi(str[3])
		VisualMachineTypeMap[str[0]] = VisualMachine{
			ID:       VisualMachineID,
			Name:     str[0],
			CPU:      cpu,
			Memory:   memory,
			isDouble: isDouble,
		}
		VisualMachineID++
	}
	//fmt.Println(VisualMachineTypeMap)
	wg.Done()
}
