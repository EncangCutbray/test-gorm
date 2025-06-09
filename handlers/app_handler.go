package handlers

import (
	"cutbray/test-gorm/request"
	"cutbray/test-gorm/utils"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

// @Summary		Helo endpoint
// @Description	Return hello From Server
// @Tags			Example
// @Accept			json
// @Produce		json
// @Success		200
// @Failure		400
// @Failure		404
// @Failure		403
// @Failure		422
// @Failure		500
// @Router			/api/v1/hello [get]
func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.SuccessResponse{
			Code:    http.StatusOK,
			Status:  "success",
			Message: "hello from server",
		})
	}
}

// @Summary		Hello from query
// @Description	Send the query and return message from server
// @Tags			Example
// @Accept			json
// @Produce		json
// @Param			title	query	string	false	"title"
// @Param			content	query	string	false	"content"
// @Success		200
// @Failure		400
// @Failure		404
// @Failure		422
// @Failure		500
// @Router			/api/v1/hello-query [get]
func HelloQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.HelloQueryRequest

		if err := c.ShouldBindQuery(&req); err != nil {
			errMap := map[string]string{"general": err.Error()}
			c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Code:    http.StatusUnprocessableEntity,
				Status:  "error",
				Message: "Bind Failed",
				Errors:  errMap,
			})
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse{
			Code:    http.StatusOK,
			Status:  "success",
			Message: "Hello Query Message",
			Data:    req,
		})
	}
}

// @Summary		Send message via form data
// @Description	Return message from server
// @Tags			Example
// @Accept			application/x-www-form-urlencoded
// @Produce		json
// @Param			title	formData	string	false	"title"
// @Param			content	formData	string	true	"content"
// @Success		200
// @Failure		400
// @Failure		404
// @Failure		422
// @Failure		500
// @Router			/api/v1/hello-form-data [post]
func HelloFormData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req request.HelloFormDataRequest

		if err := c.ShouldBind(&req); err != nil {
			errMap := map[string]string{"general": err.Error()}
			c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Code:    http.StatusUnprocessableEntity,
				Status:  "error",
				Message: "Bind failed",
				Errors:  errMap,
			})
			return
		}

		if err := validate.Struct(req); err != nil {

			errs := err.(validator.ValidationErrors)
			errMap := make(map[string]string)

			for _, e := range errs {
				fieldName := strings.ToLower(e.Field())
				errMap[fieldName] = e.Tag()
			}

			c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Code:    http.StatusUnprocessableEntity,
				Status:  "error",
				Message: "Validation failed",
				Errors:  errMap,
			})
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse{
			Code:    http.StatusOK,
			Status:  "success",
			Message: "Hello Form Data",
			Data:    req,
		})

	}
}

// @Summary		Send message via body
// @Description	Return message from server
// @Tags			Example
// @Accept			json
// @Produce		json
// @Param			request	body	request.HelloBodyRequest	true	"request body"
// @Success		200
// @Failure		400
// @Failure		403
// @Failure		404
// @Failure		422
// @Failure		500
// @Router			/api/v1/hello-body [post]
func HelloBody() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req request.HelloBodyRequest

		if err := c.ShouldBind(&req); err != nil {
			errMap := map[string]string{"general": err.Error()}
			c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Code:    http.StatusUnprocessableEntity,
				Status:  "error",
				Message: "Bind failed",
				Errors:  errMap,
			})
			return
		}

		if err := validate.Struct(req); err != nil {

			errs := err.(validator.ValidationErrors)
			errMap := make(map[string]string)

			for _, e := range errs {
				fieldName := strings.ToLower(e.Field()) // ubah ke lowercase
				errMap[fieldName] = e.Tag()             // contoh: "required", "min", "email"
			}

			c.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse{
				Code:    http.StatusUnprocessableEntity,
				Status:  "error",
				Message: "Validation failed",
				Errors:  errMap,
			})
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse{
			Code:    http.StatusOK,
			Status:  "success",
			Message: "Hello Body",
			Data:    req,
		})
	}
}
