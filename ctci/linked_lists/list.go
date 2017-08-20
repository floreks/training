package linked_lists

// PSEUDOCODE - Assuming doubly linked list
//
//func RemoveDups(ListNode *head) {
//	uniques := map[int]bool
//	node := head
//	for node != nil {
//		if uniques[node.value] {
//			prev := node.prev
//			next := node.next
//
//			prev.next = next
//
//			node.prev = nil
//			node.next = nil
//			if next != nil {
//				next.prev = prev
//				node = next
//			}
//		}
//
//		uniques[node.value] = true
//		node = node.next
//	}
//
//	return
//}
