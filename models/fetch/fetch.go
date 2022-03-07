package models_fetch

import (
	"fmt"

	database_mysql "github.com/sureshtamrakar/web-scraper/database"
)

type Entity struct {
	Id         int    `json:"id"`
	URL        string `json:"url"`
	Content    string `json:"content"`
	Created_At string `json:"created_at"`
}

func Create(url, content string) error {

	stmt, _ := database_mysql.EsConn.Conn.Prepare("INSERT INTO scrap (url, content) VALUES (?,?)")
	_, err := stmt.Exec(url, content)
	stmt.Close()
	return err

}

func LoadAll() (contents []Entity, err error) {
	stmt, err := database_mysql.EsConn.Conn.Query("SELECT * FROM `scrap`")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content Entity

		err = stmt.Scan(
			&content.Id,
			&content.URL,
			&content.Content,
			&content.Created_At,
		)

		contents = append(contents, content)
	}

	return contents, err
}

func Load(domain string) (value Entity, err error) {
	row := database_mysql.EsConn.Conn.QueryRow("SELECT * FROM `scrap` WHERE `url`=?", domain)
	err = row.Scan(
		&value.Id,
		&value.URL,
		&value.Content,
		&value.Created_At,
	)
	return value, err
}

func Paginate(cpage int) (contents []Entity, err error) {
	perPage := 2
	sqlPaging := "SELECT * FROM scrap"
	sqlPaging = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlPaging, perPage, (cpage-1)*perPage)
	fmt.Println(sqlPaging)
	stmt, err := database_mysql.EsConn.Conn.Query(sqlPaging)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content Entity

		err = stmt.Scan(
			&content.Id,
			&content.URL,
			&content.Content,
			&content.Created_At,
		)

		contents = append(contents, content)
	}

	return contents, err
}

func Search(start, end string) (contents []Entity, err error) {
	sqlPaging := fmt.Sprintf("SELECT * FROM `scrap` WHERE `created_at` BETWEEN '%s' AND '%s'", start, end)
	stmt, err := database_mysql.EsConn.Conn.Query(sqlPaging)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content Entity
		err = stmt.Scan(
			&content.Id,
			&content.URL,
			&content.Content,
			&content.Created_At,
		)
		contents = append(contents, content)
	}
	return contents, err
}
