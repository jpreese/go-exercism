package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Get_EmptyList_ThrowsError(t *testing.T) {
	var list LinkedList

	_, err := list.Get(0)

	assert.Error(t, err)
}

func Test_Get_AddedSingleValue_ReturnsValue(t *testing.T) {
	var list LinkedList
	expected := 1
	list.Add(1)

	actual, _ := list.Get(0)

	assert.Equal(t, expected, actual)
}

func Test_Get_AddedTwoValues_ReturnsSecondValue(t *testing.T) {
	var list LinkedList
	list.Add(1)
	list.Add(2)

	expected := 2

	actual, _ := list.Get(1)

	assert.Equal(t, expected, actual)
}

func Test_Add_SingleValue_ReturnsAddedNode(t *testing.T) {
	var list LinkedList

	expected := 1

	actual := list.Add(1)

	assert.Equal(t, expected, actual.Value)
}
