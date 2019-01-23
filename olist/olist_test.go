package olist

import (
	"fmt"
	"testing"
)

func NewTestNode(id uint64) *TestNode {
	return &TestNode{
		Id: id,
	}
}

type TestNode struct {
	Id uint64
}

func (this *TestNode) UniqueId() uint64 { return this.Id }
func (this *TestNode) OnEnter(id uint64) {
	fmt.Printf("Node:%d see move:%d\n", this.Id, id)
}
func (this *TestNode) OnLeave(id uint64) {
	fmt.Printf("Node:%d see leave:%d\n", this.Id, id)
}
func (this *TestNode) OnMove(id uint64) {
	fmt.Printf("Node:%d see move:%d\n", this.Id, id)
}

/*
	6						2
	5	1
	4
	3					5
	2		4
	1			3
		1	2	3	4	5	6
*/
func TestOlist(t *testing.T) {
	scene := NewMap()
	scene.AddNode(NewTestNode(1), 1, 5)
	scene.AddNode(NewTestNode(2), 6, 6)
	scene.AddNode(NewTestNode(3), 3, 1)
	scene.AddNode(NewTestNode(4), 2, 2)
	scene.AddNode(NewTestNode(5), 5, 3)
	scene.PrintAOI()

	scene.MoveNode(5, 6, 3)
	scene.MoveNode(1, 1, 1)
	scene.MoveNode(4, 2, 1)
	scene.PrintAOI()
}
