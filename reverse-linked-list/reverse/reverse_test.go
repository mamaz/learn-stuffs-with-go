package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeAbleToReverseLinkedList(t *testing.T) {
	ll := EduLinkedList{}
	ll.CreateLinkedList([]int{1, 2, 3, 4, 5})

	resultHead := Reverse(ll.head)

	ll.head = resultHead

	assert.Equal(t, []int{5, 4, 3, 2, 1}, ToList(resultHead))
}
