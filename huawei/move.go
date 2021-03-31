package huawei

func RemoveVisualMachine() {
	num := 5 * len(SerMap) / 1000
	for i, v1 := range SurplusSeres {
		if num < 0 {
			return
		}
		if v1.NodeACPUHaveNum > v1.SumCpuNum/2 || v1.NodeBCPUHaveNum > v1.SumCpuNum/2 ||
			v1.NodeAMemoryHaveNum > v1.SumMemoryNum/2 || v1.NodeBMemoryHaveNum > v1.SumMemoryNum/2 {
			continue
		}
		for j, v2 := range SurplusSeres {
			if i != j && fullServer(*v1, *v2) {
				v1.NodeA.CPU = v1.SumCpuNum / 2
				v1.NodeB.CPU = v1.SumCpuNum / 2
				v1.NodeA.Memory = v1.SumMemoryNum / 2
				v1.NodeB.Memory = v1.SumMemoryNum / 2

				v2.NodeA.CPU = 0
				v2.NodeB.CPU = 0
				v2.NodeA.Memory = 0
				v2.NodeB.Memory = 0
				num--
				Migration += len(v2.VisualMachines)

				for _, v := range v2.VisualMachines {
					//改变虚拟机的映射
					//fmt.Println(MigrationNum[v],MachineSer[v])
					MachineSer[v] = SerMap[v1.ServerID]
					MigrationNum[v] = v1.ServerID
				}

				v1.VisualMachines = append(v1.VisualMachines, v2.VisualMachines[:]...)
				v2.VisualMachines = []int{}
			}
		}
	}
}

func fullServer(v1, v2 SurplusServer) bool {
	if (v1.NodeA.CPU+v2.NodeACPUHaveNum == v1.SumCpuNum/2) && (v1.NodeB.CPU+v2.NodeBCPUHaveNum == v1.SumCpuNum/2) &&
		(v1.NodeA.Memory+v2.NodeAMemoryHaveNum == v1.SumMemoryNum/2) && (v1.NodeB.Memory+v2.NodeBMemoryHaveNum == v1.SumMemoryNum/2) {
		return true
	}
	return false
}
