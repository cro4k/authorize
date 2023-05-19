package perm

import (
	"testing"

	"github.com/cro4k/authorize/internal/db"
)

func TestCasbin(t *testing.T) {
	database := db.MustOpen(db.Config{
		Driver: "mysql",
		DSN:    "root:root@tcp(127.0.0.1:3306)/authorize?parseTime=true&charset=utf8&loc=Asia%2FShanghai",
	})
	s, err := NewService(database)
	if err != nil {
		t.Error(err)
	}
	//if err := s.AddPolicy("p", "p", []string{"alice", "book", "*"}); err != nil {
	//	t.Error(err)
	//}
	//if err := s.AddPolicy("p", "p", []string{"bob", "book", "read"}); err != nil {
	//	t.Error(err)
	//}
	//if err := s.AddPolicy("p", "p", []string{"joe", "book", "write"}); err != nil {
	//	t.Error(err)
	//}

	t.Log(s.enforce("alice", "book", "read"))
	t.Log(s.enforce("bob", "book", "read"))

}
