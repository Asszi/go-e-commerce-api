package server

import (
	"strconv"

	"github.com/asszi/go-e-commerce-api/internal/dto"
	"github.com/asszi/go-e-commerce-api/internal/services"
	"github.com/asszi/go-e-commerce-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) createCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)
	}

	productService := services.NewProductService(s.db)
	category, err := productService.CreateCategory(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to create category", err)

		return
	}

	utils.CreatedResponse(c, "category creates successfully", category)
}

func (s *Server) getCategories(c *gin.Context) {
	productService := services.NewProductService(s.db)
	categories, err := productService.GetCategories()
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to fetch categories", err)

		return
	}

	utils.SuccessResponse(c, "categories retrieved successfully", categories)
}

func (s *Server) updateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid category id", err)

		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)

		return
	}

	productService := services.NewProductService(s.db)
	category, err := productService.UpdateCategory(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to update category", err)

		return
	}

	utils.SuccessResponse(c, "category updated successfully", category)
}

func (s *Server) deleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid category id", err)

		return
	}

	productService := services.NewProductService(s.db)
	if err := productService.DeleteCategory(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete category", err)

		return
	}

	utils.SuccessResponse(c, "category deleted successfully", nil)
}

func (s *Server) createProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)

		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.CreateProduct(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to create product", err)

		return
	}

	utils.CreatedResponse(c, "product created successfully", product)
}

func (s *Server) getProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	productService := services.NewProductService(s.db)
	product, meta, err := productService.GetProducts(page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to fetch product", err)

		return
	}

	utils.PaginatedSuccessResponse(c, "products retrieved successfully", product, *meta)
}

func (s *Server) getProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid product id", err)

		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.GetProduct(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "product not found")

		return
	}

	utils.SuccessResponse(c, "product retrieved successfully", product)
}

func (s *Server) updateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid product id", err)

		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)

		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.UpdateProduct(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to update product", err)

		return
	}

	utils.SuccessResponse(c, "product updated successfully", product)
}

func (s *Server) deleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid product id", err)
	}

	productService := services.NewProductService(s.db)
	if err := productService.DeleteProduct(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete product", err)

		return
	}

	utils.SuccessResponse(c, "product deleted successfully", nil)
}
