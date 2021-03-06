package mysql

import (
	"github.com/brendensoares/storm"
	"testing"
)

type User struct {
	storm.Model `container:"users"`
	Id int64
	FirstName string
	LastName string
	Email string
}

func NewUser() *User {
	return storm.Factory(&User{}).(*User)
}

func TestMysqlCreate(t *testing.T) {
	if commError := storm.Connect("mysql", "user:@unix(/var/run/mysqld/mysqld.sock)/dbname"); commError != nil {
		// Failure
		t.Fatal("Database error")
	} else {
		// Success, create new user
		newUser := NewUser()
		newUser.Email = "brenden@test.com"
		if saveError := newUser.Save(); saveError != nil {
			// Failure
			t.Fatalf("%s %s", "Query error", saveError)
		} else {
			println("newUser id:", newUser.Id)
		}
	}
}
