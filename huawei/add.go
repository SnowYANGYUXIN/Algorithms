package huawei

import "fmt"

//TODO:如果是单机部署应该是AB结点交错部署最好
func AddVisualMachine(name string, id int) {
	visual := VisualMachineTypeMap[name]
	//fmt.Println(id,visual)
	if visual.isDouble == 0 {
		//TODO:这里要设计一个在AB之间部署的一个平衡点
		if visual.CPU <= SurplusSer.NodeA.CPU && visual.Memory <= SurplusSer.NodeA.Memory && SurplusSer.NodeA.CPU > SurplusSer.NodeB.CPU {
			SurplusSer.NodeA.CPU -= visual.CPU
			SurplusSer.NodeA.Memory -= visual.Memory
			SurplusSer.NodeACPUHaveNum += visual.CPU
			SurplusSer.NodeAMemoryHaveNum += visual.Memory

			SerMap[SurplusSer.ServerID].CPU -= visual.CPU
			SerMap[SurplusSer.ServerID].Memory -= visual.Memory
			//用于删除虚拟机时可以查询该虚拟机的具体数据

			//TODO:不能直接MachineSer[id] = SerMap[SurplusSer.ServerID] 因为都是指针类型，这样VisualMachineName和NodeName每一次都在改变
			//SerMap[SurplusSer.ServerID].VisualMachineName = visual.Name
			//SerMap[SurplusSer.ServerID].NodeName = "A"
			//增加虚拟机id映射服务器
			//MachineSer[id] = SerMap[SurplusSer.ServerID]
			MachineSer[id] = &Service{
				ID:                SerMap[SurplusSer.ServerID].ID,
				Name:              SerMap[SurplusSer.ServerID].Name,
				CPU:               SerMap[SurplusSer.ServerID].CPU,
				Memory:            SerMap[SurplusSer.ServerID].Memory,
				VisualMachineName: visual.Name,
				NodeName:          "A",
			}
			SurplusSer.VisualMachines = append(SurplusSer.VisualMachines, id)
			//MachineSer[id].VisualMachineName=visual.Name
			Record = append(Record, fmt.Sprintf("%d A", SurplusSer.ServerID))
			return
		} else if visual.CPU <= SurplusSer.NodeB.CPU && visual.Memory <= SurplusSer.NodeB.Memory {
			SurplusSer.NodeB.CPU -= visual.CPU
			SurplusSer.NodeB.Memory -= visual.Memory
			SurplusSer.NodeBCPUHaveNum += visual.CPU
			SurplusSer.NodeBMemoryHaveNum += visual.Memory

			SerMap[SurplusSer.ServerID].CPU -= visual.CPU
			SerMap[SurplusSer.ServerID].Memory -= visual.Memory
			//用于删除虚拟机时可以查询该虚拟机的具体数据
			//SerMap[SurplusSer.ServerID].VisualMachineName = visual.Name
			//SerMap[SurplusSer.ServerID].NodeName = "B"
			//增加虚拟机id映射服务器
			//MachineSer[id] = SerMap[SurplusSer.ServerID]
			//MachineSer[id].VisualMachineName=visual.Name
			MachineSer[id] = &Service{
				ID:                SerMap[SurplusSer.ServerID].ID,
				Name:              SerMap[SurplusSer.ServerID].Name,
				CPU:               SerMap[SurplusSer.ServerID].CPU,
				Memory:            SerMap[SurplusSer.ServerID].Memory,
				VisualMachineName: visual.Name,
				NodeName:          "B",
			}
			SurplusSer.VisualMachines = append(SurplusSer.VisualMachines, id)
			Record = append(Record, fmt.Sprintf("%d B", SurplusSer.ServerID))
			return
		} else {
			//重新选择剩余服务器
			ReplaceSurSer(visual.CPU, visual.Memory, name, id)
			return
		}
	} else if visual.isDouble == 1 {
		if SurplusSer.NodeA.CPU >= visual.CPU/2 && SurplusSer.NodeB.CPU >= visual.CPU/2 &&
			SurplusSer.NodeA.Memory >= visual.Memory/2 && SurplusSer.NodeB.Memory >= visual.Memory/2 {
			SurplusSer.NodeA.CPU -= visual.CPU / 2
			SurplusSer.NodeB.CPU -= visual.CPU / 2
			SurplusSer.NodeA.Memory -= visual.Memory / 2
			SurplusSer.NodeB.Memory -= visual.Memory / 2
			SurplusSer.NodeACPUHaveNum += visual.CPU / 2
			SurplusSer.NodeBCPUHaveNum += visual.CPU / 2
			SurplusSer.NodeAMemoryHaveNum += visual.Memory / 2
			SurplusSer.NodeBMemoryHaveNum += visual.Memory / 2

			SerMap[SurplusSer.ServerID].CPU -= visual.CPU
			SerMap[SurplusSer.ServerID].Memory -= visual.Memory
			//用于删除虚拟机时可以查询该虚拟机的具体数据
			//SerMap[SurplusSer.ServerID].VisualMachineName = visual.Name
			//SerMap[SurplusSer.ServerID].NodeName = "AB"
			//增加虚拟机id映射服务器
			//MachineSer[id] = SerMap[SurplusSer.ServerID]
			//MachineSer[id].VisualMachineName=visual.Name
			MachineSer[id] = &Service{
				ID:                SerMap[SurplusSer.ServerID].ID,
				Name:              SerMap[SurplusSer.ServerID].Name,
				CPU:               SerMap[SurplusSer.ServerID].CPU,
				Memory:            SerMap[SurplusSer.ServerID].Memory,
				VisualMachineName: visual.Name,
				NodeName:          "AB",
			}
			SurplusSer.VisualMachines = append(SurplusSer.VisualMachines, id)
			Record = append(Record, fmt.Sprintf("%d", SurplusSer.ServerID))

		} else {
			//重新选择剩余服务器
			ReplaceSurSer(visual.CPU, visual.Memory, name, id)
		}
	}
	return
}

