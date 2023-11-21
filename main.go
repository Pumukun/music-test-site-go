package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    //"fmt"
    "log"
    "main/handlers"
    "strings"
    "os"
    "main/config"
)

func initDB() *sql.DB {
    db, err := sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatal(err)
    }

    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,     
        "username" TEXT,
        "password" TEXT
    );`

    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    return db
}

func main() {
    db := initDB()
    defer db.Close()

    router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
    router.Static("/static", "static/")
    
	router.GET("/", func(c *gin.Context) {
	    c.HTML(http.StatusOK, "index", gin.H{})
	})

    router.GET("/index", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index", gin.H{})
    })

    router.GET("/register", func(c *gin.Context) {
        c.HTML(http.StatusOK, "register", gin.H{})
    })
    router.POST("/register", handlers.RegisterUser(db))

    router.GET("/logout", func(c *gin.Context) {
        session, _ := config.Store.Get(c.Request, "session")

        delete(session.Values, "user_id")
        delete(session.Values, "authenticated")

        session.Save(c.Request, c.Writer)
        c.Redirect(http.StatusFound, "/index")
    })

    router.GET("/test", func(c *gin.Context) {
        c.HTML(http.StatusOK, "test", gin.H{})
    })
    
    router.GET("/test2", func(c *gin.Context) {
        c.HTML(http.StatusOK, "test2", gin.H{})
    })

    router.GET("/api/get_audio_files/", func(c *gin.Context) {
        staticDir := "./static/"

        var mp3Files []string

        files, err := os.ReadDir(staticDir)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read directory"})
            return
        }

        for _, file := range files {
            if !file.IsDir() && strings.HasSuffix(file.Name(), ".mp3") {
                mp3Files = append(mp3Files, file.Name())
            }
        }

        c.JSON(http.StatusOK, gin.H{"mp3_files": mp3Files})
    })

	router.Run(":8080")
}
