package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestFloatToInt32(t *testing.T) {

	// Test Floats

	{
		result, ok := Int32Ok(float32(10))

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		result, ok := Int32Ok(float64(10))

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	// Test Float Pointers

	{
		input := float32(10)
		result, ok := Int32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		input := float64(10)
		result, ok := Int32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}
}

func TestIntToInt32(t *testing.T) {

	// Test Ints

	{
		result, ok := Int32Ok(int(10))

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		result, ok := Int32Ok(int8(10))

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		result, ok := Int32Ok(int16(10))

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		result, ok := Int32Ok(int32(10))

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		result, ok := Int32Ok(int64(10))

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	// Test Int Pointers

	{
		input := int(10)
		result, ok := Int32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		input := int8(10)
		result, ok := Int32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		input := int16(10)
		result, ok := Int32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		input := int32(10)
		result, ok := Int32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		input := int64(10)
		result, ok := Int32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}
}

func TestStringToInt32(t *testing.T) {

	{
		result, ok := Int32Ok("10")

		assert.True(t, ok)
		assert.Equal(t, result, int32(10))
	}

	{
		result, ok := Int32Ok("invalid")

		assert.False(t, ok)
		assert.Equal(t, result, int32(0))
	}
}

func TestInvalidToInt32(t *testing.T) {

	{
		result, ok := Int32Ok(map[string]interface{}{})

		assert.False(t, ok)
		assert.Equal(t, result, int32(0))
	}
}
