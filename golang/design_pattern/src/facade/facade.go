package facade

var db = map[string]string{
	"a@a.com": "a",
	"b@b.com": "b",
}

type database struct {
}

func (self *database) getNameByMail(mail string) string {
	return db[mail]
}

// 对md文件来说 只通过通过mdWriter来访问database。mdWriter作为一个装饰或封装。
type mdWriter struct {
}

func (self *mdWriter) title(title string) string {
	return "# Welcome to " + title + "'s page!"
}

type PageMaker struct {
}

func (self *PageMaker) MakeWelcomePage(mail string) string {
	database := database{}
	writer := mdWriter{}

	name := database.getNameByMail(mail)
	page := writer.title(name)

	return page
}
