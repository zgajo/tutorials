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
