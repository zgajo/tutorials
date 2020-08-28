/*Package mojtest b-tree
- balanced m-way tree (m way is order)
**************************************   maximum degree   *********************************
- order (m): 5 - how much pointers can be in node
- all data is sorted
- every leaf node must be in same leaf node
- every node has max m children (leafs)
- min children: leaf - 0, root: 2, internal node: cieling od m/2 (5/2 = 2.5 = cieling 3)
- max keys in node are m-1
- min keys in root 1, all other nodes: cieling(m/2)-1
**************************************   --------------   *********************************
**************************************   minimum degree   *********************************
- max keys in node are 2*m-1
- min keys in root 1, all other nodes: m-1
- order (m): 5 - how much min pointers can be in node
- min children: leaf - 0, root: 2, internal node: m
- max num of children in node: equal to the number of keys in it plus 1.
- all data is sorted
- every leaf node must be in same leaf node
- every node has max m children (leafs)
**************************************   --------------   *********************************

- insertion is always done in leaf node
- while inserting, sort has to be done
*/
package mojtest

import "fmt"

/**

// INSERT KEY
// check if root is null
	// create new node and set it as root
// check if root is leaf
	// if root is not full insert key into root node
	// else
		// create node left[:middle] and node right [middle+1:]
		// create new node with one key [middle] from previous root and set previous created left and right nodes as his two children
// if root is not leaf
	// search for index of last index key which is larger then key which is going to be inserted
	// use found index key as index of children array node
	// repeat this step for every nested child until last node is leaf
		// if node is not full insert key into node and sort node
		// else
			// insert and sort into array
			// split array
			// create node left[:middle] and node right [middle+1:]
			// create new node with one key [middle] from previous root and set previous created left and right nodes as his two children
			// insert new node (only key) into parent, sort parent keys array
			// set his children as key index position and key index position +1 in childrens array

*/

// Node :
type Node struct {
	Keys     []int   // An array of keys
	Children []*Node // An array of child pointers
	Parent   *Node
}

// CreateNode : Creates new node
func CreateNode() *Node {
	return &Node{
		Keys:     []int{},
		Children: []*Node{},
	}
}

func (node *Node) numOfKeys() int {
	return len(node.Keys)
}

func (node *Node) numOfChildren() int {
	return len(node.Children)
}

// isLeaf : is true when node is leaf. Otherwise false
func (node *Node) isLeaf() bool {
	return len(node.Children) == 0
}

// BTree :
type BTree struct {
	Root      *Node // Pointer to root node
	MinDegree int   // Minimum number of children
	MaxDegree int   // Max num of children
}

// InitBtree : Creates new btree
func InitBtree(m int) *BTree {
	if m*2 < 2 {
		panic("Invalid order, should be at least 2")
	}

	return &BTree{
		MaxDegree: m * 2,
		MinDegree: m,
	}
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *BTree) Get(key int) (value interface{}, found bool) {
	node, index, found := tree.searchRecursively(tree.Root, key)
	if found {
		return node.Keys[index], true
	}
	return nil, false
}

// Put inserts key-value pair node into the tree.
// If key already exists, then its value is updated with the new value.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *BTree) Put(key int) {

	if tree.Root == nil {
		tree.Root = &Node{Keys: []int{key}, Children: []*Node{}}
		// tree.size++
		return
	}

	if tree.insert(tree.Root, key) {
		// tree.size++
	}
}

func (tree *BTree) insert(node *Node, key int) (inserted bool) {
	if tree.isLeaf(node) {
		return tree.insertIntoLeaf(node, key)
	}
	return tree.insertIntoInternal(node, key)
}

func (tree *BTree) insertIntoLeaf(node *Node, key int) (inserted bool) {
	insertPosition, found := tree.search(node, key)
	if found {
		node.Keys[insertPosition] = key
		return false
	}
	// Insert entry's key in the middle of the node
	node.Keys = append(node.Keys, 0)
	copy(node.Keys[insertPosition+1:], node.Keys[insertPosition:])
	node.Keys[insertPosition] = key
	tree.split(node)
	return true
}

func (tree *BTree) insertIntoInternal(node *Node, key int) (inserted bool) {
	insertPosition, found := tree.search(node, key)
	if found {
		node.Keys[insertPosition] = key
		return false
	}
	return tree.insert(node.Children[insertPosition], key)
}

func (tree *BTree) split(node *Node) {
	if !tree.shouldSplit(node) {
		return
	}

	if node == tree.Root {
		tree.splitRoot()
		return
	}

	tree.splitNonRoot(node)
}

