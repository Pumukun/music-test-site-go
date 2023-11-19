package handlers

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "main/models"
    "net/http"
    "strings"
)

func RegisterUser(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newUser models.User

        if err := c.BindJSON(&newUser); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        newUser.Username = strings.TrimSpace(newUser.Username)

        var username string
        err := db.QueryRow("SELECT username FROM users WHERE username = ?", newUser.Username).Scan(&username)
        if username != "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
            return
        }
        if strings.TrimSpace(newUser.Password) == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Password is empty"})
            return
        }

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hash error"})
            return
        }
        newUser.Password = string(hashedPassword)

        statement, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        _, err = statement.Exec(newUser.Username, newUser.Password)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
    }
}

func LoginUser(c *gin.Context) {
}
