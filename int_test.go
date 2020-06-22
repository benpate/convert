package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestFloatToInt(t *testing.T) {

	// Test Floats

	{
		result, ok := IntOk(float32(10))

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		result, ok := IntOk(float64(10))

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	// Test Float Pointers

	{
		input := float32(10)
		result, ok := IntOk(&input)

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		input := float64(10)
		result, ok := IntOk(&input)

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}
}

func TestIntToInt(t *testing.T) {

	// Test Ints

	{
		result, ok := IntOk(int(10))

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		result, ok := IntOk(int8(10))

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		result, ok := IntOk(int16(10))

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		result, ok := IntOk(int32(10))

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		result, ok := IntOk(int64(10))

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	// Test Int Pointers

	{
		input := int(10)
		result, ok := IntOk(&input)

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		input := int8(10)
		result, ok := IntOk(&input)

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		input := int16(10)
		result, ok := IntOk(&input)

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		input := int32(10)
		result, ok := IntOk(&input)

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		input := int64(10)
		result, ok := IntOk(&input)

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}
}

func TestStringToInt(t *testing.T) {

	{
		result, ok := IntOk("10")

		assert.True(t, ok)
		assert.Equal(t, result, 10)
	}

	{
		result, ok := IntOk("invalid")

		assert.False(t, ok)
		assert.Equal(t, result, 0)
	}
}

func TestInvalidToInt(t *testing.T) {

	{
		result, ok := IntOk(map[string]interface{}{})

		assert.False(t, ok)
		assert.Equal(t, result, 0)
	}
}
