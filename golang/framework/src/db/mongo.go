package db

import (
	"cfg"
	. "helper"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var _global_ms *mgo.Session

const (
	COUNTERS = "COUNTERS"
)

func init() {
	config := cfg.Get()
	// dial mongodb
	sess, err := mgo.Dial(config["mongo_host"])
	if err != nil {
		ERR("cannot connect to", config["mongo_host"], err)
		os.Exit(-1)
	}

	// set default session mode to strong for saving player's data
	sess.SetMode(mgo.Strong, true)
	_global_ms = sess

}

//------------------------------------------------ copy connection
// !IMPORTANT!  NEVER FORGET -----> defer ms.Close() <-----
func C(collection string) (*mgo.Session, *mgo.Collection) {
	config := cfg.Get()
	ms := _global_ms.Copy()
	c := ms.DB(config["mongo_db"]).C(collection)
	return ms, c
}

type Counter struct {
	NextVal int64 /*`bson:"nextVal"`*/
}

//---------------------------------------------------------- ID GENERATOR
func NextVal(countername string) int32 {
	ms, c := C(COUNTERS)
	defer ms.Close()
	err := c.Insert(bson.M{
		"_id":     countername,
		"nextVal": 0,
	})
	if err != nil {
		DEBUG(err)
	}
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"nextVal": 1}},
		ReturnNew: true,
	}

	next := &Counter{}
	info, err := c.Find(bson.M{"_id": countername}).Apply(change, &next)
	if err != nil {
		ERR(info, err)
		return -1
	}
	// round the nextval to 2^31
	return int32(next.NextVal % 2147483648)
}
