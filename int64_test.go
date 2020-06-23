package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestInt64(t *testing.T) {

	assert.Equal(t, Int64(10, 10), int64(10))

}

func TestNilToInt64(t *testing.T) {

	result, natural := NaturalInt64(nil, -1)
	assert.Equal(t, result, int64(-1))
	assert.False(t, natural)
}

func TestFloat32ToInt64(t *testing.T) {

	result, natural := NaturalInt64(float32(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int64(10))
}

func TestFloat64ToInt64(t *testing.T) {

	result, natural := NaturalInt64(float64(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int64(10))
}

func TestIntToInt64(t *testing.T) {

	{
		result, natural := NaturalInt64(int(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := NaturalInt64(int8(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := NaturalInt64(int16(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := NaturalInt64(int32(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := NaturalInt64(int64(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}
}

func TestStringToInt64(t *testing.T) {

	{
		result, natural := NaturalInt64("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := NaturalInt64("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int64(-1))
	}
}

func TestStringerToInt64(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := NaturalInt64(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int64(-1))
	}

	s[0] = "100"
	{
		result, natural := NaturalInt64(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int64(100))
	}
}

func TestInvalidToInt64(t *testing.T) {
	result, natural := NaturalInt64(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, int64(-1))
}
