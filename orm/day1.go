package orm

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

type User1 struct {
	Name string
	Age  int
}

func Test1(user1 *User1, user2 User1) {
	uu1 := reflect.Indirect(reflect.ValueOf(user1)).Type()
	uu2 := reflect.ValueOf(user2).Type()
	fieldsCount := uu1.NumField()

	for i := 0; i < fieldsCount; i++ {
		fmt.Print(uu1.Field(i).Name, "|")
		fmt.Println(uu2.Field(i).Name)
	}
}

func Test2() {
	conn, err := sql.Open("sqlite3", "orm.db")
	if err != nil {
		log.Fatalf("connect datesource error: %v\n", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	_, _ = conn.Exec("DROP TABLE IF EXISTS User;")
	_, _ = conn.Exec("CREATE TABLE User(Name text);")

	result, err := conn.Exec("INSERT INTO User VALUES (?), (?);", "zhangsan", "lisi")
	if err != nil {
		log.Fatalf("exec sql statement error, %v\n", err)
	}

	line, _ := result.RowsAffected()
	log.Printf("insert affect %d line\n", line)

}
