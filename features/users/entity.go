package users

import "github.com/labstack/echo/v4"

type User struct {
	ID       uint
	Nama     string
	HP       string
	Password string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type Service interface {
	Register(newUser User) (User, error)
	Login(hp string, password string) (User, error)
}

type Repository interface {
	Insert(newUser User) (User, error)
	Login(hp string) (User, error)
}
