package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBot(t *testing.T) {
	data := make([]byte, 0)
	yascat := NewBot(data)
	assert.Equal(t, bot{}, yascat, "NewBot expected to return bot struct.")
}
