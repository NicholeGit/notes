package flyweight

import (
	"fmt"
	"testing"
)

//bigstr &{[0xc0820046a0 0xc0820046c0 0xc0820046a0]}
//共享了0xc0820046a0这段内存
//当对象可以枚举内容时，把内容存储起来，每个对象只用存储内容的地址
func TestFlyWeight(t *testing.T) {
	bigStr := NewBigString("121")
	result := bigStr.Print()

	expect := "-\n--\n-\n"
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
	fmt.Println(bigStr)
}
