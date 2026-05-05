package arithmatic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	ctx := context.Background()
	ad := NewAdapter()
	answer, err := ad.Add(ctx, 1, 1)
	assert.Equal(t, answer, 2)
	assert.Nil(t, err)
}

func TestSub(t *testing.T) {
	ctx := context.Background()
	ad := NewAdapter()
	answer, err := ad.Sub(ctx, 1, 1)
	assert.Equal(t, answer, 0)
	assert.Nil(t, err)
}

func TestMul(t *testing.T) {
	ctx := context.Background()
	ad := NewAdapter()
	answer, err := ad.Mul(ctx, 2, 2)
	assert.Equal(t, answer, 4)
	assert.Nil(t, err)
}

func TestDiv(t *testing.T) {
	ctx := context.Background()
	ad := NewAdapter()
	answer, err := ad.Div(ctx, 2, 2)
	assert.Equal(t, answer, float64(1))
	assert.Nil(t, err)
}
