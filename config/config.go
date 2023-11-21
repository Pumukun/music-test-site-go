package config

import (
    "github.com/gorilla/sessions"
)

var (
    Store = sessions.NewCookieStore([]byte("m0Zako86s"))
)
