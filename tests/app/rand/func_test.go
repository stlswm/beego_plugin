package rand

import (
	"github.com/stlswm/beego_plugin/app/rand"
	"strconv"
	"testing"
)

func TestRangeNum(t *testing.T) {
	min := -100
	max := 100
	for i := 0; i < 1000; i++ {
		num := rand.RangeNum(min, max)
		if num >= min && num < max {
			t.Log("ok:" + strconv.Itoa(num))
		} else {
			panic("error num:" + strconv.Itoa(num))
		}
	}
}
func TestRandomStrNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		str := rand.RandomStr(rand.KindAll, 32)
		t.Log(str)
	}
}
