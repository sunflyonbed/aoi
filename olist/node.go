package olist

import (
	"container/list"
)

//Node is double LinkNode
type Node struct {
	x, y     int
	id       uint32
	xEl, yEl *list.Element
}

func NewNode(id uint32, x, y int) *Node {
	return &Node{
		id: id,
		x:  x,
		y:  y,
	}
}

func (this *Node) Id() uint32 { return this.id }
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

func (this *Node) GetRangeMap() map[uint32]*Node {
	if this.xEl == nil || this.yEl == nil {
		return nil
	}
	inListX := make(map[uint32]bool)
	result := make(map[uint32]*Node)
	for e := this.xEl.Prev(); e != nil; e = e.Prev() {
		eNode := e.Value.(*Node)
		// fmt.Printf("check xel prev enode:%v\n", eNode)
		if eNode.X()+BroadCastRange >= this.X() {
			inListX[eNode.Id()] = true
		} else {
			break
		}
	}
	for e := this.xEl.Next(); e != nil; e = e.Next() {
		eNode := e.Value.(*Node)
		// fmt.Printf("check xel next enode:%v\n", eNode)
		if eNode.X() <= this.X()+BroadCastRange {
			inListX[eNode.Id()] = true
		} else {
			break
		}
	}

	for e := this.yEl.Prev(); e != nil; e = e.Prev() {
		eNode := e.Value.(*Node)
		// fmt.Printf("check yel prev enode:%v\n", eNode)
		if eNode.Y()+BroadCastRange >= this.Y() {
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
		if eNode.Y() <= this.Y()+BroadCastRange {
			if inListX[eNode.Id()] {
				result[eNode.Id()] = eNode
			}
		} else {
			break
		}
	}

	return result
}
