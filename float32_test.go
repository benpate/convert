package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestNilToFloat32(t *testing.T) {

	result, natural := Float32Natural(nil, float32(-1))
	assert.Equal(t, result, float32(-1))
	assert.False(t, natural)
}

func TestFloat32ToFloat32(t *testing.T) {

	assert.Equal(t, float32(10), Float32Default(float32(10), -1))
	assert.Equal(t, float32(-1), Float32Default("hello", -1))
}

func TestFloat64ToFloat32(t *testing.T) {

	result, natural := Float32Natural(float64(10), -1)

	assert.True(t, natural)
	assert.Equal(t, result, float32(10))
}

func TestIntToFloat32(t *testing.T) {

	{
		result, natural := Float32Natural(int(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := Float32Natural(int8(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := Float32Natural(int16(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := Float32Natural(int32(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := Float32Natural(int64(10), -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}
}

func TestStringToFloat32(t *testing.T) {

	{
		result, natural := Float32Natural("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(10))
	}

	{
		result, natural := Float32Natural("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(-1))
	}
}

func TestStringerToFloat32(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := Float32Natural(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(-1))
	}

	s[0] = "100"
	{
		result, natural := Float32Natural(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, float32(100))
	}
}

func TestInvalidToFloat32(t *testing.T) {
	result, natural := Float32Natural(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, float32(-1))
}
