package db

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"time"
	. "types"

	"cfg"
)

const (
	COLLECTION   = "USERS"
	COUNTER_NAME = "USERID_GEN"
)

func LoginMac(name, mac string) *User {
	ms, c := C(COLLECTION)
	defer ms.Close()

	user := &User{}
	err := c.Find(bson.M{"name": name, "mac": mac}).One(user)
	if err != nil {
		log.Println(err, mac)
		return nil
	}

	return user
}

func Query(name string) *User {
	ms, c := C(COLLECTION)
	defer ms.Close()

	user := &User{}
	err := c.Find(bson.M{"name": name}).One(user)
	if err != nil {
		log.Println(err, name)
		return nil
	}

	return user
}

func Get(id int32) *User {
	ms, c := C(COLLECTION)
	defer ms.Close()

	user := &User{}
	err := c.Find(bson.M{"id": id}).One(user)
	if err != nil {
		log.Println(err, id)
		return nil
	}

	return user
}

func GetAll() []User {
	ms, c := C(COLLECTION)
	defer ms.Close()

	var users []User
	err := c.Find(nil).All(&users)
	if err != nil {
		log.Println(err)
		return nil
	}

	return users
}

func New(name, mac string) *User {
	ms, c := C(COLLECTION)
	defer ms.Close()

	config := cfg.Get()
	user := &User{}
	err := c.Find(bson.M{"name": name}).One(user)
	if err != nil {
		user.Id = NextVal(COUNTER_NAME)
		user.Name = name
		user.Mac = mac
		user.Domain = config["domain"]
		user.CreatedAt = time.Now().Unix()
		err := c.Insert(user)
		if err != nil {
			log.Println(err, name, mac)
			return nil
		}
		return user
	}

	return nil
}
