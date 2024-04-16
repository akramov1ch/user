package functions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	st "github.com/akramov1ch/user/storage"
)

func UserInfo(users *[]st.User) {
	file, err := os.Open("github.com/akramov1ch/user/sample.txt")
	if err != nil {
		fmt.Println("Fileni ochishda xatolik yuz berdi:", err)
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var user st.User

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			*users = append(*users, user)
			user = st.User{}
			continue
		}
		data := strings.Split(line, ":")
		if len(data) < 2 { 
			fmt.Println("Ma'lumotlar noto'g'ri formatda:", line)
			continue
		}
		key, value := strings.TrimSpace(data[0]), strings.TrimSpace(data[1])
		switch key {
		case "Name":
			user.Name = value
		case "Age":
			age, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Yosh xato kiritilgan:", err)
				continue
			}
			user.Age = age
		case "Occupation":
			user.Occupation = value
		}
	}
	if user.Name != "" {
		*users = append(*users, user)
	}
}

func Printer(users []st.User) {
	for i, usr := range users {
		fmt.Printf("User %d: \n", i+1)
		fmt.Printf("Name: %s\n", usr.Name)
		fmt.Printf("Age: %d\n", usr.Age)
		fmt.Printf("Occupation: %s\n\n", usr.Occupation)
	}
}
