package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// ProductValidator echo validator for product
type ProductValidator struct {
	validator *validator.Validate
}

// Validate validates product request body
func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{
	{
		1: "mobile",
	},
	{
		2: "test",
	},
	{
		3: "laptop",
	},
}

var e = echo.New()

func HealthCheck(c echo.Context) error {
	fmt.Println("Inside handler")
	return c.String(http.StatusOK, "Test echo framework")
}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	type body struct {
		Name            string `json:"product_name" validate:"required,min=4"`
		Vendor          string `json:"vendor" validate:"min=5,max=10"`
		Email           string `json:"email" validate:"required_with=Vendor,email"`
		Website         string `json:"website" validate:"url"`
		Country         string `json:"country" validate:"len=2"`
		DefaultDeviceIP string `json:"default_device_ip" validate:"ip"`
	}
	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	e.Validator = &ProductValidator{validator: validator.New()}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}
	products = append(products, product)
	return c.JSON(http.StatusOK, products)
}

func UpdateProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, p := range products {
		for k := range p {
			if pID == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	e.Validator = &ProductValidator{validator: validator.New()}
	if err = c.Validate(reqBody); err != nil {
		return err
	}
	product[pID] = reqBody.Name
	return c.JSON(http.StatusOK, products)
}

func DeleteProduct(c echo.Context) error {
	var product map[int]string
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for i, p := range products {
		for k := range p {
			if pID == k {
				product = p
				index = i
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	slice := func(s []map[int]string, index int) []map[int]string {
		//[1,2,3,4,5]
		// [1,2]+[4,5] = [1,2,4,5]
		return append(s[:index], s[index+1:]...)
	}
	products = slice(products, index)
	return c.JSON(http.StatusOK, products)
}
