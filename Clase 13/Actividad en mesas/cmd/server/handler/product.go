package handler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	s product.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s product.Service) *productHandler {
	return &productHandler{
		s: s,
	}
}

// GetAll godoc
// @Summary      Gets all the products
// @Description  Gets all the products from the repository
// @Tags         products
// @Produce      json
// @Success      200 {object}  web.response
// @Router       /products [get]
func (h *productHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, _ := h.s.GetAll()
		web.Success(c, 200, products)
	}
}

// GetByID godoc
// @Summary      Gets a product by id
// @Description  Gets a product by id from the repository
// @Tags         products
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /products/{id} [get]
func (h *productHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("product not found"))
			return
		}
		web.Success(c, 200, product)
	}
}

// Search godoc
// @Summary      Search for a product by price
// @Description  Search for a product by price from the repository
// @Tags         products
// @Produce      json
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /products/search [get]
func (h *productHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		priceParam := c.Query("priceGt")
		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid price"))
			return
		}
		products, err := h.s.SearchPriceGt(price)
		if err != nil {
			web.Failure(c, 404, errors.New("product not found"))
			return
		}
		web.Success(c, 200, products)
	}
}

// ConsumerPrice godoc
// @Summary      Returns the prices of one or more products
// @Description  Returns the prices of one or more products from the repository
// @Tags         products
// @Produce      json
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /products/consumer_price [get]
func (h *productHandler) ConsumerPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		list, err := getList(c.Query("list"))
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		products, price, err := h.s.ConsumerPrice(list)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, gin.H{"total_price": price, "products": products})
	}
}

// getList obtiene una lista de productos desde un string
func getList(listParam string) ([]int, error) {
	if listParam == "" {
		return nil, errors.New("list can't be empty")
	}
	listParam = strings.Trim(listParam, "[]")
	list := strings.Split(listParam, ",")
	if len(list) == 0 {
		return nil, errors.New("list can't be empty")
	}
	listInt := []int{}
	for _, value := range list {
		number, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("invalid list, must be numbers")
		}
		listInt = append(listInt, number)
	}
	return listInt, nil
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(product *domain.Product) (bool, error) {
	switch {
	case product.Name == "" || product.CodeValue == "" || product.Expiration == "":
		return false, errors.New("fields can't be empty")
	case product.Quantity <= 0 || product.Price <= 0:
		if product.Quantity <= 0 {
			return false, errors.New("quantity must be greater than 0")
		}
		if product.Price <= 0 {
			return false, errors.New("price must be greater than 0")
		}
	}
	return true, nil
}

// validateExpiration valida que la fecha de expiracion sea valida
func validateExpiration(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

// Post godoc
// @Summary      Create a new product
// @Description  Create a new product in repository
// @Tags         products
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Product true "Product"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /products [post]
func (h *productHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product domain.Product
		err := c.ShouldBindJSON(&product)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&product)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateExpiration(product.Expiration)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(product)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete godoc
// @Summary      Deletes a product
// @Description  Deletes a product from the repository
// @Tags         products
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Success      204 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /products/{id} [delete]
func (h *productHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put godoc
// @Summary      Updates a product
// @Description  Updates a product from the repository
// @Tags         products
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Product true "Product"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /products/{id} [put]
func (h *productHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		var product domain.Product
		err = c.ShouldBindJSON(&product)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&product)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateExpiration(product.Expiration)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, product)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, p)
	}
}



// Patch godoc
// @Summary      Updates selected fields
// @Description  Updates selected fields from a product from the repository
// @Tags         products
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Product true "Product"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /products/{id} [patch]
func (h *productHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name        string  `json:"name,omitempty"`
		Quantity    int     `json:"quantity,omitempty"`
		CodeValue   string  `json:"code_value,omitempty"`
		IsPublished bool    `json:"is_published,omitempty"`
		Expiration  string  `json:"expiration,omitempty"`
		Price       float64 `json:"price,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Product{
			Name:        r.Name,
			Quantity:    r.Quantity,
			CodeValue:   r.CodeValue,
			IsPublished: r.IsPublished,
			Expiration:  r.Expiration,
			Price:       r.Price,
		}
		if update.Expiration != "" {
			valid, err := validateExpiration(update.Expiration)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, p)
	}
}
