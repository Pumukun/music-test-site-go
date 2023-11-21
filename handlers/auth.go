package handlers

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "main/models"
    "net/http"
    "strings"
    "main/config"
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
            var storedUser models.User
            err := db.QueryRow("SELECT username, password FROM users WHERE username = ?", newUser.Username).Scan(&storedUser.Username, &storedUser.Password)
            if err != nil {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "SQL query error"})
                return
            }

            pwdErr := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(newUser.Password))
            if pwdErr == nil {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
                return
            }

            session, _ := config.Store.Get(c.Request, "session")
            session.Values["authenticated"] = true
            session.Values["user_id"] = newUser.ID
            session.Save(c.Request, c.Writer)

            c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
            return
        }
        if strings.TrimSpace(newUser.Password) == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Password is empty"})
            return
        }

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
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

