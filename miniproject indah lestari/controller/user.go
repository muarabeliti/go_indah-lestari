package controller

import (
	"net/http"
	"project/config"
	"project/lib/database"
	"project/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {

	// var users []model.User
	users, e := database.GetUser()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {

	id := c.Param("id")
	user := model.User{}

	config.DB.Where("ID = ?", id).First(&user)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, user)
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {

	id := c.Param("id")

	config.DB.Delete(&model.User{}, id)

	return c.JSON(http.StatusOK, nil)
}

// update user by id
func UpdateUserController(c echo.Context) error {

	id := c.Param("id")
	user := model.User{}

	config.DB.Where("ID = ?", id).First(&user)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	payload := model.User{}
	c.Bind(&payload)

	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = payload.Password
	config.DB.Save(&user)

	return c.JSON(http.StatusOK, user)
}

func GetUserDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	user, err := database.GetDetailUser((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "sukses",
		"user ":  user,
	})
}

// func LoginUserController(c echo.Context) error {

// 	user := model.User{}
// 	c.Bind(&user)

// 	user, e := database.LoginUser(user)
// 	if e != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest,e.Error)
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success login",
// 		"user":    user,
// 	})

// }

// func LoginCon(echo echo.Context)error{
// 	user := model.User{}
// 	echo.Bind(&user)

// 	user,err := database.Login(&user)
// 	if err !=nil{
// 		return echo.JSON(http.StatusBadRequest,map[string]interface{}{
// 			"message":"fail login",
// 			"error" : err.Error(),
// 		})
// 	}
// 	return echo.JSON(http.StatusOK,user)
// }
