package decorator

type display interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
	Show(display display) string
}

type defaultDisplay struct {
	display
}

//因为golang里面没有多态机制，所以这里必须把display传入。
//否则defaultDisplay中并没有实现getRows和getRowText方法
func (self *defaultDisplay) Show(display display) string {
	str := ""
	for i := 0; i < display.getRows(); i++ {
		str += display.getRowText(i)
	}
	return str
}

type StringDisplay struct {
	*defaultDisplay
	str string
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		&defaultDisplay{},
		str,
	}
}

func (self *StringDisplay) getColumns() int {
	return len(self.str)
}

func (self *StringDisplay) getRows() int {
	return 1
}

func (self *StringDisplay) getRowText(row int) string {
	if row == 0 {
		return self.str
	} else {
		return ""
	}
}

type border struct {
	*defaultDisplay
	display display
}

// SideBorder装饰了border把一些方法重写
type SideBorder struct {
	*border
	borderChar string
}

func NewSideBorder(display display, borderChar string) *SideBorder {
	return &SideBorder{
		&border{display: display},
		borderChar,
	}
}

func (self *SideBorder) getColumns() int {
	return len(self.borderChar)*2 + self.display.getColumns()
}

func (self *SideBorder) getRows() int {
	return self.display.getRows()
}

func (self *SideBorder) getRowText(row int) string {
	return self.borderChar + self.display.getRowText(row) + self.borderChar
}
