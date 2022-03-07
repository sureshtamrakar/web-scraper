package models_login

import (
	database_mysql "github.com/sureshtamrakar/web-scraper/database"
)

type Entity struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(email string) (login Entity, err error) {
	row := database_mysql.EsConn.Conn.QueryRow("SELECT email, password FROM `user` WHERE `email`=?", email)
	err = row.Scan(
		&login.Email,
		&login.Password,
	)
	return login, err
}

func Create(email, password string) error {
	stmt, _ := database_mysql.EsConn.Conn.Prepare("INSERT INTO user (email, password) VALUES (?,?)")
	_, err := stmt.Exec(email, password)
	if err != nil {
		return nil
	}
	stmt.Close()
	return nil
}
