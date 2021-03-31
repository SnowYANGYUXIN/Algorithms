package huawei

import "sync"

type Service struct {
	ID       int
	Name     string
	CPU      int
	Memory   int
	Money    int
	MoneyDay int

	VisualMachineName string //用于在删除虚拟机时知道服务器上该虚拟机的类型
	NodeName          string //部署的结点类型 有A、B、AB
}

type ServerNode struct {
	CPU    int
	Memory int
}

type VisualMachine struct {
	ID       int
	Name     string
	CPU      int
	Memory   int
	isDouble int
}

type SurplusServer struct {
	ServerID           int
	NodeACPUHaveNum    int //结点A的cpu使用数
	NodeBCPUHaveNum    int //结点B的cpu使用数
	NodeAMemoryHaveNum int //结点A的内存使用数
	NodeBMemoryHaveNum int //结点B的内存使用数
	SumCpuNum          int
	SumMemoryNum       int
	MoneyDay           int //每天的钱
	NodeA              ServerNode
	NodeB              ServerNode
	VisualMachines     []int
}

var ServerTypeID int    //服务器类型ID
var VisualMachineID int //虚拟机类型ID

var MachineSer map[int]*Service //根据虚拟机ID号映射到服务器号

var ServerTypeMap []*Service                      //初始化服务器类型的的map
var VisualMachineTypeMap map[string]VisualMachine //初始化服务器类型的的map

//var Day int //工作天数
var wg sync.WaitGroup

var SerMap []*Service //购买服务器map

var SurplusSeres map[int]*SurplusServer //通过服务器ID查询当前服务器剩余CUP和内存 得是map，因为撤销虚拟机后将空闲cup和内存
var SurplusSer *SurplusServer           //取SurplusSer的地址 则对SurplusSer的操作就是对SurplusSeres的操作，反之亦然

var Purchase int    //每日购买数
var Record []string //购买记录
var Migration int   //虚拟机迁移数
var RMB int

var PurchaseSer map[string]int
var MigrationNum map[int]int
