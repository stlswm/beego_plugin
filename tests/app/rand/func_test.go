package rand

import (
	"github.com/stlswm/beego_plugin/app/rand"
	"strconv"
	"testing"
)

func TestRangeNum(t *testing.T) {
	for i := 0; i < 1000; i++ {
		num := rand.RangeNum(1, 100)
		if num >= 1 && num <= 100 {
			t.Log("ok:" + strconv.Itoa(num))
		} else {
			t.Error("error num:" + strconv.Itoa(num))
		}
	}
}
func TestRandomStrNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		str := rand.RandomStr(rand.KindAll, 32)
		t.Log("ok:" + str)
	}
}
