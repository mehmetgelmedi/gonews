package repositories

import (
	"gonews/db"
	"gonews/models"
)

func GetUser(userLogin models.UserLogin) models.User {
	db := db.Connect()
	defer db.Close()

	row, err := db.Query("select * from public.user where username=$1 and password=$2", userLogin.Username, userLogin.Password)
	if err != nil {
		panic(err.Error())
	}
	var user models.User
	for row.Next() {
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.IsAdmin)
		if err != nil {
			panic(err.Error())
		}
	}

	return user
}

func CreateUser(user *models.User) *models.User {
	db := db.Connect()
	defer db.Close()

	sqlStatement := `INSERT INTO public.user (username, password, is_admin) VALUES ($1, $2, $3) RETURNING id`

	err := db.QueryRow(sqlStatement, user.Username, user.Password, user.IsAdmin).Scan(&user.ID)
	if err != nil {
		panic(err)
	}

	return user
}
