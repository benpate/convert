package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestInt(t *testing.T) {
	assert.Equal(t, Int(10, 10), int(10))
}

func TestNilToInt(t *testing.T) {

	result, natural := NaturalInt(nil, int(-1))
	assert.Equal(t, result, int(-1))
	assert.False(t, natural)
}

func TestFloat32ToInt(t *testing.T) {

	result, natural := NaturalInt(float32(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int(10))
}

func TestFloat64ToInt(t *testing.T) {

	result, natural := NaturalInt(float64(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int(10))
}

func TestIntToInt(t *testing.T) {

	{
		result, natural := NaturalInt(int(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := NaturalInt(int8(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := NaturalInt(int16(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := NaturalInt(int32(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := NaturalInt(int64(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}
}

func TestStringToInt(t *testing.T) {

	{
		result, natural := NaturalInt("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := NaturalInt("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int(-1))
	}
}

func TestStringerToInt(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := NaturalInt(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int(-1))
	}

	s[0] = "100"
	{
		result, natural := NaturalInt(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int(100))
	}
}

func TestInvalidToInt(t *testing.T) {
	result, natural := NaturalInt(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, int(-1))
}
