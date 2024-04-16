package main

import (
	fnc "github.com/akramov1ch/user/internal/functions"
	st "github.com/akramov1ch/user/storage"
)

func main() {
	var users []st.User
	fnc.UserInfo(&users)
	fnc.Printer(users)
}