func ReplaceSurSer(cpu, memory int, name string, id int) {
	for i, v := range SurplusSeres {
		if v.NodeA.CPU >= cpu && v.NodeB.CPU >= cpu && v.NodeA.Memory >= memory && v.NodeB.Memory >= memory {
			SurplusSer = SurplusSeres[i]
			AddVisualMachine(name, id)
			return
		}
	}
	BuyServer(cpu, memory, id)
	AddVisualMachine(name, id)
}

//TODO:需要算法得买哪一种服务器更划算以及虚拟机的迁移  cpu和memory的倍数可以做参考
func BuyServer(cpu, memory, id int) {
	for _, v := range ServerTypeMap {
		if v.CPU > 3*cpu && v.Memory > 3*memory {
			SerMap = append(SerMap, &Service{
				ID:     len(SerMap),
				Name:   v.Name,
				CPU:    v.CPU,
				Memory: v.Memory,
			})

			SurplusSer = &SurplusServer{
				ServerID:           len(SerMap) - 1, //这里要减1，因为通过上面的append使serMap的len加1
				NodeACPUHaveNum:    0,
				NodeBCPUHaveNum:    0,
				NodeAMemoryHaveNum: 0,
				NodeBMemoryHaveNum: 0,
				SumCpuNum:          v.CPU,
				SumMemoryNum:       v.Memory,
				NodeA: ServerNode{
					CPU:    v.CPU / 2,
					Memory: v.Memory / 2,
				},
				NodeB: ServerNode{
					CPU:    v.CPU / 2,
					Memory: v.Memory / 2,
				},
				MoneyDay:       v.MoneyDay,
				VisualMachines: []int{},
			}
			SurplusSeres[len(SerMap)-1] = SurplusSer

			Purchase++
			PurchaseSer[v.Name] += 1
			RMB = v.Money
			return
		}
	}
	//TODO：如果没找到就把最好的服务器给出来
	//暂时为最后一个
	SerMap = append(SerMap, &Service{
		ID:     len(SerMap),
		Name:   ServerTypeMap[len(ServerTypeMap)-1].Name,
		CPU:    ServerTypeMap[len(ServerTypeMap)-1].CPU,
		Memory: ServerTypeMap[len(ServerTypeMap)-1].Memory,
	})

	SurplusSer = &SurplusServer{
		ServerID:           len(SerMap) - 1,
		NodeACPUHaveNum:    0,
		NodeBCPUHaveNum:    0,
		NodeAMemoryHaveNum: 0,
		NodeBMemoryHaveNum: 0,
		SumCpuNum:          ServerTypeMap[len(ServerTypeMap)-1].CPU,
		SumMemoryNum:       ServerTypeMap[len(ServerTypeMap)-1].Memory,
		NodeA: ServerNode{
			CPU:    ServerTypeMap[len(ServerTypeMap)-1].CPU / 2,
			Memory: ServerTypeMap[len(ServerTypeMap)-1].Memory / 2,
		},
		NodeB: ServerNode{
			CPU:    ServerTypeMap[len(ServerTypeMap)-1].CPU / 2,
			Memory: ServerTypeMap[len(ServerTypeMap)-1].Memory / 2,
		},
		MoneyDay:       ServerTypeMap[len(ServerTypeMap)-1].MoneyDay,
		VisualMachines: []int{},
	}
	SurplusSeres[len(SerMap)-1] = SurplusSer

	Purchase++
	PurchaseSer[ServerTypeMap[len(ServerTypeMap)-1].Name] += 1
	RMB = ServerTypeMap[len(ServerTypeMap)-1].Money
}

