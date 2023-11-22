package priority_queue

import "testing"

type Item struct {
	Value    string
	Priority int
}

func TestPriorityQueue(t *testing.T) {
	items := []*Item{
		{Value: "init1", Priority: 6},
		{Value: "init2", Priority: 9},
		{Value: "init3", Priority: 8},
	}
	pq := NewPriorityQueue(items, func(a, b *Item) bool {
		return a.Priority > b.Priority
	})
	pq.Push(&Item{
		Value:    "value1",
		Priority: 5,
	})
	pq.Push(&Item{
		Value:    "value2",
		Priority: 3,
	})
	pq.Push(&Item{
		Value:    "value3",
		Priority: 4,
	})
	pq.Push(&Item{
		Value:    "value4",
		Priority: 1,
	})
	for i := 0; i < 10; i++ {
		popItem := pq.Pop()
		t.Logf("%+v", popItem)
	}
}
