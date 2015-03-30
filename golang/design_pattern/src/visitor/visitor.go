package visitor

import (
	"strconv"
)

type visitor interface {
	visitFile(file *file) string
	visitDir(directory *directory) string
}

type element interface {
	Accept(visitor visitor) string
}

type entry interface {
	element
	getName() string
	getSize() int
	Add(entry entry)
}

type defaultEntry struct {
	entry
	name string
}

func (self *defaultEntry) getName() string {
	return self.name
}

func (self *defaultEntry) print(entry entry) string {
	return entry.getName() + " (" + strconv.Itoa(entry.getSize()) + ")\n"
}

type file struct {
	*defaultEntry
	size int
}

func (self *file) getSize() int {
	return self.size
}

func (self *file) Add(entry entry) {}

func (self *file) Accept(visitor visitor) string {
	return visitor.visitFile(self)
}

type directory struct {
	*defaultEntry
	dir []entry
}

func (self *directory) getSize() int {
	size := 0
	for _, dir := range self.dir {
		size += dir.getSize()
	}
	return size
}

func (self *directory) Add(entry entry) {
	self.dir = append(self.dir, entry)
}

//调用visitor来访问自己的方法，这样自己的结构变化和接口变化也只用修改visitor内，无需改变Accept结构
func (self *directory) Accept(visitor visitor) string {
	return visitor.visitDir(self)
}

// 接口visitor的实现，这个类隔离了file与directory类的变化
type listVisitor struct {
	currentDir string
}

func (self *listVisitor) visitFile(file *file) string {
	return self.currentDir + "/" + file.print(file)
}

func (self *listVisitor) visitDir(directory *directory) string {
	saveDir := self.currentDir
	list := self.currentDir + "/" + directory.print(directory)
	self.currentDir += "/" + directory.getName()
	for _, dir := range directory.dir {
		list += dir.Accept(self)
	}
	self.currentDir = saveDir
	return list
}

func NewFile(name string, size int) *file {
	return &file{
		defaultEntry: &defaultEntry{name: name},
		size:         size,
	}
}
func NewDirectory(name string) *directory {
	return &directory{defaultEntry: &defaultEntry{name: name}}
}
