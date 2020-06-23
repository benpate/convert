package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestNilToFloat32(t *testing.T) {

	result, natural := NaturalFloat32(nil, float32(-1))
	assert.Equal(t, result, float32(-1))
	assert.False(t, natural)
}

func TestFloat32ToFloat32(t *testing.T) {

	assert.Equal(t, float32(10), Float32(float32(10), -1))
	assert.Equal(t, float32(-1), Float32("hello", -1))
}

func TestFloat64ToFloat32(t *testing.T) {

	result, natural := NaturalFloat32(float64(10), -1)

	assert.True(t, natural)
	assert.Equal(t, result, float32(10))
}

func TestIntToFloat32(t *testing.T) {

	{
		result, natural := NaturalFloat32(int(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := NaturalFloat32(int8(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := NaturalFloat32(int16(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := NaturalFloat32(int32(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := NaturalFloat32(int64(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}
}

func TestStringToFloat32(t *testing.T) {

	{
		result, natural := NaturalFloat32("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := NaturalFloat32("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(-1))
	}
}

func TestStringerToFloat32(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := NaturalFloat32(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(-1))
	}

	s[0] = "100"
	{
		result, natural := NaturalFloat32(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(100))
	}
}

func TestInvalidToFloat32(t *testing.T) {
	result, natural := NaturalFloat32(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, float32(-1))
}
