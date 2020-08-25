/*Package mojtest b-tree
- balanced m-way tree (m way is order)
- order (m): 5 - how much pointers can be in node
- all data is sorted
- every leaf node must be in same leaf node
- every node has max m children (leafs)
- min children: leaf - 0, root: 2, internal node: cieling od m/2 (5/2 = 2.5 = cieling 3)
- max keys in node are m-1
- min keys in root 1, all other nodes: cieling(m/2)-1
- insertion is always done in leaf node
- while inserting, sort has to be done
*/
package mojtest
