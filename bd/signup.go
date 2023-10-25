package bd

import (
	"fmt"
	"gambit/models"
	"gambit/tools"
)

func SignUp(sig models.SignUp) error {
	println("Comienza registro")
	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	setencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySql() + "')"
	fmt.Println(setencia)

	_, err = Db.Exec(setencia)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	fmt.Println("SignUp > Ejecuciòn existosa")
	return nil
}
