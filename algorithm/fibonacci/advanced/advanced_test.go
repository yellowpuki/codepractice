package advanced_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yellowpuki/codepractice/advanced"
)

func TestCachedFibo(t *testing.T) {
	assert.Equal(t, 8, advanced.CachedFibo(5))
}
