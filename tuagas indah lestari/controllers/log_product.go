package controllers

import (
	"fmt"
	"myapp/dto"
	"myapp/interfaces"
	"myapp/middlewares"
	"myapp/models"
	services "myapp/services/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	LogProductController interfaces.LogProductInterface
)

func init() {
	LogProductController = new(logProductController)
}

type logProductController struct{}

// @Summary Create log product with increase type
// @Description Create log product with increase type
// @Tags log_product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.Response
// @Failure 401
// @Failure 400
// @Failure 500
// @Router /log-product/increase [post]
func (c *logProductController) CreateFromIncrease(ctx *gin.Context) {

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	var newLogProduct dto.CreateFromIncreaseParam
	err := ctx.ShouldBind(&newLogProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	productTx := services.BeginTransaction()

	product, _ := services.Database.ProductGetByIDForLogProduct(newLogProduct.ProductId)

	productTx.Commit()

	if product == nil {
		newProductTx := services.BeginTransaction()

		newProduct, newProductErr := services.Database.ProductCreate(models.Product{
			Code:          newLogProduct.ProductCode,
			ProductTypeId: newLogProduct.ProductTypeId,
			Name:          newLogProduct.ProductName,
			Qty:           newLogProduct.ProductQty,
			Price:         newLogProduct.ProductPrice,
		})
		if newProductErr != nil {
			ctx.JSON(http.StatusNotFound, dto.Response{
				Status: false,
				Data:   nil,
				Error:  newProductErr.Error(),
			})
			return
		}
		product = newProduct

		if newProductTxErr := newProductTx.Commit().Error; newProductTxErr != nil {
			ctx.AbortWithError(http.StatusInternalServerError, newProductTxErr)
			return
		}
	} else {
		productUpdateTx := services.BeginTransaction()

		services.Database.ProductUpdate(models.Product{
			ID:            product.ID,
			Code:          newLogProduct.ProductCode,
			ProductTypeId: newLogProduct.ProductTypeId,
			Name:          newLogProduct.ProductName,
			Description:   product.Description,
			Price:         newLogProduct.ProductPrice,
			Qty:           newLogProduct.ProductQty + product.Qty,
		})

		if productUpdateTxErr := productUpdateTx.Commit().Error; productUpdateTxErr != nil {
			ctx.AbortWithError(http.StatusInternalServerError, productUpdateTxErr)
			return
		}
	}

	logProductCreateTx := services.BeginTransaction()

	logProduct, err := services.Database.LogProductCreate(models.LogProduct{
		Status:     "masuk",
		TotalQty:   newLogProduct.ProductQty,
		TotalPrice: 1.20 * (newLogProduct.ProductQty * product.Price),
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if logProductCreateTxErr := logProductCreateTx.Commit().Error; logProductCreateTxErr != nil {
		ctx.AbortWithError(http.StatusInternalServerError, logProductCreateTxErr)
		return
	}

	pivotCreateTx := services.BeginTransaction()

	pivot, err := services.Database.PivotCreate(models.Pivot{
		IDProduct:    product.ID,
		IDLogProduct: logProduct.ID,
		Qty:          newLogProduct.ProductQty,
		Price:        product.Price,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}
	fmt.Println(pivot)

	if pivotCreateTxErr := pivotCreateTx.Commit().Error; pivotCreateTxErr != nil {
		ctx.AbortWithError(http.StatusInternalServerError, pivotCreateTxErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data:   logProduct,
		Error:  "",
	})
}

// @Summary Create log product with decrease type
// @Description Create log product with decrease type
// @Tags log_product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} dto.Response
// @Failure 401
// @Failure 400
// @Failure 500
// @Router /log-product/decrease [post]
func (c *logProductController) CreateFromDecrease(ctx *gin.Context) {

	authorizedUser := middlewares.AuthCtx(ctx.Request.Context())
	if authorizedUser == nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Not Logged In!",
		})
		return
	}

	var newLogProduct dto.CreateFromIncreaseParam
	err := ctx.ShouldBind(&newLogProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errH(err))
		return
	}

	getProductTx := services.BeginTransaction()

	product, err := services.Database.ProductGetByID(newLogProduct.ProductId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if product.Qty < newLogProduct.ProductQty {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status: false,
			Data:   nil,
			Error:  "Product Qty is not enough to be decreased!",
		})
		return
	}

	if getProductTxErr := getProductTx.Commit().Error; getProductTxErr != nil {
		ctx.AbortWithError(http.StatusInternalServerError, getProductTxErr)
		return
	}

	updateProductTx := services.BeginTransaction()

	services.Database.ProductUpdate(models.Product{
		ID:            product.ID,
		Code:          product.Code,
		ProductTypeId: product.ProductTypeId,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		Qty:           product.Qty - newLogProduct.ProductQty,
	})

	logProduct, err := services.Database.LogProductCreate(models.LogProduct{
		Status:     "keluar",
		TotalQty:   newLogProduct.ProductQty,
		TotalPrice: 1.02 * (newLogProduct.ProductQty * product.Price),
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	if updateProductTxErr := updateProductTx.Commit().Error; updateProductTxErr != nil {
		ctx.AbortWithError(http.StatusInternalServerError, updateProductTxErr)
		return
	}

	createPivotTx := services.BeginTransaction()

	pivot, err := services.Database.PivotCreate(models.Pivot{
		IDProduct:    product.ID,
		IDLogProduct: logProduct.ID,
		Qty:          newLogProduct.ProductQty,
		Price:        product.Price,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}
	fmt.Println(pivot)

	if createPivotTxErr := createPivotTx.Commit().Error; createPivotTxErr != nil {
		ctx.AbortWithError(http.StatusInternalServerError, createPivotTxErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: true,
		Data:   logProduct,
		Error:  "",
	})
}
