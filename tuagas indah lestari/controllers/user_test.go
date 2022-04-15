package controllers

import (
	"bytes"
	"encoding/json"
	"myapp/config"
	"myapp/dto"
	"myapp/middlewares"
	services "myapp/services/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func init() {
	config.SetEnv()
	config.ConnectGorm()
}

func Test_Register(t *testing.T) {
	router := gin.New()
	c := new(userController)
	router.POST("api/register", c.Register)

	t.Run("no body", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/register", nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := map[string]interface{}{
			"message": "invalid request",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("name required", func(t *testing.T) {
		var param = dto.UserRegisterParam{}
		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := map[string]interface{}{
			"message": "Name is a required field",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("email required", func(t *testing.T) {
		var param = dto.UserRegisterParam{
			Name: "Louis Aldorio",
		}
		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := map[string]interface{}{
			"message": "Email is a required field",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("password required", func(t *testing.T) {
		var param = dto.UserRegisterParam{
			Name:  "Louis Aldorio",
			Email: "louisaldorio@gmail.com",
		}
		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := map[string]interface{}{
			"message": "Password is a required field",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("register success", func(t *testing.T) {
		var param = dto.UserRegisterParam{
			Name:     "Louis Aldorio",
			Email:    "louisaldorio@gmail.com",
			Password: "12345",
		}
		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		tx := services.BeginTransaction()
		user, _ := services.Database.UserGetByEmail(param.Email)
		tx.Commit()

		expected := dto.Response{
			Status: true,
			Data: dto.UserTokenResponse{
				Token: middlewares.JwtGenerate(user.ID),
			},
			Error: "",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("Email already exist", func(t *testing.T) {
		var param = dto.UserRegisterParam{
			Name:     "Louis Aldorio",
			Email:    "louisaldorio@gmail.com",
			Password: "12345",
		}
		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Email already exists!",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})
}

func Test_Login(t *testing.T) {
	router := gin.New()
	UserController = new(userController)
	router.POST("api/login", UserController.Login)

	t.Run("login success", func(t *testing.T) {
		var param = dto.UserLoginParam{
			Email:    "louisaldorio@gmail.com",
			Password: "12345",
		}

		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		tx := services.BeginTransaction()
		user, _ := services.Database.UserGetByEmail(param.Email)
		tx.Commit()

		expected := dto.Response{
			Status: true,
			Data: dto.UserTokenResponse{
				Token: middlewares.JwtGenerate(user.ID),
			},
			Error: "",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("user not found", func(t *testing.T) {
		var param = dto.UserLoginParam{
			Email:    "emailasalasal@gmail.com",
			Password: "12345",
		}

		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := dto.Response{
			Status: false,
			Data:   nil,
			Error:  "user not available",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("invalid password", func(t *testing.T) {
		var param = dto.UserLoginParam{
			Email:    "louisaldorio@gmail.com",
			Password: "54321",
		}

		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Invalid Password!",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})
}
