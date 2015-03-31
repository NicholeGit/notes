package proxy

type printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(str string) string
}

type printer struct {
	name string
}

func (self *printer) SetPrinterName(name string) {
	self.name = name
}

func (self *printer) GetPrinterName() string {
	return self.name
}

func (self *printer) Print(str string) string {
	return self.name + ":" + str
}

// PrinterProxy代理printer,重写了所有接口
type PrinterProxy struct {
	Name string
	real *printer // 实际的printer对象
}

func (self *PrinterProxy) SetPrinterName(name string) {
	if self.real != nil {
		self.real.SetPrinterName(name)
	}
	self.Name = name
}

func (self *PrinterProxy) GetPrinterName() string {
	return self.Name
}

func (self *PrinterProxy) Print(str string) string {
	self.realize()
	return self.real.Print(str)
}

func (self *PrinterProxy) realize() {
	if self.real == nil {
		self.real = &printer{self.Name}
	}
}
