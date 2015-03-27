package abstract_factory

//产品类继承层次
type item interface {
	toString() string
}

type link interface {
	item
}

type tray interface {
	item
	AddToTray(item item)
}

type page interface {
	AddToContent(item item)
	Output() string
}

//=============================
// 抽象工厂
type factory interface {
	CreateLink(caption, url string) link
	CreateTray(caption string) tray
	CreatePage(title, author string) page
}

//=============================
//产品实现
type mdLink struct {
	caption, url string
}

func (self *mdLink) toString() string {
	return "[" + self.caption + "](" + self.url + ")"
}

type mdTray struct {
	tray    []item
	caption string
}

func (self *mdTray) AddToTray(item item) {
	self.tray = append(self.tray, item)
}

func (self *mdTray) toString() string {
	tray := "- " + self.caption + "\n"
	for _, item := range self.tray {
		tray += item.toString() + "\n"
	}
	return tray
}

type basePage struct {
}

type mdPage struct {
	content       []item
	title, author string
}

func (self *mdPage) AddToContent(item item) {
	self.content = append(self.content, item)
}

func (self *mdPage) Output() string {
	content := "title: " + self.title + "\n"
	content += "author: " + self.author + "\n"
	for _, item := range self.content {
		content += item.toString() + "\n"
	}
	return content
}

//============================
// 具体的MD工厂
type MdFactory struct {
}

func (self *MdFactory) CreateLink(caption, url string) link {
	return &mdLink{caption, url}
}
func (self *MdFactory) CreateTray(caption string) tray {
	return &mdTray{caption: caption}
}
func (self *MdFactory) CreatePage(title, author string) page {
	return &mdPage{title: title, author: author}
}
