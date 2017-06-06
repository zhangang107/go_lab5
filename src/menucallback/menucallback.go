package main

import "fmt"
import "os"
import "os/exec"
import(. "linkcallback")
import "unsafe"


var head *LinkTable 
func help() int {
	ShowAllCmd(head)
	return 0
}
func quit() int {
	os.Exit(1)
	return 0
}
func add() int {
	var x int
	var y int
	fmt.Println("pls input two int numbers")
	fmt.Scanf("%d %d",&x,&y)
	ret := x+y
	fmt.Printf("The result :%d\n",ret)
	return 0
}

func sub() int {
	var x, y int
	fmt.Println("pls input two int numbers")
	fmt.Scanf("%d %d",&x,&y)
	ret := x-y
	fmt.Printf("The result :%d\n",ret)
	return 0
}

func newfile() int {
	var cmdstr, filename string
	cmdstr = "touch"
	fmt.Println("Pls input filename")
	fmt.Scanf("%s",&filename)
	cmd :=exec.Command(cmdstr,filename)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if string(out) != "" {
		fmt.Println(string(out))
	}
	fmt.Printf("%s success!",cmdstr)
	return 0
}

func del() int {
	var cmdstr, filename string
	cmdstr = "rm"
	fmt.Println("Pls input filename")
	fmt.Scanf("%s",&filename)
	cmd :=exec.Command(cmdstr,filename)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if string(out) != "" {
		fmt.Println(string(out))
	}
	fmt.Printf("%s success!",cmdstr)
	return 0
}

func pwd() int {
	cmd :=exec.Command("pwd")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if string(out) != "" {
		fmt.Println(string(out))
	}
	return 0
}

func ls() int {
	cmd :=exec.Command("ls")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if string(out) != "" {
		fmt.Println(string(out))
	}
	return 0
}

type DataNode struct {
	pNext *LinkTable
	cmd string
	desc string
	handler func() int
}

func SearchCondition(pLinkTableNode *LinkTableNode, args unsafe.Pointer) int {
	pcmd := (*string)(unsafe.Pointer(args))
	dNode := (*DataNode)(unsafe.Pointer(pLinkTableNode))
	if dNode.cmd == *pcmd {
		return SUCCESS
	}
	return FAILURE
}

func FindCmd(head *LinkTable, cmd string) *DataNode {
	//tNode := GetLinkTableHead(head)
	tNode:= SearchLinkTableNode(SearchCondition, head, unsafe.Pointer(&cmd))
	return (*DataNode)(unsafe.Pointer(tNode))
}

func ShowAllCmd(head *LinkTable) int {
	tNode := GetLinkTableHead(head)
	fmt.Printf("********Menu List:**********\n")
	for tNode != nil {
		dNode := (*DataNode)(unsafe.Pointer(tNode))
		fmt.Printf("%s - %s\n",dNode.cmd, dNode.desc)
		tNode = GetNextLinkTableNode(head,tNode)
	}
	fmt.Printf("****************************\n")
	return 0
}

func InitMenuData() *LinkTable {
	var pLinkTable = CreateLinkTable();
	var pNode = new(DataNode)
	pNode.cmd = "help"
	pNode.desc = "Show The Command List"
	pNode.handler=help
	AddLinkTableNode(pLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode)))
	var pNode1 *DataNode = new(DataNode)
	pNode1.cmd = "add"
	pNode1.desc = "Add Two Numbers"
	pNode1.handler=add
	AddLinkTableNode(pLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode1)))
	var pNode2 *DataNode = new(DataNode)
	pNode2.cmd = "sub"
	pNode2.desc = "Minus Two Numbers"
	pNode2.handler=sub
	AddLinkTableNode(pLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode2)))
	var pNode3 *DataNode = new(DataNode)
	pNode3.cmd = "new"
	pNode3.desc = "New File"
	pNode3.handler=newfile
	AddLinkTableNode(pLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode3)))
	var pNode4 *DataNode = new(DataNode)
	pNode4.cmd = "del"
	pNode4.desc = "Delete File"
	pNode4.handler=del
	AddLinkTableNode(pLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode4)))
	var pNode5 *DataNode = new(DataNode)
	pNode5.cmd = "pwd"
	pNode5.desc = "Show The path"
	pNode5.handler=pwd
	AddLinkTableNode(pLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode5)))
	var pNode6 *DataNode = new(DataNode)
	pNode6.cmd = "ls"
	pNode6.desc = "Show Files"
	pNode6.handler=ls
	AddLinkTableNode(pLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode6)))
	var pNode7 *DataNode = new(DataNode)
	pNode7.cmd = "quit"
	pNode7.desc = "Quit The Menu!"
	pNode7.handler=quit
	AddLinkTableNode(pLinkTable,(*LinkTableNode)(unsafe.Pointer(pNode7)))
	return pLinkTable
}



func main() {
	head = InitMenuData()
	for true {
		var cmd string
		fmt.Print("Command->")
		fmt.Scanf("%s",&cmd)
		p := FindCmd(head,cmd)
		if p == nil {
			fmt.Println("Command not exit,you can try help")
			continue
		}
		fmt.Printf("%s - %s\n",p.cmd,p.desc)
		if p.handler != nil {
			p.handler()
		}
	}
}