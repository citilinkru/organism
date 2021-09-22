package organism

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOrganism(t *testing.T) {
	o := New()
	assert.True(t, o.IsAlive())
	assert.False(t, o.IsReady())

	o.Ready()

	assert.True(t, o.IsReady())

	o.Die()

	assert.False(t, o.IsAlive())
	assert.True(t, o.IsReady())
}

func TestOrganism_GrowLimb_One(t *testing.T) {
	o := New()
	limb := o.GrowLimb()

	assert.True(t, limb.IsAlive())
	assert.False(t, limb.IsReady())

	limb.Ready()

	assert.True(t, limb.IsAlive())
	assert.True(t, limb.IsReady())
	assert.True(t, o.IsAlive())
	assert.False(t, o.IsReady())

	o.Ready()

	assert.True(t, limb.IsAlive())
	assert.True(t, limb.IsReady())
	assert.True(t, o.IsAlive())
	assert.True(t, o.IsReady())

	limb.Die()

	assert.False(t, limb.IsAlive())
	assert.True(t, limb.IsReady())
	assert.False(t, o.IsAlive())
	assert.True(t, o.IsReady())

	o.Die()

	assert.False(t, limb.IsAlive())
	assert.True(t, limb.IsReady())
	assert.False(t, o.IsAlive())
	assert.True(t, o.IsReady())
}
