package mgo
import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Ts struct {
	Id 		bson.ObjectId 	`bson:"_id,omitempty"`
	Name 	string			`bson:"name"`
}

var e = &Exec{
	Database: "deep-space",
	Username: "root",
	Password: "guhao19890412",

	Query: make(map[string]interface{}),
}

func Test_Mgo(t *testing.T) {
	e.Collection = "test"

	var ts Ts
	e.Query = bson.M{"name":"test4"}
	err := e.Find(&ts)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ts.Name)

	/*ts1 := &Ts{Name: "test1"}
	ts2 := &Ts{Name: "test3"}
	ts3 := &Ts{Name: "test4"}
	ts5 := &Ts{Name: "test5"}

	err := e.Insert(ts1, ts2,ts3,ts5)

	fmt.Println(err)*/

	/*e.Query = bson.M{"name": bson.M{"$in": []string{"test3","test4"}}}

	err := e.Remove()
	if err != nil {
		fmt.Println(err)
	}*/

	//var ts []Ts

	/*e.Sort = []string{"-name"}

	err := e.FindAll(&ts)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ts)*/

	/*session := Session()
	defer session.Close()

	c := session.DB("tianshu").C("test")

	var ts Ts

	c.Find(bson.M{"name": "test1"}).One(&ts)

	fmt.Println(ts)*/

}
