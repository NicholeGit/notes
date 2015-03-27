package observer

type numberGenerator struct {
	observers []observer
}

func (self *numberGenerator) AddObserver(observer observer) {
	self.observers = append(self.observers, observer)
}

func (self *numberGenerator) notifyObservers() []int {
	var result []int
	for _, observer := range self.observers {
		result = append(result, observer.update())
	}
	return result
}

type randomNumberGenerator struct {
	*numberGenerator
}

func NewRandomNumberGenerator() *randomNumberGenerator {
	return &randomNumberGenerator{&numberGenerator{}}
}

func (self *randomNumberGenerator) Execute() []int {
	return self.notifyObservers()
}

type observer interface {
	update() int
}

type DigitObserver struct {
	generator int
}

func (self *DigitObserver) update() int {
	return self.generator
}
