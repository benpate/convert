package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestFloatToInt64(t *testing.T) {

	// Test Floats

	{
		result, ok := Int64Ok(float32(10))

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		result, ok := Int64Ok(float64(10))

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	// Test Float Pointers

	{
		input := float32(10)
		result, ok := Int64Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		input := float64(10)
		result, ok := Int64Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}
}

func TestIntToInt64(t *testing.T) {

	// Test Ints

	{
		result, ok := Int64Ok(int(10))

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		result, ok := Int64Ok(int8(10))

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		result, ok := Int64Ok(int16(10))

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		result, ok := Int64Ok(int32(10))

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		result, ok := Int64Ok(int64(10))

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	// Test Int Pointers

	{
		input := int(10)
		result, ok := Int64Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		input := int8(10)
		result, ok := Int64Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		input := int16(10)
		result, ok := Int64Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		input := int32(10)
		result, ok := Int64Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		input := int64(10)
		result, ok := Int64Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}
}

func TestStringToInt64(t *testing.T) {

	{
		result, ok := Int64Ok("10")

		assert.True(t, ok)
		assert.Equal(t, result, int64(10))
	}

	{
		result, ok := Int64Ok("invalid")

		assert.False(t, ok)
		assert.Equal(t, result, int64(0))
	}
}

func TestInvalidToInt64(t *testing.T) {

	{
		result, ok := Int64Ok(map[string]interface{}{})

		assert.False(t, ok)
		assert.Equal(t, result, int64(0))
	}
}