func (tree *BTree) splitNonRoot(node *Node) {
	middle := tree.middle()
	parent := node.Parent

	left := &Node{Keys: append([]int(nil), node.Keys[:middle]...), Parent: parent}
	right := &Node{Keys: append([]int(nil), node.Keys[middle+1:]...), Parent: parent}

	// Move children from the node to be split into left and right nodes
	if !tree.isLeaf(node) {
		left.Children = append([]*Node(nil), node.Children[:middle+1]...)
		right.Children = append([]*Node(nil), node.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	insertPosition, _ := tree.search(parent, node.Keys[middle])

	// Insert middle key into parent
	parent.Keys = append(parent.Keys, 0)
	copy(parent.Keys[insertPosition+1:], parent.Keys[insertPosition:])
	parent.Keys[insertPosition] = node.Keys[middle]

	// Set child left of inserted key in parent to the created left node
	parent.Children[insertPosition] = left

	// Set child right of inserted key in parent to the created right node
	parent.Children = append(parent.Children, nil)
	copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
	parent.Children[insertPosition+1] = right

	tree.split(parent)
}

func (tree *BTree) splitRoot() {
	middle := tree.middle()

	left := &Node{Keys: append([]int(nil), tree.Root.Keys[:middle]...)}
	right := &Node{Keys: append([]int(nil), tree.Root.Keys[middle+1:]...)}

	// Move children from the node to be split into left and right nodes
	if !tree.isLeaf(tree.Root) {
		left.Children = append([]*Node(nil), tree.Root.Children[:middle+1]...)
		right.Children = append([]*Node(nil), tree.Root.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	// Root is a node with one entry and two children (left and right)
	newRoot := &Node{
		Keys:     []int{tree.Root.Keys[middle]},
		Children: []*Node{left, right},
	}

	left.Parent = newRoot
	right.Parent = newRoot
	tree.Root = newRoot
}

func setParent(nodes []*Node, parent *Node) {
	for _, node := range nodes {
		node.Parent = parent
	}
}

// Right :
func (tree *BTree) Right() {
	i := tree.Root
	fmt.Println("************** Right **************")
	fmt.Println("Root", i)

	for i.numOfChildren() != 0 {
		i = i.Children[i.numOfChildren()-1]
		fmt.Println("Node", i)
		for _, c := range i.Children {

			fmt.Println("child", c)
		}
	}

	fmt.Println("Leaf", i)

	fmt.Println(i.Parent)
	fmt.Println(i.Parent.Parent)
	fmt.Println("**************")
}

// Left :
func (tree *BTree) Left() {
	fmt.Println("************** Left **************")
	i := tree.Root
	fmt.Println("Root", i)
	for i.numOfChildren() != 0 {
		i = i.Children[0]
		fmt.Println("Node", i)
		for _, c := range i.Children {

			fmt.Println("child", c)
		}
	}

	fmt.Println("Leaf", i)
	fmt.Println(i.Parent)
	fmt.Println(i.Parent.Parent)
	fmt.Println("**************")
}

// Empty returns true if tree does not contain any nodes
func (tree *BTree) Empty() bool {
	return len(tree.Root.Keys) == 0
}

func (tree *BTree) isLeaf(node *Node) bool {
	return len(node.Children) == 0
}

func (tree *BTree) isFull(node *Node) bool {
	return len(node.Keys) == tree.maxEntries()
}

func (tree *BTree) shouldSplit(node *Node) bool {
	return len(node.Keys) > tree.maxEntries()
}

func (tree *BTree) maxChildren() int {
	return tree.MaxDegree - 1
}

func (tree *BTree) minChildren() int {
	return tree.MaxDegree
}

func (tree *BTree) maxEntries() int {
	return tree.maxChildren() - 1
}

func (tree *BTree) minEntries() int {
	return tree.minChildren() - 1
}

func (tree *BTree) middle() int {
	return tree.MinDegree - 1 // "-1" to favor right nodes to have more keys when splitting
}

// search searches only within the single node among its entries
func (tree *BTree) search(node *Node, key int) (index int, found bool) {
	low, high := 0, len(node.Keys)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		compare := tree.IntComparator(key, node.Keys[mid])
		switch {
		case compare > 0:
			low = mid + 1
		case compare < 0:
			high = mid - 1
		case compare == 0:
			return mid, true
		}
	}
	return low, false
}

// searchRecursively searches recursively down the tree starting at the startNode
func (tree *BTree) searchRecursively(startNode *Node, key int) (node *Node, index int, found bool) {
	if tree.Empty() {
		return nil, -1, false
	}
	node = startNode
	for {
		index, found = tree.search(node, key)
		if found {
			return node, index, true
		}
		if tree.isLeaf(node) {
			return nil, -1, false
		}
		node = node.Children[index]
	}
}

// IntComparator provides a basic comparison on int
func (tree *BTree) IntComparator(a int, b int) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

/**

// INSERT KEY
// 👍 check if root is null
	// 👍 create new node and set it as root
// 👍 check if root is leaf
	// 👍 if root is not full, sort and insert key into root node
	// 👍 else
		// 👍 create node left[:middle] and node right [middle+1:]
		// 👍 create new node with one key [middle] from previous root and set previous created left and right nodes as his two children

// if root is not leaf
	// search for index of last index key which is larger then key which is going to be inserted
	// use found index key as index of children array node
	// repeat this step for every nested child until last node is leaf
		// if node is not full insert key into node and sort node
		// else
			// insert and sort into array
			// split array
			// create node left[:middle] and node right [middle+1:]
			// create new node with one key [middle] from previous root and set previous created left and right nodes as his two children
			// insert new node (only key) into parent, sort parent keys array
			// set his children as key index position and key index position +1 in childrens array

*/
// // InsertKey : insert the key
// func (tree *BTree) InsertKey(key int) {
// 	if tree.Root == nil {
// 		tree.Root = CreateNode(tree.MaxDegree)
// 		tree.Root.Keys = append(tree.Root.Keys, key)
// 		return
// 	}

// 	tree.insert(tree.Root, key)
// }

// func (tree *BTree) insert(node *Node, key int) {
// 	if node.isLeaf() {
// 		tree.insertIntoLeaf(node, key)
// 		return
// 	}

// 	tree.insertIntoInternal(node, key)
// }

// func (tree *BTree) insertIntoInternal(node *Node, key int) {
// 	i := node.numOfKeys() - 1

// 	// The following loop does two things
// 	// a) Finds the location of new key to be inserted
// 	// b) Moves all greater keys to one place ahead
// 	for i >= 0 && node.Keys[i] > key {
// 		i--
// 	}
// 	// check if node is full
// 	if node.numOfKeys() == tree.MaxDegree-1 {
// 		// Insert entry's key in the middle of the node
// 		tree.split(node)
// 		return
// 	}

// 	tree.insert(&node.Children[i+1], key)

// }

// func (tree *BTree) insertIntoLeaf(node *Node, key int) {
// 	// check if node is full
// 	if node.numOfKeys() == tree.MaxDegree-1 {
// 		// Insert entry's key in the middle of the node
// 		node.sortAndInsert(key)

// 		tree.split(node)
// 		return
// 	}

// 	node.insertNonFull(key)

// }

// func (tree *BTree) split(node *Node) {
// 	if node == tree.Root {
// 		tree.splitRoot()
// 		return
// 	}

// 	// tree.splitNonRoot(node)
// }

// // insertNonFull : insert the key when non full condition
// func (node *Node) insertNonFull(key int) {
// 	node.sortAndInsert(key)
// }

// func (node *Node) sortAndInsert(key int) {
// 	// Initialize index as index of rightmost element
// 	i := node.numOfKeys() - 1
// 	node.Keys = append(node.Keys, 0)

// 	// The following loop does two things
// 	// a) Finds the location of new key to be inserted
// 	// b) Moves all greater keys to one place ahead
// 	for i >= 0 && node.Keys[i] > key {
// 		node.Keys[i+1] = node.Keys[i]
// 		i--
// 	}

// 	node.Keys[i+1] = key
// }

// func (tree *BTree) splitNonRoot(node *Node) {
// 	middle := tree.MinDegree

// 	left := &Node{Keys: append([]int(nil), tree.Root.Keys[:middle]...)}
// 	right := &Node{Keys: append([]int(nil), tree.Root.Keys[middle+1:]...)}

// 	// Move children from the node to be split into left and right nodes
// 	if !node.isLeaf() {
// 		left.Children = append([]Node(nil), node.Children[:middle+1]...)
// 		right.Children = append([]Node(nil), node.Children[middle+1:]...)
// 	}

// 	i := node.numOfKeys() - 1

// 	// The following loop does two things
// 	// a) Finds the location of new key to be inserted
// 	// b) Moves all greater keys to one place ahead
// 	for i >= 0 && node.Keys[i] > node.Keys[mid] {
// 		i--
// 	}

// 	insertPosition, _ := tree.search(parent, node.Entries[middle].Key)

// 	// Insert middle key into parent
// 	parent.Entries = append(parent.K, nil)
// 	copy(parent.Entries[insertPosition+1:], parent.Entries[insertPosition:])
// 	parent.Entries[insertPosition] = node.Entries[middle]

// 	// Set child left of inserted key in parent to the created left node
// 	parent.Children[insertPosition] = left

// 	// Set child right of inserted key in parent to the created right node
// 	parent.Children = append(parent.Children, nil)
// 	copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
// 	parent.Children[insertPosition+1] = right

// 	tree.split(parent)
// }

// func (tree *BTree) splitRoot() {
// 	middle := tree.MinDegree

// 	left := &Node{Keys: append([]int(nil), tree.Root.Keys[:middle]...)}
// 	right := &Node{Keys: append([]int(nil), tree.Root.Keys[middle+1:]...)}

// 	// Move children from the node to be split into left and right nodes
// 	if !tree.Root.isLeaf() {
// 		left.Children = append([]Node(nil), tree.Root.Children[:middle+1]...)
// 		right.Children = append([]Node(nil), tree.Root.Children[middle+1:]...)
// 	}

// 	// Root is a node with one entry and two children (left and right)
// 	newRoot := &Node{
// 		Keys:     []int{tree.Root.Keys[middle]},
// 		Children: []Node{*left, *right},
// 	}

// 	tree.Root = newRoot
// }

// // func (bt *BTree) splitChild(x *Node, i int) {}
