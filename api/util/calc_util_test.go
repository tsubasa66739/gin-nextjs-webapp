package util_test

import (
	"testing"

	"github.com/tsubasa66739/gin-nextjs-webapp/util"
)

func TestSum(t *testing.T) {
	got := util.Sum(2, 4)
	want := 6
	if got != want {
		t.Errorf("Sum(2, 4) = %d; want 6", got)
	}
}
