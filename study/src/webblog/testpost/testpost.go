package testpost

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Invoice struct {
	UserId string `json:"userId"`
	Data   struct {
		InvoiceNo   string  `json:"invoiceNo"`
		InvoiceType string  `json:"invoiceType"`
		SendUser    string  `json:"sendUser"`
		GetUser     string  `json:"getUser"`
		CustomerID  int     `json:"customerId"`
		Price       float64 `json:"price"`
		CreateDate  string  `json:"createDate"`
	} `json:"data"`
}

var ins Invoice

func Tests(c *gin.Context) {
	if err := c.BindJSON(&ins); err == nil {
		userId, _ := strconv.Atoi(ins.UserId)
		c.JSON(http.StatusOK, gin.H{
			//"userId": strconv.Atoi(ins.UserId),
			"userId": userId,
			"data":   ins.Data,
		})
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"status":  401,
			"message": err,
		})
	}
	return
}
