// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-12-18 9:58
// version: 1.0.0
// desc   : 

package sum

import (
	"testing"
)

func Test(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	t.Log(_1(nums, target))

	t.Log("====")

	t.Log(_2(nums, target))
}
