package controller

import (
	"fmt"
	"go_shop/config"
	"go_shop/helper"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Id        int64  `gorm:"primaryKey"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	Pattern   string `json:"pattern"`
	Figures   int    `json:"figures"`
	Active    bool   `json:"active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetProductList(c *fiber.Ctx) error {
	var products []Product
	db := config.DBConn

	sql := "SELECT * FROM products WHERE active = true"

	if gender := c.Query("gender"); gender != "" {
		sql = fmt.Sprintf("%s AND gender = '%s'", sql, gender)
	}
	if size := c.Query("size"); size != "" {
		sql = fmt.Sprintf("%s AND size = '%s'", sql, size)
	}
	if color := c.Query("color"); color != "" {
		sql = fmt.Sprintf("%s AND color = '%s'", sql, color)
	}
	if pattern := c.Query("pattern"); pattern != "" {
		sql = fmt.Sprintf("%s AND pattern = '%s'", sql, pattern)
	}
	if figures := c.Query("figures"); figures != "" {
		sql = fmt.Sprintf("%s AND figures = '%s'", sql, figures)
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	var total int64

	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, limit, (page-1)*limit)

	db.Raw(sql).Find(&products).Count(&total)
	res := helper.ListResponse(true, "OK", page, limit, total, products)
	return c.JSON(res)
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.DBConn
	var products []Product
	db.Where("id = ?", id).First(&products)
	res := helper.BuildResponse(true, "OK", products)
	return c.JSON(res)
}

func AddProduct(c *fiber.Ctx) error {
	db := config.DBConn
	products := new(Product)

	err := c.BodyParser(products)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", "Check your input", err)
		return c.Status(500).JSON(res)
	}
	err = db.Create(&products).Error
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", "Could not create note", err)
		return c.Status(500).JSON(res)
	}
	res := helper.BuildResponse(true, "OK", products)
	return c.JSON(res)

}
