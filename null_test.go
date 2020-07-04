package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNull(t *testing.T) {

	{
		v := NullBool(true)
		assert.True(t, v.Bool())
		assert.True(t, v.IsPresent())
		assert.False(t, v.IsNull())
	}

	{
		v := NullBool("true")
		assert.False(t, v.Bool())
		assert.False(t, v.IsPresent())
		assert.True(t, v.IsNull())
	}

}
