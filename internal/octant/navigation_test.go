package octant_test

import (
	"testing"

	"github.com/heptio/developer-dash/internal/octant"
	"github.com/stretchr/testify/assert"
)

func Test_NewNavigation(t *testing.T) {
	path := "/path"
	title := "title"

	nav := octant.NewNavigation(title, path)

	assert.Equal(t, path, nav.Path)
	assert.Equal(t, title, nav.Title)
}