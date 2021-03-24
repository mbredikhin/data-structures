package Tree

// Tree data structure
type Tree struct {
	value    int
	children []*Tree
}

// NewTree - Tree constructor
func NewTree(value int, children []*Tree) *Tree {
	return &Tree{value, children}
}

// Traverse - walking the tree
func (t *Tree) Traverse(cb func(t *Tree)) {
	cb(t)
	for _, c := range t.children {
		c.Traverse(cb)
	}
}

// Add children
func (t *Tree) Add(c *Tree) {
	t.children = append(t.children, c)
}

// Find children
func (t *Tree) Find(value int) *Tree {
	if t.value == value {
		return t
	}
	if len(t.children) == 0 {
		return nil
	}
	for _, c := range t.children {
		c.Find(value)
	}
	return nil
}
