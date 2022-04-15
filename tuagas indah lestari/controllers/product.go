package controllers

import (
	"myapp/dto"
	"myapp/interfaces"
	"myapp/middlewares"
	"myapp/models"
	services "myapp/services/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ProductController interfaces.ProductInterface
)

func init() {
	ProductController = new(productController)
}

type productController struct{}


// @Summary Get All product
// @Description Get All product type, return list of products
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.Response
// @Failure 401
// @Failure 500
// @Router /product [get]
func (c *productController) GetAll(ctx *gin.Context) {

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	products, err := services.Database.ProductGetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data:   products,
		Error:  "",
	})
}


// @Summary Get certain product by id
// @Description Get certain product by id
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param id path string true "product ID"
// @Success 200
// @Failure 500
// @Router /product/{id} [get]
func (c *productController) GetByID(ctx *gin.Context) {

	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	product, err := services.Database.ProductGetByID(productId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data:   product,
		Error:  "",
	})
}


// @Summary Create New product
// @Description Create New product
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param AdminData body dto.ProductParam true "product type data"
// @Success 200 {object} dto.Response
// @Failure 500
// @Router /product [post]
func (c *productController) Create(ctx *gin.Context) {

	var newProduct dto.ProductParam
	err := ctx.ShouldBind(&newProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	product, err := services.Database.ProductCreate(models.Product{
		Code: newProduct.Code,
		ProductTypeId: newProduct.ProductTypeId,
		Name:   newProduct.Name,
		Description: newProduct.Description,
		Price: newProduct.Price,
		Qty: 0,
	})

	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data:   product,
		Error:  "",
	})
}

func (c *productController) Update(ctx *gin.Context) {

	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	var product dto.ProductUpdateParam
	err = ctx.ShouldBind(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	newUpdatedProduct, err := services.Database.ProductUpdate(models.Product{
		ID: productId,
		Code: product.Code,
		ProductTypeId: product.ProductTypeId,
		Name:   product.Name,
		Description: product.Description,
		Price: product.Price,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data:   newUpdatedProduct,
		Error:  "",
	})
}

func (c *productController) Delete(ctx *gin.Context) {

	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	tx := services.BeginTransaction()

	_, err = services.Database.ProductDelete(productId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data:   productId,
		Error:  "",
	})
}