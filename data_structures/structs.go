package data_structures

// Linked List Implementation

type Node struct{
	data int
	next *Node     // Pointer to the next node of the single linked list
}

var start *Node = nil  	//  Start Pointer to the Linked list

func Add(elem int) {
	node := Node{elem, nil}   // crate a new node
	if start == nil{    // if initially the list is empty, then add as a first node
		start = &node
	}else{				// else, traverse the list to end then add the new node
		ptr:=start
		for ptr.next!=nil{
			ptr = ptr.next
		}
		ptr.next = &node
	}
	return
}

func Search(elem int) *Node {   // traverse from start to end and match each node with the required elem
	ptr := start
	for ptr!=nil{
		if ptr.data == elem{  // if found then return elem
			return ptr
		} 
		ptr = ptr.next
	}
	return nil   // return nil if not found
}
