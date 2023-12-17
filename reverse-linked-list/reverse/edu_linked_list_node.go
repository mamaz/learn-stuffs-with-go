package reverse

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func (node *EduLinkedListNode) GetData() int {
	return node.data
}

func ToList(head *EduLinkedListNode) []int {
	result := []int{}
	current := head

	for current != nil {
		result = append(result, current.data)
		current = current.next
	}

	return result
}

func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = nil
	return node
}
