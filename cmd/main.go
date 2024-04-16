package main

import (
	fnc "user/internal/functions"
	st "user/storage"
)

func main() {
	var users []st.User
	fnc.UserInfo(&users)
	fnc.Printer(users)
}
