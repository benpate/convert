package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestNilToFloat64(t *testing.T) {

	result, natural := Float64Natural(nil, float64(-1))
	assert.Equal(t, result, float64(-1))
	assert.False(t, natural)
}

func TestFloat32ToFloat64(t *testing.T) {

	result, natural := Float64Natural(float32(10), -1)

	assert.True(t, natural)
	assert.Equal(t, result, float64(10))
}

func TestFloat64ToFloat64(t *testing.T) {

	assert.Equal(t, float64(10), Float64Default(float64(10), -1))
	assert.Equal(t, float64(-1), Float64Default("hello", -1))
}

func TestIntToFloat64(t *testing.T) {

	{
		result, natural := Float64Natural(int(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := Float64Natural(int8(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := Float64Natural(int16(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := Float64Natural(int32(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := Float64Natural(int64(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}
}

func TestStringToFloat64(t *testing.T) {

	{
		result, natural := Float64Natural("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(10))
	}

	{
		result, natural := Float64Natural("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(-1))
	}
}

func TestStringerToFloat64(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := Float64Natural(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(-1))
	}

	s[0] = "100"
	{
		result, natural := Float64Natural(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, float64(100))
	}
}

func TestInvalidToFloat64(t *testing.T) {
	result, natural := Float64Natural(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, float64(-1))
}
