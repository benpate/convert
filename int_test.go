package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestFloatToInt(t *testing.T) {

	// Test Floats

	{
		result, err := Int(float32(10))

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		result, err := Int(float64(10))

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	// Test Float Pointers

	{
		input := float32(10)
		result, err := Int(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		input := float64(10)
		result, err := Int(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}
}

func TestIntToInt(t *testing.T) {

	// Test Ints

	{
		result, err := Int(int(10))

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		result, err := Int(int8(10))

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		result, err := Int(int16(10))

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		result, err := Int(int32(10))

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		result, err := Int(int64(10))

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	// Test Int Pointers

	{
		input := int(10)
		result, err := Int(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		input := int8(10)
		result, err := Int(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		input := int16(10)
		result, err := Int(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		input := int32(10)
		result, err := Int(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		input := int64(10)
		result, err := Int(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}
}

func TestStringToInt(t *testing.T) {

	{
		result, err := Int("10")

		assert.Nil(t, err)
		assert.Equal(t, result, 10)
	}

	{
		result, err := Int("invalid")

		assert.NotNil(t, err)
		assert.Equal(t, result, 0)
	}
}

func TestInvalidToInt(t *testing.T) {

	{
		result, err := Int(map[string]interface{}{})

		assert.NotNil(t, err)
		assert.Equal(t, result, 0)
	}
}
