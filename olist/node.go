package olist

import (
	"container/list"
)

//Node is double LinkNode
type Node struct {
	MyNode
	x, y     int
	xEl, yEl *list.Element
}

func NewNode(node MyNode, x, y int) *Node {
	return &Node{
		MyNode: node,
		x:      x,
		y:      y,
	}
}

func (this *Node) Id() uint64 { return this.UniqueId() }
func (this *Node) X() int     { return this.x }
func (this *Node) Y() int     { return this.y }
func (this *Node) SetX(x int) { this.x = x }
func (this *Node) SetY(y int) { this.y = y }

func (this *Node) SetXElement(el *list.Element) { this.xEl = el }
func (this *Node) SetYElement(el *list.Element) { this.yEl = el }
func (this *Node) XElement() *list.Element      { return this.xEl }
func (this *Node) YElement() *list.Element      { return this.yEl }

func (this *Node) ClearElement() {
	this.xEl = nil
	this.yEl = nil
}

func (this *Node) GetRangeMap(rangeSize int) map[uint64]*Node {
	if this.xEl == nil || this.yEl == nil {
		return nil
	}
	inListX := make(map[uint64]bool)
	result := make(map[uint64]*Node)
	for e := this.xEl.Prev(); e != nil; e = e.Prev() {
		eNode := e.Value.(*Node)
		// fmt.Printf("check xel prev enode:%v\n", eNode)
		if eNode.X()+rangeSize >= this.X() {
			inListX[eNode.Id()] = true
		} else {
			break
		}
	}
	for e := this.xEl.Next(); e != nil; e = e.Next() {
		eNode := e.Value.(*Node)
		// fmt.Printf("check xel next enode:%v\n", eNode)
		if eNode.X() <= this.X()+rangeSize {
			inListX[eNode.Id()] = true
		} else {
			break
		}
	}

	for e := this.yEl.Prev(); e != nil; e = e.Prev() {
		eNode := e.Value.(*Node)
		// fmt.Printf("check yel prev enode:%v\n", eNode)
		if eNode.Y()+rangeSize >= this.Y() {
			if inListX[eNode.Id()] {
				result[eNode.Id()] = eNode
			}
		} else {
			break
		}
	}

	for e := this.yEl.Next(); e != nil; e = e.Next() {
		eNode := e.Value.(*Node)
		// fmt.Printf("check yel next enode:%v\n", eNode)
		if eNode.Y() <= this.Y()+rangeSize {
			if inListX[eNode.Id()] {
				result[eNode.Id()] = eNode
			}
		} else {
			break
		}
	}

	return result
}
