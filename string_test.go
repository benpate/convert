package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Thes tests still dont account for overflow errors.

func TestString(t *testing.T) {

	assert.Equal(t, "hello there", String("hello there", ""))
	assert.Equal(t, "10", String(10, ""))
}

func TestNilToString(t *testing.T) {
	assert.Equal(t, "default", String(nil, "default"))
}

func TestStringToString(t *testing.T) {

	result, natural := NaturalString("Hello there", "default")

	assert.True(t, natural)
	assert.Equal(t, result, "Hello there")
}

func TestByteArrayToString(t *testing.T) {
	result, natural := NaturalString([]byte("...from a certain point of view."), "default")

	assert.True(t, natural)
	assert.Equal(t, result, "...from a certain point of view.")
}

func TestIntToString(t *testing.T) {

	{
		result, natural := NaturalString(int(10), "default")

		assert.False(t, natural)
		assert.Equal(t, result, "10")
	}

	{
		result, natural := NaturalString(int8(10), "default")

		assert.False(t, natural)
		assert.Equal(t, result, "10")
	}

	{
		result, natural := NaturalString(int16(10), "default")

		assert.False(t, natural)
		assert.Equal(t, result, "10")
	}

	{
		result, natural := NaturalString(int32(10), "default")

		assert.False(t, natural)
		assert.Equal(t, result, "10")
	}

	{
		result, natural := NaturalString(int64(10), "default")

		assert.False(t, natural)
		assert.Equal(t, result, "10")
	}
}

func TestFloatToString(t *testing.T) {

	{
		result, natural := NaturalString(float32(10), "default")

		assert.False(t, natural)
		assert.Equal(t, result, "10")
	}

	{
		result, natural := NaturalString(float64(10), "default")

		assert.False(t, natural)
		assert.Equal(t, result, "10")
	}
}

func TestStringerToString(t *testing.T) {

	s := getTestStringer()

	{
		result, natural := NaturalString(s, "")

		assert.False(t, natural)
		assert.Equal(t, result, "hello")
	}
}

func TestInvalidToString(t *testing.T) {
	result, natural := NaturalString(map[string]interface{}{}, "default")

	assert.False(t, natural)
	assert.Equal(t, result, "default")
}
