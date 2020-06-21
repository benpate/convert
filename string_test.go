package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestString(t *testing.T) {

	{
		result, err := String("Hello there")

		assert.Nil(t, err)
		assert.Equal(t, result, "Hello there")
	}

	{
		pointer := "I have the high ground."
		result, err := String(&pointer)

		assert.Nil(t, err)
		assert.Equal(t, result, "I have the high ground.")
	}

	{
		result, err := String([]byte("...from a certain point of view."))

		assert.Nil(t, err)
		assert.Equal(t, result, "...from a certain point of view.")
	}

	{
		result, err := String(map[string]interface{}{})

		assert.NotNil(t, err)
		assert.Equal(t, result, "")
	}
}

func TestIntToString(t *testing.T) {

	// Test Ints

	{
		result, err := String(int(10))

		assert.Nil(t, err)
		assert.Equal(t, result, "10")
	}

	{
		result, err := String(int8(10))

		assert.Nil(t, err)
		assert.Equal(t, result, "10")
	}

	{
		result, err := String(int16(10))

		assert.Nil(t, err)
		assert.Equal(t, result, "10")
	}

	{
		result, err := String(int32(10))

		assert.Nil(t, err)
		assert.Equal(t, result, "10")
	}

	{
		result, err := String(int64(10))

		assert.Nil(t, err)
		assert.Equal(t, result, "10")
	}

}
