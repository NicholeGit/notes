package chain_of_responsibility

import (
	"strconv"
)

type Trouble struct {
	number int
}

func (self *Trouble) getNumber() int {
	return self.number
}

//=====================
type support interface {
	resolve(trouble Trouble) bool
	Handle(support support, trouble Trouble) string
}

type defaultSupport struct {
	support
	name string
	next support //记录下个处理链
}

//support接口和defaultSupport类完成职责链，（需要两个组合完成是以为golang的接口不支持方法定义）。
//support提供基本的接口，
//defaultSupport添加一些默认的方法，这个类不能实例化因为resolve没有被实现，所以必须用具体类来实现resolve。
//=====================

func (self *defaultSupport) SetNext(next support) {
	self.next = next
}

func (self *defaultSupport) Handle(support support, trouble Trouble) string {
	if support.resolve(trouble) { //因为resolve没有被实现，所以需要把自己的support传递进来
		return self.done(trouble)
	} else if self.next != nil {
		return self.next.Handle(self.next, trouble)
	} else {
		return self.fail(trouble)
	}
}

func (self *defaultSupport) done(trouble Trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " is resolved by " + self.name
}

func (self *defaultSupport) fail(trouble Trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " cannot be resolved"
}

type noSupport struct {
	*defaultSupport
}

func (self *noSupport) resolve(trouble Trouble) bool {
	return false
}

func NewNoSupport(name string) *noSupport {
	return &noSupport{&defaultSupport{name: name}}
}

type limitSupport struct {
	*defaultSupport
	limit int
}

func (self *limitSupport) resolve(trouble Trouble) bool {
	if trouble.getNumber() < self.limit {
		return true
	} else {
		return false
	}
}

func NewLimitSupport(name string, limit int) *limitSupport {
	return &limitSupport{&defaultSupport{name: name}, limit}
}
