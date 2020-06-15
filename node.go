package container

type Node struct {
	next, prev *Node
	Value      interface{}
}

// NewNode create a new node, with preset value
func NewNode(value interface{}) *Node {
	return &Node{nil, nil, value}
}

// Next node
func (n *Node) Next() *Node {
	return n.next
}

// Previous node
func (n *Node) Previous() *Node {
	return n.prev
}

// InsertAfter will insert this node after the target.
func (n *Node) InsertAfter(target *Node) {
	if target.next != nil {
		n.next = target.next
		target.next.prev = n
	}
	target.next = n
	n.prev = target
}

// Remove this node from the list
func (n *Node) Remove() {
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	n.next = nil
	n.prev = nil
}
