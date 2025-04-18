package db

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func Test(t *testing.T) {
	users := make([]Users, 0, 10)

	for k := 0; k < 10; k++ {
		token := uuid.New().String()
		nn := fmt.Sprintf("mx%d", k)
		n := fmt.Sprintf("ych%d", k)
		users = append(users, Users{
			Name:     n,
			Email:    "test",
			Password: "test",
			Token:    token,
			Nickname: nn,
		})
	}

	fmt.Println(BatchCreate(users))
}
