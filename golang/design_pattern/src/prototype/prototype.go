package prototype

type producter interface {
	clone() producter
	GetName() string
}

type Manager struct {
	productMap map[string]producter
}

func (self *Manager) Register(pdter producter) {
	if self.productMap == nil {
		self.productMap = make(map[string]producter)
	}
	self.productMap[pdter.GetName()] = pdter
}

func (self *Manager) Create(name string) producter {
	if v, ok := self.productMap[name]; ok {
		return v
	}
	return nil
}

type Product struct {
	name string
}

func (self *Product) SetUp() {
	// something takes time...
}

func (self *Product) GetName() string {
	return self.name
}

func (self *Product) clone() producter {
	return &Product{self.name}
}
