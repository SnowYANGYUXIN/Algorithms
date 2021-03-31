package huawei

func DeleteVisualMachine(id int) {
	server := MachineSer[id]
	visual := VisualMachineTypeMap[server.VisualMachineName]
	SerMap[server.ID].CPU += visual.CPU
	SerMap[server.ID].Memory += visual.Memory

	//通过结点类型恢复结点的相应数据
	switch server.NodeName {
	case "A":
		SurplusSeres[server.ID].NodeA.CPU += visual.CPU
		SurplusSeres[server.ID].NodeA.Memory += visual.Memory
		SurplusSeres[server.ID].NodeACPUHaveNum -= visual.CPU
		SurplusSeres[server.ID].NodeAMemoryHaveNum -= visual.Memory
		//因为SurplusSe存的是SurplusSeres的地址，则不用再重新改变了
		break
	case "B":
		SurplusSeres[server.ID].NodeB.CPU += visual.CPU
		SurplusSeres[server.ID].NodeB.Memory += visual.Memory
		SurplusSeres[server.ID].NodeBCPUHaveNum -= visual.CPU
		SurplusSeres[server.ID].NodeBMemoryHaveNum -= visual.Memory
		//因为SurplusSe存的是SurplusSeres的地址，则不用再重新改变了
		break
	case "AB":
		SurplusSeres[server.ID].NodeA.CPU += visual.CPU / 2
		SurplusSeres[server.ID].NodeA.Memory += visual.Memory / 2
		SurplusSeres[server.ID].NodeB.CPU += visual.CPU / 2
		SurplusSeres[server.ID].NodeB.Memory += visual.Memory / 2
		SurplusSeres[server.ID].NodeACPUHaveNum -= visual.CPU / 2
		SurplusSeres[server.ID].NodeAMemoryHaveNum -= visual.Memory / 2
		SurplusSeres[server.ID].NodeBCPUHaveNum -= visual.CPU / 2
		SurplusSeres[server.ID].NodeBMemoryHaveNum -= visual.Memory / 2
		break
	}
	//fmt.Println(visual,*SurplusSeres[server.ID])

	//删除虚拟机id到服务器的映射
	delete(MachineSer, id)
}
