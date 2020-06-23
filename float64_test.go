package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestNilToFloat64(t *testing.T) {

	result, natural := NaturalFloat64(nil, float64(-1))
	assert.Equal(t, result, float64(-1))
	assert.False(t, natural)
}

func TestFloat32ToFloat64(t *testing.T) {

	result, natural := NaturalFloat64(float32(10), -1)

	assert.True(t, natural)
	assert.Equal(t, result, float64(10))
}

func TestFloat64ToFloat64(t *testing.T) {

	assert.Equal(t, float64(10), Float64(float64(10), -1))
	assert.Equal(t, float64(-1), Float64("hello", -1))
}

func TestIntToFloat64(t *testing.T) {

	{
		result, natural := NaturalFloat64(int(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := NaturalFloat64(int8(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := NaturalFloat64(int16(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := NaturalFloat64(int32(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := NaturalFloat64(int64(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}
}

func TestStringToFloat64(t *testing.T) {

	{
		result, natural := NaturalFloat64("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := NaturalFloat64("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(-1))
	}
}

func TestStringerToFloat64(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := NaturalFloat64(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(-1))
	}

	s[0] = "100"
	{
		result, natural := NaturalFloat64(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(100))
	}
}

func TestInvalidToFloat64(t *testing.T) {
	result, natural := NaturalFloat64(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, float64(-1))
}
