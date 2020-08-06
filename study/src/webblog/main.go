package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	sc "webblog/middleware"
	sqls "webblog/sqlconfig"
	"webblog/sqlite"
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

type Insurance struct {
	UserId int `json:"userId"`
	Data   struct {
		CarId       int     `json:"CarId"`
		InsuranceNo string  `json:"InsuranceNo"`
		Type        string  `json:"Type"`
		Price       float64 `json:"Price"`
		Company     string  `json:"Company"`
		StartTime   string  `json:"StartTime"`
	} `json:"data"`
}

var wg sync.WaitGroup
var ch = make(chan interface{}, 3)

func main() {

	r := gin.Default()

	r.POST("/test/js", func(c *gin.Context) {
		var v Insurance
		err := c.BindJSON(&v)
		if err != nil {
			c.JSON(333, gin.H{
				"status":  "333",
				"message": err,
			})
			return
		}

		c.JSON(333, gin.H{
			"status":  "200",
			"message": v,
		})
		return
	})

	r.POST("/addInsurance", func(c *gin.Context) {

		var insurance Insurance
		//var insurance sqlite.T_base_enterprise_vehicle_insurance

		err := c.BindJSON(&insurance)
		if err != nil {
			c.JSON(333, gin.H{
				"status":  "333",
				"message": err,
			})
			return
		}

		err = sqlite.InitDB()
		if err != nil {
			c.JSON(333, gin.H{
				"status":  "333",
				"message": err,
			})
		}

		operator := sqlite.User{
			Id: insurance.UserId,
		}

		datas := sqlite.ValidateTest{
			CarId:       insurance.Data.CarId,
			InsuranceNo: insurance.Data.InsuranceNo,
		}

		//// 获取操作人信息
		wg.Add(1)
		//
		//// 检查用户是否存在
		go func() {
			sqlite.GetUser(&operator)
			wg.Done()
		}()

		// 检查车辆是否存在
		wg.Add(1)
		go func() {
			sqlite.GetCarInfo(&datas)
			wg.Done()
		}()

		// 检查保险号是否重复
		wg.Add(1)
		go func() {
			sqlite.GetInsNo(&datas)
			wg.Done()
		}()

		wg.Wait()
		//for range ch {
		//	if err := <-ch; err != nil {
		//		c.JSON(333, gin.H{
		//			"status":  "333",
		//			"message": fmt.Sprintf("未找到操作人信息1，请核实%v", err),
		//		})
		//		return
		//	}
		//}
		//if errs := <-ch;errs != nil {
		//	c.JSON(333, gin.H{
		//		"status":  "333",
		//		"message": fmt.Sprintf("未找到操作人信息1，请核实%v", errs),
		//	})
		//	return
		//}

		if operator.IsUser != nil || datas.IsHaveNo != nil || datas.IsCar != nil{
			c.JSON(333, gin.H{
				"status":  "333",
				"message": "信息有误，请核实",
			})
			return
		}

		//if err != nil {
		//	c.JSON(333, gin.H{
		//		"status":  "333",
		//		"message": fmt.Sprintf("未找到操作人信息1，请核实%v", err),
		//	})
		//}
		//c.JSON(333, gin.H{
		//	"status":  "333",
		//	"message": fmt.Sprintf("未找到操作人信息2，请核实%v %v", operator,datas),
		//})
		//return
		//for err := range ch {
		//	if err != nil {
		//		c.JSON(333, gin.H{
		//			"status":  "333",
		//			"message": fmt.Sprintf("未找到操作人信息1，请核实%v", err),
		//		})
		//		return
		//	}
		//}
		//
		//if datas.InsuranceId != 0 || datas.CarId == 0 {
		//	c.JSON(555, gin.H{
		//		"status":  "555",
		//		"message": fmt.Sprintf("未找到操作人信息3，请核实%s、%s", "InsuranceId", "CarId"),
		//	})
		//	return
		//}
		//c.JSON(444, gin.H{
		//	"status":  "444",
		//	"message": fmt.Sprintf("未找到操作人信息2，请核实%v", datas),
		//})
		//
		//return
		//close(ch)
		var insertInsurance sqlite.T_base_enterprise_vehicle_insurance

		insertInsurance.CarId = insurance.Data.CarId
		insertInsurance.InsuranceNo = insurance.Data.InsuranceNo
		insertInsurance.Type = insurance.Data.Type
		insertInsurance.Price = int(insurance.Data.Price * 100)
		insertInsurance.Company = insurance.Data.Company
		insertInsurance.StartTime = insurance.Data.StartTime
		insertInsurance.EnterpriseId = operator.EnterpriseId
		insertInsurance.RetailId = operator.RetailId
		insertInsurance.Operator = operator.Realname

		// 插入数据
		data := sqlite.InsertRowDemo(&insertInsurance)

		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"message2": data,
		})

		return
	})

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

		route.GET("/", sc.StatCost(), func(c *gin.Context) {
			//c.String(http.StatusOK, "Hello World!")
			c.JSON(http.StatusOK, gin.H{
				"name":    c.MustGet("name"),
				"message": "Hello World!",
			})
		})

		route.GET("/sql", func(c *gin.Context) {

			err := sqls.InitDB()
			if err != nil {
				log.Printf("error from %v", err)
			} else {
				c.String(200, "that is ok")
			}
			c.String(200, "that is ok")
			return
			//c.String(http.StatusOK, "Hello World!")
			//c.JSON(http.StatusOK, gin.H{
			//	"name":    c.MustGet("name"),
			//	"message": "Hello World!",
			//})
		})

	}

	r.GET("/getOrderBill", func(c *gin.Context) {
		err := sqlite.InitDB()
		if err != nil {
			c.String(403, fmt.Sprintf("init db failed,err: %v\n", err))
		}
		var orderBill = make(map[string]interface{})
		orderBill = sqlite.QueryRow()

		c.JSON(http.StatusOK, gin.H{
			"id":         orderBill["id"],
			"orderNo":    orderBill["orderNo"],
			"recordType": orderBill["recordType"],
		})
	})

	r.GET("/getOrderBillStruct", func(c *gin.Context) {
		err := sqlite.InitDB()
		if err != nil {
			c.String(403, fmt.Sprintf("init db failed,err: %v\n", err))
		}

		//type Qbs struct {
		//	Id         int
		//	OrderNo    string
		//	RecordType string
		//}
		var orderBill sqlite.Obs
		//
		//qps := sqlite.QueryRowStruct()
		//orderBill.Id = qps.Id
		//orderBill.OrderNo = qps.OrderNo
		//orderBill.RecordType = qps.RecordType
		orderBill = sqlite.QueryRowStruct()
		//c.String(333, "%T", orderBill)
		//return
		c.JSON(http.StatusOK, gin.H{
			"id":         orderBill.Id,
			"orderNo":    orderBill.OrderNo,
			"recordType": orderBill.RecordType,
			"data":       orderBill.Data,
		})
	})

	r.Run(":8899")
}
