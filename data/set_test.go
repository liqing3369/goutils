package data

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"testing"
)

func TestSet(t *testing.T) {
	set := mapset.NewSet("1", "2", "3")
	set.Each(func(i interface{}) bool {
		fmt.Println(i)
		return false
	})
}
