package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"myapp/config"
	"myapp/dto"
	"myapp/middlewares"
	services "myapp/services/db"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func init() {
	config.SetEnv()
	config.ConnectGorm()

	tx := services.BeginTransaction()
	tx.Exec("INSERT INTO product_type(id, name) VALUES ('1', 'alat mandi')")
	tx.Exec("INSERT INTO product(id, code, product_type_id, name, description, price, qty) VALUES ('1', 'P001', '1', 'sabun', 'wangi', '10000', '1')")
	tx.Commit()
}

func Test_GetAll(t *testing.T) {
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	c := new(productController)
	router.GET("api/product", c.GetAll)

	t.Run("not logged in", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/product", nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("GetAll success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/product", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjUwMjEzMzQyLCJpYXQiOjE2NDk5NTQxNDJ9.Ig857yAX5KbDGBWrQvgu5jUVCkOgItlRDwTESNRCePA")

		router.ServeHTTP(w, req)

		tx := services.BeginTransaction()
		products, _ := services.Database.ProductGetAll()
		tx.Commit()

		expected := dto.Response{
			Status: true,
			Data:   products,
			Error:  "",
		}
		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})
}

func Test_GetByID(t *testing.T) {
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	c := new(productController)
	router.GET("api/product/:id", c.GetByID)

	t.Run("not logged in", func(t *testing.T) {
		var id = "1"
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/product/"+id, nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("record not found", func(t *testing.T) {
		var id = "9999"
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/product/"+id, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjUwMjEzMzQyLCJpYXQiOjE2NDk5NTQxNDJ9.Ig857yAX5KbDGBWrQvgu5jUVCkOgItlRDwTESNRCePA")

		router.ServeHTTP(w, req)

		tx := services.BeginTransaction()
		intID, _ := strconv.Atoi(id)
		_, err := services.Database.ProductGetByID(intID)
		tx.Commit()

		expected := dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 404, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("GetByID success", func(t *testing.T) {
		var id = "1"
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/product/"+id, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjUwMjEzMzQyLCJpYXQiOjE2NDk5NTQxNDJ9.Ig857yAX5KbDGBWrQvgu5jUVCkOgItlRDwTESNRCePA")

		router.ServeHTTP(w, req)

		tx := services.BeginTransaction()
		intID, _ := strconv.Atoi(id)
		product, _ := services.Database.ProductGetByID(intID)
		tx.Commit()

		expected := dto.Response{
			Status: true,
			Data:   product,
			Error:  "",
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})
}

func Test_Create(t *testing.T) {
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	c := new(productController)
	router.POST("api/product", c.Create)

	t.Run("not logged in", func(t *testing.T) {
		var param = dto.ProductParam{
			Code:          "P001",
			ProductTypeId: 1,
			Name:          "sabun",
			Description:   "wangi, bikin badan bersih",
			Price:         10000,
		}
		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/product", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("create success", func(t *testing.T) {
		var param = dto.ProductParam{
			Code:          "P001",
			ProductTypeId: 1,
			Name:          "sabun",
			Description:   "wangi, bikin badan bersih",
			Price:         10000,
		}
		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/product", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjUwMjEzMzQyLCJpYXQiOjE2NDk5NTQxNDJ9.Ig857yAX5KbDGBWrQvgu5jUVCkOgItlRDwTESNRCePA")

		router.ServeHTTP(w, req)

		response := dto.Response{}
		json.Unmarshal(w.Body.Bytes(), &response)
		productResp := response.Data.(map[string]interface{})

		tx := services.BeginTransaction()
		product, _ := services.Database.ProductGetByID(int(productResp["id"].(float64)))
		tx.Commit()

		product.CreatedAt = product.CreatedAt.UTC()
		product.UpdatedAt = product.UpdatedAt.UTC()

		expected := dto.Response{
			Status: true,
			Data:   product,
			Error:  "",
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})
}

func Test_Update(t *testing.T) {
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	c := new(productController)
	router.PUT("api/product/:id", c.Update)

	t.Run("not logged in", func(t *testing.T) {
		var id = "1"
		var param = dto.ProductUpdateParam{
			Code:          "P001",
			ProductTypeId: 1,
			Name:          "sabun",
			Description:   "wangi, bikin badan bersih + shining shimmering splendid",
			Price:         20000,
		}
		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/product/"+id, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("Update success", func(t *testing.T) {
		var id = "1"
		var param = dto.ProductUpdateParam{
			Code:          "P001",
			ProductTypeId: 1,
			Name:          "sabun",
			Description:   "wangi, bikin badan bersih + shining shimmering splendid",
			Price:         20000,
		}
		body, _ := json.Marshal(param)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/product/"+id, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjUwMjEzMzQyLCJpYXQiOjE2NDk5NTQxNDJ9.Ig857yAX5KbDGBWrQvgu5jUVCkOgItlRDwTESNRCePA")

		router.ServeHTTP(w, req)

		intID, _ := strconv.Atoi(id)
		tx := services.BeginTransaction()
		product, _ := services.Database.ProductGetByID(intID)
		tx.Commit()

		fmt.Println(product.CreatedAt)
		fmt.Println(product.CreatedAt.Local())
		fmt.Println(product.CreatedAt.UTC())

		expected := dto.Response{
			Status: true,
			Data:   product,
			Error:  "",
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})
}

func Test_Delete(t *testing.T) {
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	c := new(productController)
	router.DELETE("api/product/:id", c.Delete)

	t.Run("not logged in", func(t *testing.T) {
		var id = "1"

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/product/"+id, nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		expected := dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 401, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})

	t.Run("Delete success", func(t *testing.T) {
		var id = "1"

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/product/"+id, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjUwMjEzMzQyLCJpYXQiOjE2NDk5NTQxNDJ9.Ig857yAX5KbDGBWrQvgu5jUVCkOgItlRDwTESNRCePA")

		router.ServeHTTP(w, req)

		intID, _ := strconv.Atoi(id)
		expected := dto.Response{
			Status: true,
			Data:   intID,
			Error:  "",
		}

		expctedByte, _ := json.Marshal(expected)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expctedByte), w.Body.String())
	})
}
