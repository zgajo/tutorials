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
	Keys     []int  // An array of keys
	Children []Node // An array of child pointers
	Parent   *Node
}

// CreateNode : Creates new node
func CreateNode() *Node {
	return &Node{
		Keys:     []int{},
		Children: []Node{},
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

// isLeaf : is true when node is leaf. Otherwise false
func (tree *BTree) isNodeFull(node *Node) bool {
	return node.numOfKeys() == tree.MaxDegree-1
}

// InsertKey - INSERT KEY
func (tree *BTree) InsertKey(key int) {
	// check if root is null
	if tree.Root == nil {
		// create new node and set it as root
		tree.Root = CreateNode()
		tree.Root.Keys = append(tree.Root.Keys, key)
		return
	}
	// check if root is leaf
	if tree.Root.isLeaf() {
		// if root is not full, sort and insert key into root node
		if !tree.isNodeFull(tree.Root) {
			tree.insertKeyIntoLeaf(tree.Root, key)
			return
		}

		// else
		tree.insertKeyIntoLeaf(tree.Root, key)
		tree.splitRootWhenLeaf()
		return
	}

	// if root is not leaf
	tree.searchChild(tree.Root, key)
}

func (tree *BTree) insertKeyIntoLeaf(node *Node, key int) {
	// Initialize index as index of rightmost element
	i := node.numOfKeys() - 1
	// insert key as last element
	node.Keys = append(node.Keys, key)

	// The following loop:
	// a) Moves all greater keys than last inserted to one place ahead
	for i >= 0 && node.Keys[i] > key {
		tmp := node.Keys[i+1]
		node.Keys[i+1] = node.Keys[i]
		node.Keys[i] = tmp
		i--
	}
}

func (tree *BTree) splitRootWhenLeaf() {
	middle := tree.MinDegree
	// create node left[:middle] and node right [middle+1:]
	left := &Node{Keys: append([]int(nil), tree.Root.Keys[:middle]...)}
	right := &Node{Keys: append([]int(nil), tree.Root.Keys[middle+1:]...)}

	// create new node with one key [middle] from previous root and set previous created left and right nodes as his two children
	newRoot := &Node{
		Keys: []int{tree.Root.Keys[middle]},
	}

	left.Parent = newRoot
	right.Parent = newRoot

	newRoot.Children = []Node{*left, *right}

	tree.Root = newRoot
}

func (tree *BTree) searchChild(node *Node, key int) {
	// search for index of last index key which is larger then key which is going to be inserted
	childIndex := node.searchChildrenIndexPosition(key)

	// use found index key as index of children array node
	// repeat this step for every nested child until last node is leaf
	if !node.Children[childIndex].isLeaf() {
		fmt.Println("Not found leaf")
		tree.searchChild(&node.Children[childIndex], key)
	}

	// insert key into leaf
	// if node is not full insert key into node and sort node
	if !tree.isNodeFull(&node.Children[childIndex]) {
		tree.insertKeyIntoLeaf(&node.Children[childIndex], key)
		return
	}

	// else
	tree.insertKeyIntoLeaf(&node.Children[childIndex], key)
	tree.splitLeaf(&node.Children[childIndex], key)
}

func (tree *BTree) splitLeaf(node *Node, key int) {
	middle := tree.MinDegree
	prevParent := node.Parent
	// create node left[:middle] and node right [middle+1:]
	left := &Node{Keys: append([]int(nil), node.Keys[:middle]...)}
	right := &Node{Keys: append([]int(nil), node.Keys[middle+1:]...)}

	// create new node with one key [middle] from previous root and set previous created left and right nodes as his two children
	newNode := &Node{
		Keys: []int{node.Keys[middle]},
	}

	left.Parent = prevParent
	right.Parent = prevParent

	newNode.Children = []Node{*left, *right}

	node = newNode

	tree.insertIntoParent(prevParent, node)
}

func (tree *BTree) insertIntoParent(parent *Node, node *Node) {
	if !tree.isNodeFull(parent) {
		// Initialize index as index of rightmost element
		i := parent.numOfKeys() - 1
		// insert key as last element
		key := node.Keys[0]
		parent.Keys = append(parent.Keys, key)

		// The following loop:
		// a) Moves all greater keys than last inserted to one place ahead
		for i >= 0 && parent.Keys[i] > key {
			tmp := parent.Keys[i+1]
			parent.Keys[i+1] = parent.Keys[i]
			parent.Keys[i] = tmp
			i--
		}

		// Insert empty node which will be deleted
		parent.Children = append(parent.Children, Node{})

		pos := i + 1

		for j := pos; j < len(parent.Children)-1; j++ {
			tmp := parent.Children[j+1]
			parent.Children[j+1] = parent.Children[j]
			parent.Children[j] = tmp
		}

		parent.Children[pos] = node.Children[0]
		parent.Children[pos+1] = node.Children[1]

	}
}

func (node *Node) searchChildrenIndexPosition(key int) int {
	low, high := 0, node.numOfKeys()-1

	for low <= high {
		middle := (high + low) / 2

		if node.Keys[middle] > key {
			high = middle - 1
		} else if node.Keys[middle] < key {
			low = middle + 1
		} else {
			return middle
		}
	}

	return low
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
