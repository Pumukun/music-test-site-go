package models

import (
    _ "database/sql"
)

type User struct {
    ID int
    Username string
    Password string
}
