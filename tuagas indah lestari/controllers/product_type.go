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
	ProductTypeController interfaces.ProductTypeInterface
)

func init() {
	ProductTypeController = new(productTypeController)
}

type productTypeController struct{}


// @Summary Get All product type
// @Description Get All product type, return list of product types
// @Tags product_type
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.Response
// @Failure 401
// @Failure 500
// @Router /product-type [get]
func (c *productTypeController) GetAll(ctx *gin.Context) {

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

	productTypes, err := services.Database.ProductTypeGetAll()
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
		Data:   productTypes,
		Error:  "",
	})
}


// @Summary Get certain product type by id
// @Description Get certain product type by id
// @Tags product_type
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param id path string true "product type ID"
// @Success 200
// @Failure 500
// @Router /product-type/{id} [get]
func (c *productTypeController) GetByID(ctx *gin.Context) {

	productTypeId, err := strconv.Atoi(ctx.Param("id"))
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

	productType, err := services.Database.ProductTypeGetByID(productTypeId)
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
		Data:   productType,
		Error:  "",
	})
}


// @Summary Create New product type
// @Description Create New product type
// @Tags product_type
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param AdminData body dto.ProductTypeParam true "product type data"
// @Success 200 {object} dto.Response
// @Failure 500
// @Router /product-type [post]
func (c *productTypeController) Create(ctx *gin.Context) {

	var newProductType dto.ProductTypeParam
	err := ctx.ShouldBind(&newProductType)
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

	productType, err := services.Database.ProductTypeCreate(models.ProductType{
		Name:   newProductType.Name,
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
		Data:   productType,
		Error:  "",
	})
}

func (c *productTypeController) Update(ctx *gin.Context) {

	productTypeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	var productType dto.ProductTypeUpdateParam
	err = ctx.ShouldBind(&productType)
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

	newUpdatedProductType, err := services.Database.ProductTypeUpdate(models.ProductType{
		ID:     productTypeId,
		Name:   productType.Name,
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
		Data:   newUpdatedProductType,
		Error:  "",
	})
}

func (c *productTypeController) Delete(ctx *gin.Context) {

	productTypeId, err := strconv.Atoi(ctx.Param("id"))
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

	_, err = services.Database.ProductTypeDelete(productTypeId)
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
		Data:   productTypeId,
		Error:  "",
	})
}