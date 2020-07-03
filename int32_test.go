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

	result, natural := Int32Natural(nil, int32(-1))
	assert.Equal(t, result, int32(-1))
	assert.False(t, natural)
}

func TestFloat32ToInt32(t *testing.T) {

	result, natural := Int32Natural(float32(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int32(10))
}

func TestFloat64ToInt32(t *testing.T) {

	result, natural := Int32Natural(float64(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int32(10))
}

func TestIntToInt32(t *testing.T) {

	{
		result, natural := Int32Natural(int(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := Int32Natural(int8(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := Int32Natural(int16(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := Int32Natural(int32(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := Int32Natural(int64(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int32(10))
	}
}

func TestStringToInt32(t *testing.T) {

	{
		result, natural := Int32Natural("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int32(10))
	}

	{
		result, natural := Int32Natural("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int32(-1))
	}
}

func TestStringerToInt32(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := Int32Natural(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int32(-1))
	}

	s[0] = "100"
	{
		result, natural := Int32Natural(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int32(100))
	}
}

func TestInvalidToInt32(t *testing.T) {
	result, natural := Int32Natural(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, int32(-1))
}
