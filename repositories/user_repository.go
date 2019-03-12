package repositories

import(
	"gonews/models"
	"gonews/db"
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
		err:=row.Scan(&user.ID, &user.Username, &user.Password, &user.IsAdmin)
		if err != nil {
			panic(err.Error())
		}
	}

	return user
}