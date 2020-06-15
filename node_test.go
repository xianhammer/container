package container

import "testing"

func TestNodeInsert0(t *testing.T) {
	n0 := NewNode(0)
	n1 := NewNode(1)
	n2 := NewNode(2)

	n1.InsertAfter(n0)
	n2.InsertAfter(n1)

	if n0.Next() != n1 {
		t.Errorf("Node 0 next was not node 1.\n")
	}
	if n1.Next() != n2 {
		t.Errorf("Node 1 next was not node 2.\n")
	}
	if n2.Next() != nil {
		t.Errorf("Node 2 next was not nil.\n")
	}
}

func TestNodeInsert1(t *testing.T) {
	n0 := NewNode(0)
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	n3.InsertAfter(n0)
	n2.InsertAfter(n0)
	n1.InsertAfter(n0)

	if n0.Previous() != nil {
		t.Errorf("Node 0 prev was not nil.\n")
	}
	if n1.Previous() != n0 {
		t.Errorf("Node 1 prev was not node 0.\n")
	}
	if n2.Previous() != n1 {
		t.Errorf("Node 2 prev was not node 1.\n")
	}
	if n3.Previous() != n2 {
		t.Errorf("Node 3 prev was not node 2.\n")
	}
}

func TestNodeRemove(t *testing.T) {
	n0 := NewNode(1)
	n1 := NewNode(2)
	n2 := NewNode(3)

	n1.InsertAfter(n0)
	n2.InsertAfter(n1)

	n1.Remove()
	if n0.Next() != n2 {
		t.Errorf("Node 0 next was not node 2.\n")
	}
	if n2.Previous() != n0 {
		t.Errorf("Node 2 previous was not node 0.\n")
	}

	if n1.Next() != nil {
		t.Errorf("Node 1 next was not nil.\n")
	}
	if n1.Previous() != nil {
		t.Errorf("Node 1 previous was not nil.\n")
	}

	n1.Remove() // Should yield no changes
	if n0.Next() != n2 {
		t.Errorf("Node 0 next was not node 2.\n")
	}
	if n2.Previous() != n0 {
		t.Errorf("Node 2 previous was not node 0.\n")
	}

	if n1.Next() != nil {
		t.Errorf("Node 1 next was not nil.\n")
	}
	if n1.Previous() != nil {
		t.Errorf("Node 1 previous was not nil.\n")
	}

	n0.Remove() // Now all nodes should be detached
	if n0.Next() != nil {
		t.Errorf("Node 0 next was not nil.\n")
	}
	if n0.Previous() != nil {
		t.Errorf("Node 0 previous was not nil.\n")
	}

	if n1.Next() != nil {
		t.Errorf("Node 1 next was not nil.\n")
	}
	if n1.Previous() != nil {
		t.Errorf("Node 1 previous was not nil.\n")
	}

	if n2.Next() != nil {
		t.Errorf("Node 0 next was not nil.\n")
	}
	if n2.Previous() != nil {
		t.Errorf("Node 0 previous was not nil.\n")
	}
}
