package user

import (
	"net/http"
	"sampleprj/features/users"
	"sampleprj/helper/jwt"
	"sampleprj/helper/responses"
	"strings"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type userController struct {
	srv users.Service
}

func New(s users.Service) users.Handler {
	return &userController{
		srv: s,
	}
}

func (uc *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(UserRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(users.User)
		inputProcess.Nama = input.Nama
		inputProcess.Password = input.Password
		inputProcess.HP = input.Hp

		result, err := uc.srv.Register(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar ada sistem"
			}

			return responses.PrintResponse(c, statusCode, message, nil)
		}

		var response = new(UserResponse)
		response.Hp = result.HP
		response.Nama = result.Nama
		response.ID = result.ID

		return responses.PrintResponse(c, http.StatusCreated, "success create data", response)
	}
}

func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		result, err := uc.srv.Login(input.Hp, input.Password)

		if err != nil {
			c.Logger().Error("ERROR Login, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		strToken, err := jwt.GenerateJWT(result.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika mengenkripsi data",
			})
		}

		var response = new(LoginResponse)
		response.Hp = result.HP
		response.Nama = result.Nama
		response.ID = result.ID
		response.Token = strToken

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (uc *userController) GetListUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, err := jwt.ExtractToken(c.Get("user").(*gojwt.Token))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"message": "tidak ada kuasa untuk mengakses",
			})
		}

		return c.JSON(http.StatusOK, userid)
	}
}
