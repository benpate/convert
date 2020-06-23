package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilToBool(t *testing.T) {

	{
		result, natural := NaturalBool(nil, true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(nil, false)
		assert.False(t, result)
		assert.False(t, natural)
	}
}

func TestBoolToBool(t *testing.T) {

	assert.True(t, Bool(true, true))
	assert.True(t, Bool(true, false))
	assert.False(t, Bool(false, true))
	assert.False(t, Bool(false, false))
}

func TestIntToBool(t *testing.T) {

	{
		result, natural := NaturalBool(0, false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(1, false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(0, true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(1, true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestInt8ToBool(t *testing.T) {

	{
		result, natural := NaturalBool(int8(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int8(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int8(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int8(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestInt16ToBool(t *testing.T) {

	{
		result, natural := NaturalBool(int16(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int16(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int16(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int16(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestInt32ToBool(t *testing.T) {

	{
		result, natural := NaturalBool(int32(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int32(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int32(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int32(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestInt64ToBool(t *testing.T) {

	{
		result, natural := NaturalBool(int64(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int64(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int64(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(int64(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestFloat32ToBool(t *testing.T) {

	{
		result, natural := NaturalBool(float32(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(float32(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(float32(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(float32(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestFloat64ToBool(t *testing.T) {

	{
		result, natural := NaturalBool(float64(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(float64(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(float64(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(float64(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestStringToBool(t *testing.T) {

	{
		result, natural := NaturalBool("true", true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool("false", true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool("true", false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool("false", false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool("Somethig Else", true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool("Somethig Else", false)
		assert.False(t, result)
		assert.False(t, natural)
	}
}

func TestStringerToBool(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := NaturalBool(s, true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(s, false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	s[0] = "true"

	{
		result, natural := NaturalBool(s, true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(s, false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	s[0] = "false"

	{
		result, natural := NaturalBool(s, true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := NaturalBool(s, false)
		assert.False(t, result)
		assert.False(t, natural)
	}
}

func TestInvalidToBool(t *testing.T) {

	result, natural := NaturalBool(map[string]interface{}{}, true)

	assert.False(t, natural)
	assert.True(t, result)
}
