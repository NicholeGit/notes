package builder

//builder为一个接口
//Director里面有一个bulider对象
//TextBuilder实现了builder接口
//只要实现builder接口都可以被Director调用，如果增加新的功能需要实现builder接口并且director := Director{&TextBuilder{}}就完成了新的功能
type builder interface {
	makeTitle(title string) string
	makeString(str string) string
	makeItems(items []string) string
	close() string
}

type Director struct {
	builder builder
}

func (self *Director) Construct() string {
	result := self.builder.makeTitle("Title")
	result += self.builder.makeString("String")
	result += self.builder.makeItems([]string{
		"Item1",
		"Item2",
	})
	result += self.builder.close()
	return result
}

type TextBuilder struct {
}

func (self *TextBuilder) makeTitle(title string) string {
	return "# " + title + "\n"
}

func (self *TextBuilder) makeString(str string) string {
	return "## " + str + "\n"
}

func (self *TextBuilder) makeItems(items []string) string {
	var result string
	for _, item := range items {
		result += "- " + item + "\n"
	}
	return result
}

func (self *TextBuilder) close() string {
	return "\n"
}
