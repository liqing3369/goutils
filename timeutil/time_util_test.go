package timeutil

import (
	"github.com/assertgo/assert"
	"testing"
)

func TestAddTimeSuffix(t *testing.T) {
	ass := assert.New(t)
	dataStr := "2020-04-18"
	ass.ThatString("2020-04-18 23:59:59").IsEqualTo(AddTimeSuffix(dataStr))
}
