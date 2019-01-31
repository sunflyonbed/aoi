package tower

//Node is double LinkNode
type Node struct {
	MyNode
	x, y int
	t    *Tower
}

func NewNode(node MyNode, x, y int) *Node {
	return &Node{
		MyNode: node,
		x:      x,
		y:      y,
	}
}

func (this *Node) Id() uint64        { return this.UniqueId() }
func (this *Node) X() int            { return this.x }
func (this *Node) Y() int            { return this.y }
func (this *Node) SetX(x int)        { this.x = x }
func (this *Node) SetY(y int)        { this.y = y }
func (this *Node) SetTower(t *Tower) { this.t = t }
func (this *Node) Tower() *Tower     { return this.t }
