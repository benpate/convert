package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestString(t *testing.T) {

	{
		result, ok := StringOk("Hello there")

		assert.True(t, ok)
		assert.Equal(t, result, "Hello there")
	}

	{
		pointer := "I have the high ground."
		result, ok := StringOk(&pointer)

		assert.True(t, ok)
		assert.Equal(t, result, "I have the high ground.")
	}

	{
		result, ok := StringOk([]byte("...from a certain point of view."))

		assert.True(t, ok)
		assert.Equal(t, result, "...from a certain point of view.")
	}

	{
		result, ok := StringOk(map[string]interface{}{})

		assert.False(t, ok)
		assert.Equal(t, result, "")
	}
}

func TestIntToString(t *testing.T) {

	// Test Ints

	{
		result, ok := StringOk(int(10))

		assert.True(t, ok)
		assert.Equal(t, result, "10")
	}

	{
		result, ok := StringOk(int8(10))

		assert.True(t, ok)
		assert.Equal(t, result, "10")
	}

	{
		result, ok := StringOk(int16(10))

		assert.True(t, ok)
		assert.Equal(t, result, "10")
	}

	{
		result, ok := StringOk(int32(10))

		assert.True(t, ok)
		assert.Equal(t, result, "10")
	}

	{
		result, ok := StringOk(int64(10))

		assert.True(t, ok)
		assert.Equal(t, result, "10")
	}

}
