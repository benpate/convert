package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestFloatToFloat32(t *testing.T) {

	// Test Floats

	{
		result, err := Float32(float32(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		result, err := Float32(float64(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	// Test Float Pointers

	{
		input := float32(10)
		result, err := Float32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		input := float64(10)
		result, err := Float32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}
}

func TestIntToFloat32(t *testing.T) {

	// Test Ints

	{
		result, err := Float32(int(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		result, err := Float32(int8(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		result, err := Float32(int16(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		result, err := Float32(int32(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		result, err := Float32(int64(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	// Test Int Pointers

	{
		input := int(10)
		result, err := Float32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		input := int8(10)
		result, err := Float32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		input := int16(10)
		result, err := Float32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		input := int32(10)
		result, err := Float32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		input := int64(10)
		result, err := Float32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}
}

func TestStringToFloat32(t *testing.T) {

	{
		result, err := Float32("10")

		assert.Nil(t, err)
		assert.Equal(t, result, float32(10))
	}

	{
		result, err := Float32("invalid")

		assert.NotNil(t, err)
		assert.Equal(t, result, float32(0))
	}
}

func TestInvalidToFloat32(t *testing.T) {

	{
		result, err := Float32(map[string]interface{}{})

		assert.NotNil(t, err)
		assert.Equal(t, result, float32(0))
	}
}
