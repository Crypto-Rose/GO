package main

import (
	"fmt"
	"time"

	us "./user"
)

type crypto struct {
	us.User
}

func main() {
	u := new(crypto)
	u.HighUser(1, "Rose", time.Now(), true)
	fmt.Println(u.User)
}
