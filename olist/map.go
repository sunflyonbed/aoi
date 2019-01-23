package olist

import (
	"container/list"
	"fmt"
)

const (
	BroadCastRange = 2
)

type Map struct {
	xList, yList *list.List
	nodes        map[uint32]*Node

	moveMap  map[uint32]*Node
	enterMap map[uint32]*Node
	leaveMap map[uint32]*Node
}

func NewMap() *Map {
	return &Map{
		xList: list.New(),
		yList: list.New(),
		nodes: make(map[uint32]*Node),

		moveMap:  make(map[uint32]*Node),
		enterMap: make(map[uint32]*Node),
		leaveMap: make(map[uint32]*Node),
	}
}

func (this *Map) GetNode(id uint32) *Node { return this.nodes[id] }
func (this *Map) AddNode(id uint32, x, y int) {
	if this.GetNode(id) != nil {
		return
	}
	node := NewNode(id, x, y)
	this.nodes[id] = node
	insert := false
	inListX := make(map[uint32]bool)
	var xEl, yEl *list.Element
	for e := this.xList.Front(); e != nil; e = e.Next() {
		eNode := e.Value.(*Node)
		diff := eNode.X() - node.X()
		if abs(diff) <= BroadCastRange {
			inListX[eNode.Id()] = true
		}
		if !insert && e.Value.(*Node).X() > node.X() {
			xEl = this.xList.InsertBefore(node, e)
			insert = true
		}
		if diff > BroadCastRange {
			break
		}
	}
	if !insert {
		xEl = this.xList.PushBack(node)
	}

	insert = false
	for e := this.yList.Front(); e != nil; e = e.Next() {
		eNode := e.Value.(*Node)
		diff := eNode.Y() - node.Y()
		if abs(diff) <= BroadCastRange && inListX[eNode.Id()] {
			this.enterMap[eNode.Id()] = eNode
		}
		if !insert && e.Value.(*Node).Y() > node.Y() {
			yEl = this.yList.InsertBefore(node, e)
			insert = true
		}
		if diff > BroadCastRange {
			break
		}
	}
	if !insert {
		yEl = this.yList.PushBack(node)
	}
	node.SetXElement(xEl)
	node.SetYElement(yEl)
	this.BroadCast(node)
}

func (this *Map) MoveNode(id uint32, x, y int) {
	node := this.GetNode(id)
	if node == nil {
		return
	}

	oldMap := node.GetRangeMap()
	// fmt.Printf("get oldMap:%d %v\n", id, oldMap)
	this.ChangePosition(node, x, y)
	newMap := node.GetRangeMap()
	// fmt.Printf("get newMap:%d %v\n", id, newMap)

	for id, v := range oldMap {
		if newMap[id] != nil {
			this.moveMap[id] = v
		}
	}

	for id, v := range oldMap {
		if this.moveMap[id] == nil {
			this.leaveMap[id] = v
		}
	}

	for id, v := range newMap {
		if oldMap[id] == nil {
			this.enterMap[id] = v
		}
	}

	this.BroadCast(node)
}

func (this *Map) ChangePosition(node *Node, x, y int) {
	oldX := node.X()
	oldY := node.Y()
	node.SetX(x)
	node.SetY(y)
	originXEl := node.XElement()
	originYEl := node.YElement()
	var xEl, yEl *list.Element
	insert := false

	if x > oldX {
		for el := originXEl.Next(); el != nil; el = el.Next() {
			// fmt.Printf("ChangePosition xel next:%d %v\n", node.Id(), el.Value.(*Node))
			if el.Value.(*Node).X() > x {
				xEl = this.xList.InsertBefore(node, el)
				insert = true
				break
			}
		}
		if !insert {
			xEl = this.xList.PushBack(node)
		}
	} else if x < oldX {
		for el := originXEl.Prev(); el != nil; el = el.Prev() {
			// fmt.Printf("ChangePosition xel prev:%d %v\n", node.Id(), el.Value.(*Node))
			if el.Value.(*Node).X() < x {
				xEl = this.xList.InsertAfter(node, el)
				insert = true
				break
			}
		}
		if !insert {
			xEl = this.xList.PushFront(node)
		}
	}
	if xEl != nil {
		node.SetXElement(xEl)
		this.xList.Remove(originXEl)
	}

	if y > oldY {
		for el := originYEl.Next(); el != nil; el = el.Next() {
			// fmt.Printf("ChangePosition yel next:%d %v\n", node.Id(), el.Value.(*Node))
			if el.Value.(*Node).Y() > y {
				yEl = this.yList.InsertBefore(node, el)
				insert = true
				break
			}
		}
		if !insert {
			yEl = this.yList.PushBack(node)
		}
	} else if y < oldY {
		for el := originYEl.Prev(); el != nil; el = el.Prev() {
			// fmt.Printf("ChangePosition yel prev:%d %v\n", node.Id(), el.Value.(*Node))
			if el.Value.(*Node).Y() < y {
				yEl = this.yList.InsertAfter(node, el)
				insert = true
				break
			}
		}
		if !insert {
			yEl = this.yList.PushFront(node)
		}
	}
	if yEl != nil {
		node.SetYElement(yEl)
		this.yList.Remove(originYEl)
	}
}

func (this *Map) BroadCast(node *Node) {
	for _, v := range this.moveMap {
		fmt.Printf("Node:%d see move:%d\n", v.Id(), node.Id())
	}
	for _, v := range this.enterMap {
		fmt.Printf("Node:%d see enter:%d\n", v.Id(), node.Id())
	}
	for _, v := range this.leaveMap {
		fmt.Printf("Node:%d see leave:%d\n", v.Id(), node.Id())
	}
	this.moveMap = make(map[uint32]*Node)
	this.enterMap = make(map[uint32]*Node)
	this.leaveMap = make(map[uint32]*Node)
}

func (this *Map) LeaveNode(id uint32) {
	node := this.GetNode(id)
	if node == nil {
		return
	}
	this.leaveMap = node.GetRangeMap()
	this.xList.Remove(node.XElement())
	this.yList.Remove(node.YElement())
	this.BroadCast(node)
	delete(this.nodes, id)
}

func (this *Map) PrintAOI() {
	for _, v := range this.nodes {
		fmt.Printf("printAOI:%d %v\n", v.Id(), v.GetRangeMap())
	}
}
