package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestInt(t *testing.T) {
	assert.Equal(t, IntDefault(10, 10), int(10))
}

func TestNilToInt(t *testing.T) {

	result, natural := IntOk(nil, int(-1))
	assert.Equal(t, result, int(-1))
	assert.False(t, natural)
}

func TestFloat32ToInt(t *testing.T) {

	result, natural := IntOk(float32(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int(10))
}

func TestFloat64ToInt(t *testing.T) {

	result, natural := IntOk(float64(10), -1)

	assert.False(t, natural)
	assert.Equal(t, result, int(10))
}

func TestIntToInt(t *testing.T) {

	{
		result, natural := IntOk(int(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := IntOk(int8(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := IntOk(int16(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := IntOk(int32(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := IntOk(int64(10), -1)

		assert.True(t, natural)
		assert.Equal(t, result, int(10))
	}
}

func TestStringToInt(t *testing.T) {

	{
		result, natural := IntOk("10", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int(10))
	}

	{
		result, natural := IntOk("invalid", -1)

		assert.False(t, natural)
		assert.Equal(t, result, int(-1))
	}
}

func TestStringerToInt(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := IntOk(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int(-1))
	}

	s[0] = "100"
	{
		result, natural := IntOk(s, -1)

		assert.False(t, natural)
		assert.Equal(t, result, int(100))
	}
}

func TestInvalidToInt(t *testing.T) {
	result, natural := IntOk(map[string]interface{}{}, -1)

	assert.False(t, natural)
	assert.Equal(t, result, int(-1))
}
