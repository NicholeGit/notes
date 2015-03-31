package memento

type memento struct {
	money int
}

func (self *memento) getMoney() int {
	return self.money
}

type Game struct {
	Money int
}

// 建立一个备份
func (self *Game) CreateMemento() *memento {
	return &memento{
		self.Money,
	}
}

// 还原备份
func (self *Game) RestoreMemento(memento *memento) {
	self.Money = memento.getMoney()
}
