package controller

import (
	"fmt"
	"go_shop/config"
	"go_shop/helper"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	Id          int64     `gorm:"primaryKey"`
	ProductId   int       `json:"product_id"`
	DeliveryAt  string    `json:"delivery_at"`
	Quantity    int       `json:"quantity"`
	OrderStatus string    `json:"order_status"`
	PaidStatus  string    `json:"paid_status"`
	PaidDate    time.Time `gorm:"autoCreateTime"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetOrderList(c *fiber.Ctx) error {
	var orders []Order
	db := config.DBConn

	sql := "SELECT * FROM orders WHERE active = true"

	if order_status := c.Query("orderStatus"); order_status != "" {
		sql = fmt.Sprintf("%s AND order_status = '%s'", sql, order_status)
	}
	if paid_start := c.Query("paidStart"); paid_start != "" {
		sql = fmt.Sprintf("%s AND DATE('%s%%') <= DATE(paid_date)", sql, paid_start)
	}
	if paid_end := c.Query("paidEnd"); paid_end != "" {
		sql = fmt.Sprintf("%s AND DATE('%s%%') >= DATE(paid_date)", sql, paid_end)
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	var total int64

	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, limit, (page-1)*limit)

	db.Raw(sql).Find(&orders).Count(&total)
	res := helper.ListResponse(true, "OK", page, limit, total, orders)
	return c.JSON(res)
}

func AddOrder(c *fiber.Ctx) error {
	db := config.DBConn
	orders := new(Order)

	err := c.BodyParser(orders)
	if err != nil {
		res := helper.BuildResponse(false, "Bad request", "Check your input")
		return c.Status(500).JSON(res)
	}
	err = db.Create(&orders).Error
	if err != nil {
		res := helper.BuildResponse(false, "Bad request", "Could not create note")
		return c.Status(500).JSON(res)
	}
	res := helper.BuildResponse(true, "OK", orders)
	return c.JSON(res)

}
