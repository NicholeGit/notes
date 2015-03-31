package adapter

// 通过这个接口统一适配
type Decorator interface {
	Decorate() string
}

// 原结构体
type Banner struct {
	str string
}

func (self *Banner) getString() string {
	return "*" + self.str + "*"
}

// Decorator的实现，Decorate()方法中可以把Banner转变为需要的结构
type EmbeddedDecorateBanner struct {
	*Banner
}

func NewEmbeddedDecorateBanner(str string) *EmbeddedDecorateBanner {
	return &EmbeddedDecorateBanner{&Banner{str}}
}

func (self *EmbeddedDecorateBanner) Decorate() string {
	return self.getString()
}
