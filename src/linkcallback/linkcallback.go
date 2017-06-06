package linkcallback

import (
	"sync"
	"fmt"
	"unsafe"
)

const (
	SUCCESS=0
	FAILURE=-1
)

type LinkTableNode struct {
	PNext *LinkTableNode
}

type LinkTable struct {
	PHead *LinkTableNode
	PTail *LinkTableNode
	SumOfNode int
	LMutex sync.Mutex
}

/*
 * Create a LinkTable
 */
func CreateLinkTable() *LinkTable {
	var pLinkTable *LinkTable = new(LinkTable)
	return pLinkTable
}

/*
 * Delete a LinkTable
 */
func DeleteLinkTable(pLinkTable *LinkTable) int {
	if pLinkTable == nil {
		fmt.Println("It is an empty Linktable!")
		return FAILURE
	}
	if pLinkTable.PHead == nil {
		return FAILURE
	}
	for pLinkTable.PHead != pLinkTable.PTail {
		pLinkTable.PHead = pLinkTable.PHead.PNext
	}
	return SUCCESS
}

/*
 * Add a LinkTableNode to LinkTable
 */
func AddLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
	if pLinkTable == nil || pNode == nil {
		fmt.Println("Don't exist this NodeTable or Node!")
		return FAILURE
	}
	if pLinkTable.PHead == nil {
		pLinkTable.PHead = pNode	
	} else {
		pLinkTable.PTail.PNext = pNode
	}
	pLinkTable.PTail = pNode
	pLinkTable.SumOfNode +=1
	return SUCCESS
}

/*
 * Delete a LinkTableNode from LinkTable
 */
func DeleteLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
	if pLinkTable == nil || pLinkTable.PHead ==nil || pNode ==  nil {
		fmt.Println("Don't exist this NodeTable or Node!")
		return FAILURE
	}
	var ptrNode *LinkTableNode = GetLinkTableHead(pLinkTable)
	if ptrNode == pNode {
		pLinkTable.PHead = ptrNode.PNext
		pLinkTable.SumOfNode -= 1
		return SUCCESS
	}
	for ptrNode.PNext != nil {
		if ptrNode.PNext == pNode {
			ptrNode.PNext = ptrNode.PNext.PNext
			pLinkTable.SumOfNode -= 1
			if pLinkTable.SumOfNode == 0 {
				pLinkTable.PTail = nil
			}
			return SUCCESS
		}
		ptrNode = ptrNode.PNext
	}
	return FAILURE
}

/*
 * Search a LinkTableNode from LinkTable
 * int Condition(pNode *LinkTableNode, agrs ptr);
 */
func SearchLinkTableNode(Condition func(*LinkTableNode, unsafe.Pointer) int, pLinkTable *LinkTable, agrs unsafe.Pointer) *LinkTableNode {
	if pLinkTable == nil || Condition == nil {
		return nil
	}
	pNode := pLinkTable.PHead
	for pNode != nil {
		if Condition(pNode, agrs)== SUCCESS {
			return pNode
		}
		pNode = pNode.PNext
	}
	return nil 
}

/*
 * get LinkTableHead
 */
func GetLinkTableHead(pLinkTable *LinkTable) *LinkTableNode {
	if pLinkTable.PHead == nil {
		fmt.Println("don't exist HeadNode!")
		return nil
	}
	return pLinkTable.PHead
}

/*
 * get next LinkTableNode
 */
func GetNextLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) *LinkTableNode {
	if pLinkTable == nil || pLinkTable.PHead == nil {
		fmt.Println("Do not exist an NodeTable!")
		return nil
	}
	if pNode == nil {
		fmt.Println("Node is nil")
		return nil
	}
	var phead *LinkTableNode = GetLinkTableHead(pLinkTable)
	for phead != nil {
		if phead == pNode {
			return phead.PNext
		}
		phead = phead.PNext
	}
	return nil
}