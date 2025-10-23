package test

import (
	"fmt"
	"testing"
	"yves-go/config"
)

type Person struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Mobile   string `db:"mobile"`
}

func TestQuery(t *testing.T) {
	var person []Person
	err := config.Db.Select(&person, "select id, username,mobile from user_info where id < ?", 255)
	defer func() {
		_ = config.Db.Close()
	}()

	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	for _, v := range person {
		fmt.Println("select succ:", v)
	}
}
