package strategy

import (
	"math/rand"
)

const (
	GUU = iota
	CHO
	PAA
)

var hands []*hand

func init() {
	hands = []*hand{
		&hand{GUU},
		&hand{CHO},
		&hand{PAA},
	}
}

type hand struct {
	handValue int
}

func getHand(handValue int) *hand {
	return hands[handValue]
}

func (self *hand) IsStrongerThan(h *hand) bool {
	return self.fight(h) == 1
}

func (self *hand) IsWeakerThan(h *hand) bool {
	return self.fight(h) == -1
}

func (self *hand) fight(h *hand) int {
	if self == h {
		return 0
	} else if (self.handValue+1)%3 == h.handValue {
		return 1
	} else {
		return -1
	}
}

//策略接口
type strategy interface {
	NextHand() *hand
	study(win bool)
}

// 一个具体的策略，只要实现strategy接口就是一个具体实现
// 策略是可以被多个类使用。（这点是模板方法模式（Template Method）最大区别）
type winningStrategy struct {
	seed     int64
	won      bool
	prevHand *hand
}

func (self *winningStrategy) NextHand() *hand {
	if !self.won {
		// rand.Seed(self.seed)
		self.prevHand = getHand(rand.Intn(3))
	}
	return self.prevHand
}

func (self *winningStrategy) study(win bool) {
	self.won = win
}

// 实际使用策略类
type Player struct {
	Name                           string
	Strategy                       strategy
	wincount, losecount, gamecount int
}

func (self *Player) NextHand() *hand {
	return self.Strategy.NextHand()
}

func (self *Player) Win() {
	self.wincount++
	self.gamecount++
}

func (self *Player) Lose() {
	self.losecount++
	self.gamecount++
}

func (self *Player) Even() {
	self.gamecount++
}
