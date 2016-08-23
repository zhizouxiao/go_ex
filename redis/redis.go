package main

import (
	//"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

// Dial dials the local Redis server and selects database 9. To prevent
// stomping on real data, DialTestDB fails if database 9 contains data. The
// returned connection flushes database 9 on close.
func Dial() (redis.Conn, error) {
	c, err := redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		return nil, err
	}

	_, err = c.Do("SELECT", "9")
	if err != nil {
		c.Close()
		return nil, err
	}

	//n, err := redis.Int(c.Do("DBSIZE"))
	//if err != nil {
	//c.Close()
	//return nil, err
	//}

	//if n != 0 {
	//c.Close()
	//return nil, errors.New("database #9 is not empty, test can not continue")
	//}

	return c, nil
}
func main() {
	conn, err := Dial()
	//res, err := conn.Do("SELECT", 1)
	if err != nil {
		conn.Close()
		fmt.Println(err)
	}
	//fmt.Println(res)
	res1, err := conn.Do("GET", "a")
	if err != nil {
		conn.Close()
		fmt.Println(err)
	}
	fmt.Println(res1)
	conn.Close()

}
