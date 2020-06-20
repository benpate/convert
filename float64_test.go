package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestFloatToFloat64(t *testing.T) {

	// Test Floats

	{
		result, err := Float64(float32(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		result, err := Float64(float64(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	// Test Float Pointers

	{
		input := float32(10)
		result, err := Float64(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		input := float64(10)
		result, err := Float64(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}
}

func TestIntToFloat64(t *testing.T) {

	// Test Ints

	{
		result, err := Float64(int(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		result, err := Float64(int8(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		result, err := Float64(int16(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		result, err := Float64(int32(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		result, err := Float64(int64(10))

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	// Test Int Pointers

	{
		input := int(10)
		result, err := Float64(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		input := int8(10)
		result, err := Float64(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		input := int16(10)
		result, err := Float64(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		input := int32(10)
		result, err := Float64(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		input := int64(10)
		result, err := Float64(&input)

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}
}

func TestStringToFloat64(t *testing.T) {

	{
		result, err := Float64("10")

		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	}

	{
		result, err := Float64("invalid")

		assert.NotNil(t, err)
		assert.Equal(t, result, float64(0))
	}
}

func TestInvalidToFloat64(t *testing.T) {

	{
		result, err := Float64(map[string]interface{}{})

		assert.NotNil(t, err)
		assert.Equal(t, result, float64(0))
	}
}
