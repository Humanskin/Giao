package testpost

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Invoice struct {
	UserId int `json:"userId"`
	Data   struct {
		InvoiceNo   string `json:"invoiceNo"`
		InvoiceType string `json:"invoiceType"`
		SendUser    string `json:"sendUser"`
		GetUser     string `json:"getUser"`
		CustomerID  int `json:"customerId"`
		Price       float64 `json:"price"`
		CreateDate  string `json:"createDate"`
	} `json:"data"`
}

var ins Invoice

func Tests(c *gin.Context) {
	if err := c.BindJSON(&ins); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"userId": ins.UserId,
			"data":ins.Data,
		})
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"status":401,
			"message":err,
		})
	}
	return
}
