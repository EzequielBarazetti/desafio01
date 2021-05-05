package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func main() {
	db, err1 := sql.Open("mysql", "root:root@tcp(db:3306)/db_mysql")

	if err1 != nil {
		fmt.Println("Deu erro na conexao com o DB 1")
		panic(err1.Error())
	}
	defer db.Close()

	results, err2 := db.Query("SELECT id, nome FROM cursos")
	if err2 != nil {
		fmt.Println("Deu erro na conexao com o DB 2")
		panic(err2.Error())
	}

	s := make([]string, 0)

	for results.Next() {
		var tag Tag
		err3 := results.Scan(&tag.ID, &tag.Nome)
		if err3 != nil {
			panic(err3.Error())
		}

		s = append(s, tag.Nome)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, s)
	})
	r.Run(":8080")
}
