package BinarySearchTree

// BST data structure
type BST struct {
	value int
	left  *BST
	right *BST
}

// NewBST - Binary Search Tree constructor
func NewBST(value int, left *BST, right *BST) *BST {
	return &BST{value, left, right}
}

// Traverse - walking the tree
func (t *BST) Traverse(cb func(t *BST)) {
	cb(t)
	if t.left != nil {
		t.left.Traverse(cb)
	}
	if t.right != nil {
		t.right.Traverse(cb)
	}
}

// Find children
func (t *BST) Find(value int) *BST {
	if t == nil {
		return nil
	}
	if t.value == value {
		return t
	}
	if t.value > value {
		return t.left.Find(value)
	}
	if t.value < value {
		return t.right.Find(value)
	}
	return nil
}

// AddLeft - add left children
func (t *BST) AddLeft(c *BST) {
	t.left = c
}

// AddRight - add right children
func (t *BST) AddRight(c *BST) {
	t.right = c
}

// Add - add children
func (t *BST) Add(c *BST) {
	if c.value > t.value {
		if t.right == nil {
			t.AddRight(c)
			return
		}
		t.right.Add(c)
	} else if c.value < t.value {
		if t.left == nil {
			t.AddLeft(c)
			return
		}
		t.left.Add(c)
	}
	return
}
