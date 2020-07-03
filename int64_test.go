package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestInt64(t *testing.T) {

	assert.Equal(t, Int64Default(10, 10), int64(10))

}

func TestNilToInt64(t *testing.T) {

	result, natural := Int64Natural(nil, -1)
	assert.Equal(t, result, int64(-1))
	assert.False(t, natural)
}

func TestFloat32ToInt64(t *testing.T) {

	result, natural := Int64Natural(float32(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int64(10))
}

func TestFloat64ToInt64(t *testing.T) {

	result, natural := Int64Natural(float64(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int64(10))
}

func TestIntToInt64(t *testing.T) {

	{
		result, natural := Int64Natural(int(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := Int64Natural(int8(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := Int64Natural(int16(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := Int64Natural(int32(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := Int64Natural(int64(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int64(10))
	}
}

func TestStringToInt64(t *testing.T) {

	{
		result, natural := Int64Natural("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int64(10))
	}

	{
		result, natural := Int64Natural("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int64(-1))
	}
}

func TestStringerToInt64(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := Int64Natural(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int64(-1))
	}

	s[0] = "100"
	{
		result, natural := Int64Natural(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int64(100))
	}
}

func TestInvalidToInt64(t *testing.T) {
	result, natural := Int64Natural(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, int64(-1))
}
