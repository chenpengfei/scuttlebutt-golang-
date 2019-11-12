package event

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmit(t *testing.T) {
	event := NewEvent()
	total1 := 0
	event.On("_update", func(opts ...interface{}) {
		total1 = total1 + opts[0].(int)
	})

	total2 := 0
	event.On("_update", func(opts ...interface{}) {
		total2 = total2 + opts[0].(int)
	})

	type user struct {
		Name string
		Age  int
	}
	counter := 0
	event.On("_user", func(opts ...interface{}) {
		user := opts[0].(user)
		assert.Equal(t, "Alice", user.Name)
		assert.Equal(t, 3, user.Age)
		counter++
	})

	event.Emit("_update", 1)
	event.Emit("_user", user{
		Name: "Alice",
		Age:  3,
	})
	event.Emit("_update", 2)
	event.Emit("_user", user{
		Name: "Alice",
		Age:  3,
	})

	assert.Equal(t, 3, total1)
	assert.Equal(t, 3, total2)
	assert.Equal(t, 2, counter)
}