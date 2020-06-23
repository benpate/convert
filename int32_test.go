package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestInt32(t *testing.T) {

	assert.Equal(t, Int32(10, 10), int32(10))

}

func TestNilToInt32(t *testing.T) {

	result, natural := NaturalInt32(nil, int32(-1))
	assert.Equal(t, result, int32(-1))
	assert.False(t, natural)
}

func TestFloat32ToInt32(t *testing.T) {

	result, natural := NaturalInt32(float32(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int32(10))
}

func TestFloat64ToInt32(t *testing.T) {

	result, natural := NaturalInt32(float64(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int32(10))
}

func TestIntToInt32(t *testing.T) {

	{
		result, natural := NaturalInt32(int(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := NaturalInt32(int8(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := NaturalInt32(int16(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := NaturalInt32(int32(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := NaturalInt32(int64(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}
}

func TestStringToInt32(t *testing.T) {

	{
		result, natural := NaturalInt32("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := NaturalInt32("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int32(-1))
	}
}

func TestStringerToInt32(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := NaturalInt32(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int32(-1))
	}

	s[0] = "100"
	{
		result, natural := NaturalInt32(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int32(100))
	}
}

func TestInvalidToInt32(t *testing.T) {
	result, natural := NaturalInt32(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, int32(-1))
}
