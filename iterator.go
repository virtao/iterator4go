package iterator4go

type Iterator struct {
	current int
	data    []interface{}
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data, current: -1}
}

func (it *Iterator) Value() interface{} {
	return it.data[it.current]
}
func (it *Iterator) Next() bool {
	if it.current+1 >= len(it.data) {
		return false
	}
	it.current++
	return true
}

type IntIterator struct {
	current int
	data    []int
}

func NewIntIterator(data []int) *IntIterator {
	return &IntIterator{data: data, current: -1}
}

func (it *IntIterator) Value() int {
	return it.data[it.current]
}
func (it *IntIterator) Next() bool {
	if it.current+1 >= len(it.data) {
		return false
	}
	it.current++
	return true
}
