package reverse

func Reverse(head *EduLinkedListNode) *EduLinkedListNode {
	var prev *EduLinkedListNode = nil
	current := head
	next := head.next

	if current == nil {
		return nil
	}

	if next == nil {
		return current
	}

	for next != nil {
		prev = current
		current = next
		next = current.next
		current.next = prev
		if prev == head {
			prev.next = nil
		}
	}

	return current
}
