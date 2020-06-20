package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestFloatToInt32(t *testing.T) {

	// Test Floats

	{
		result, err := Int32(float32(10))

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		result, err := Int32(float64(10))

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	// Test Float Pointers

	{
		input := float32(10)
		result, err := Int32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		input := float64(10)
		result, err := Int32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}
}

func TestIntToInt32(t *testing.T) {

	// Test Ints

	{
		result, err := Int32(int(10))

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		result, err := Int32(int8(10))

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		result, err := Int32(int16(10))

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		result, err := Int32(int32(10))

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		result, err := Int32(int64(10))

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	// Test Int Pointers

	{
		input := int(10)
		result, err := Int32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		input := int8(10)
		result, err := Int32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		input := int16(10)
		result, err := Int32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		input := int32(10)
		result, err := Int32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		input := int64(10)
		result, err := Int32(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}
}

func TestStringToInt32(t *testing.T) {

	{
		result, err := Int32("10")

		assert.Nil(t, err)
		assert.Equal(t, result, int32(10))
	}

	{
		result, err := Int32("invalid")

		assert.NotNil(t, err)
		assert.Equal(t, result, int32(0))
	}
}

func TestInvalidToInt32(t *testing.T) {

	{
		result, err := Int32(map[string]interface{}{})

		assert.NotNil(t, err)
		assert.Equal(t, result, int32(0))
	}
}
