package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilToBool(t *testing.T) {

	{
		result, natural := BoolNatural(nil, true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(nil, false)
		assert.False(t, result)
		assert.False(t, natural)
	}
}

func TestBoolToBool(t *testing.T) {

	assert.True(t, BoolDefault(true, true))
	assert.True(t, BoolDefault(true, false))
	assert.False(t, BoolDefault(false, true))
	assert.False(t, BoolDefault(false, false))
}

func TestIntToBool(t *testing.T) {

	{
		result, natural := BoolNatural(0, false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(1, false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(0, true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(1, true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestInt8ToBool(t *testing.T) {

	{
		result, natural := BoolNatural(int8(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int8(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int8(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int8(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestInt16ToBool(t *testing.T) {

	{
		result, natural := BoolNatural(int16(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int16(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int16(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int16(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestInt32ToBool(t *testing.T) {

	{
		result, natural := BoolNatural(int32(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int32(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int32(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int32(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestInt64ToBool(t *testing.T) {

	{
		result, natural := BoolNatural(int64(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int64(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int64(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(int64(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestFloat32ToBool(t *testing.T) {

	{
		result, natural := BoolNatural(float32(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(float32(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(float32(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(float32(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestFloat64ToBool(t *testing.T) {

	{
		result, natural := BoolNatural(float64(0), false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(float64(1), false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(float64(0), true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(float64(1), true)
		assert.True(t, result)
		assert.False(t, natural)
	}
}

func TestStringToBool(t *testing.T) {

	{
		result, natural := BoolNatural("true", true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural("false", true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural("true", false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural("false", false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural("Somethig Else", true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural("Somethig Else", false)
		assert.False(t, result)
		assert.False(t, natural)
	}
}

func TestStringerToBool(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := BoolNatural(s, true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(s, false)
		assert.False(t, result)
		assert.False(t, natural)
	}

	s[0] = "true"

	{
		result, natural := BoolNatural(s, true)
		assert.True(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(s, false)
		assert.True(t, result)
		assert.False(t, natural)
	}

	s[0] = "false"

	{
		result, natural := BoolNatural(s, true)
		assert.False(t, result)
		assert.False(t, natural)
	}

	{
		result, natural := BoolNatural(s, false)
		assert.False(t, result)
		assert.False(t, natural)
	}
}

func TestInvalidToBool(t *testing.T) {

	result, natural := BoolNatural(map[string]interface{}{}, true)

	assert.False(t, natural)
	assert.True(t, result)
}
