package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestFloatToFloat32(t *testing.T) {

	// Test Floats

	{
		result, ok := Float32Ok(float32(10))

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		result, ok := Float32Ok(float64(10))

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	// Test Float Pointers

	{
		input := float32(10)
		result, ok := Float32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		input := float64(10)
		result, ok := Float32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}
}

func TestIntToFloat32(t *testing.T) {

	// Test Ints

	{
		result, ok := Float32Ok(int(10))

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		result, ok := Float32Ok(int8(10))

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		result, ok := Float32Ok(int16(10))

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		result, ok := Float32Ok(int32(10))

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		result, ok := Float32Ok(int64(10))

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	// Test Int Pointers

	{
		input := int(10)
		result, ok := Float32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		input := int8(10)
		result, ok := Float32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		input := int16(10)
		result, ok := Float32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		input := int32(10)
		result, ok := Float32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		input := int64(10)
		result, ok := Float32Ok(&input)

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}
}

func TestStringToFloat32(t *testing.T) {

	{
		result, ok := Float32Ok("10")

		assert.True(t, ok)
		assert.Equal(t, result, float32(10))
	}

	{
		result, ok := Float32Ok("invalid")

		assert.False(t, ok)
		assert.Equal(t, result, float32(0))
	}
}

func TestInvalidToFloat32(t *testing.T) {

	{
		result, ok := Float32Ok(map[string]interface{}{})

		assert.False(t, ok)
		assert.Equal(t, result, float32(0))
	}
}
