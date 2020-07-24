package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	tp "webblog/testpost"
)

type User struct {
	Username string
	Password string
	Id       int64
}

type Invoice struct {
	UserId string `json:"userId"`
	Data   struct {
		InvoiceNo   string `json:"invoiceNo"`
		InvoiceType string `json:"invoiceType"`
		SendUser    string `json:"sendUser"`
		GetUser     string `json:"getUser"`
		CustomerID  string `json:"customerId"`
		Price       string `json:"price"`
		CreateDate  string `json:"createDate"`
	} `json:"data"`
}

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "信息，系统管理等模块")
	})

	r.GET("/call/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "name: %s", name)
	})

	r.GET("/users", func(c *gin.Context) {
		user := c.Query("name")
		role := c.DefaultQuery("role", "everything")
		c.String(http.StatusOK, "%s is %s", user, role)
	})

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123456")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Any("/any", func(c *gin.Context) {
		id := c.Query("id")
		username := c.PostForm("username")
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"username": username,
		})
	})
	var user User

	r.Any("/jsons", func(c *gin.Context) {
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"password": user.Password,
				"id":       user.Id,
			})
		} else {
			c.String(404, fmt.Sprintf("error from should bind : %s", err))
		}
		return
	})

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": err,
			})
			return
		}

		log.Println(file.Filename)
		c.SaveUploadedFile(file, "D:/stt/go/stt_test/src/webblog/files/"+file.Filename)
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": file.Filename + " is uploaded",
		})
	})

	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["f2"]

		for index, file := range files {
			log.Println(index, file.Filename)
			dst := fmt.Sprintf("D:/stt/go/stt_test/src/webblog/files/%d-%s", index, file.Filename)
			c.SaveUploadedFile(file, dst)

		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Any("/tests", func(c *gin.Context) {
		c.Request.URL.Path = "/jsons"
		r.HandleContext(c)
	})

	var InvoiceEr Invoice
	//var JsonData interface{}
	//var Invoices Invoice

	route := r.Group("/test")
	{
		route.POST("/add", func(c *gin.Context) {
			if err := c.BindJSON(&InvoiceEr); err != nil {
				//c.JSON(http.StatusOK, gin.H{
				//	"userId":      InvoiceEr.UserID,
				//	"invoiceNo":   InvoiceEr.InvoiceNo,
				//	"invoiceType": InvoiceEr.InvoiceType,
				//	"sendUser":    InvoiceEr.SendUser,
				//	"getUser":     InvoiceEr.GetUser,
				//	"customerId":  InvoiceEr.CustomerID,
				//	"price":       InvoiceEr.Price,
				//	"createDate":  InvoiceEr.CreateDate,
				//})
				c.JSON(http.StatusOK, gin.H{
					"userId": InvoiceEr.UserId,
					"data":   InvoiceEr.Data,
				})

			} else {
				c.JSON(403, gin.H{
					"status": fmt.Sprintf("error: %s", err),
				})
			}
			return
		})

		route.POST("/adds", func(c *gin.Context) {
			tp.Tests(c)
		})
	}

	r.Run(":8899")
}
