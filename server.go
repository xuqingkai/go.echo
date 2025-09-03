package main

import (
    "os"
    "database/sql"
    "log"
	"net/http"
	"github.com/labstack/echo/v4"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)
import ()


type User struct {
    ID       int    `db:"id" json:"id"`
    Name     string `db:"test_key" json:"name"`
}
var db *sql.DB
func main() {

    err1 := godotenv.Load()
    if err1 != nil { log.Fatal("Error loading .env file") }
    dbHost := os.Getenv("DB_HOST")

	echoServer := echo.New()
    echoServer.Static("/static", "assets")

    var err error
    db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // 测试连接
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    
    
	echoServer.GET("/", func(context echo.Context) error {
		return context.HTML(http.StatusOK, dbHost + "<b><img src=\"./static/1.jpg\">Hello, World!</b>")
	})
    echoServer.GET("/users", getUsers)
	echoServer.Logger.Fatal(echoServer.Start(":1323"))
}
func getUsers(c echo.Context) error {
    rows, err := db.Query("SELECT id, test_key FROM ylyq_peis_test")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "查询失败",
        })
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "数据解析失败",
            })
        }
        users = append(users, user)
    }
    
    return c.JSON(http.StatusOK, users)
}