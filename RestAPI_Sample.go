package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "empusr:emp@/emp_detail")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	type outputdata struct {
		ID    string
		Name  string
		DOJ   string
		Dept  string
		Grade string
	}
	router := gin.Default()
	router.GET("/getthruid/:id", func(c *gin.Context) {
		var (
			getthruid outputdata
			result    gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("SELECT * FROM emp_table WHERE id = ?;", id)
		err = row.Scan(&getthruid.ID, &getthruid.Name, &getthruid.DOJ, &getthruid.Dept, &getthruid.Grade)
		if err != nil {

			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": getthruid,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// GET multiple data
	router.GET("/getthrudept/:dept", func(c *gin.Context) {
		var (
			getthruid   outputdata
			getthrudept []outputdata
		)
		dept := c.Param("dept")
		rows, err := db.Query("SELECT * FROM emp_table WHERE dept = ?;", dept)
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&getthruid.ID, &getthruid.Name, &getthruid.DOJ, &getthruid.Dept, &getthruid.Grade)
			getthrudept = append(getthrudept, getthruid)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": getthrudept,
			"count":  len(getthrudept),
		})
	})

	router.Run(":8100")
}
