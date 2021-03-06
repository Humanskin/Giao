package sqlite

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"sync"
)

var Db *sql.DB

type t_base_enterprise_orderBill struct {
	Id              int
	OrderNo         string
	createTime      string
	RecordType      string
	billType        string
	amount          int
	billDetail      string
	billNote        string
	billTime        string
	payType         string
	name            string
	paid            int
	payTime         string
	payNote         string
	connectPayOrder string
}

type T_base_enterprise_vehicle_insurance struct {
	id           int    `json:"id"`
	EnterpriseId int    `json:"enterpriseId"`
	RetailId     int    `json:"retailId"`
	CarId        int    `json:"carId"`
	InsuranceNo  string `json:"insuranceNo"`
	Type         string `json:"type"`
	Price        int    `json:"price"`
	Company      string `json:"company"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	Operator     string `json:"operator"`
	IsStop       int    `json:"isStop"`
}

type User struct {
	Id           int    `json:"id"`
	EnterpriseId int    `json:"enterpriseId"`
	RetailId     int    `json:"retailId"`
	Realname     string `json:"realname"`
	IsUser       error  `json:"isUser"`
}

type ValidateTest struct {
	EnterpriseId int
	RetailId     int
	CarId        int
	InsuranceId  int
	InsuranceNo  string
	IsHaveNo     error
	IsCar        error
}

// add data
func InsertRowDemo(ins *T_base_enterprise_vehicle_insurance) interface{} {
	sqlStr := fmt.Sprintf("insert into t_base_enterprise_vehicle_insurance (enterpriseId,retailId,carId,insuranceNo,type,price,company,startTime,operator) values (%d,%d,%d,'%s','%s',%d,'%s','%s','%s')",
		ins.EnterpriseId, ins.RetailId, ins.CarId, ins.InsuranceNo, ins.Type, ins.Price, ins.Company, ins.StartTime, ins.Operator)

	ret, err := Db.Exec(sqlStr)
	if err != nil {
		return err
	}

	theId, err := ret.LastInsertId()
	if err != nil {
		return err
	}

	return theId
}

// select id,enterpriseId,retailId and realname from User table
func GetUser(user *User, wg *sync.WaitGroup) {
	sqlStr := fmt.Sprintf("select enterpriseId,retailId,realname from t_base_enterprise_user where id=%d", user.Id)
	err := Db.QueryRow(sqlStr).Scan(&user.EnterpriseId, &user.RetailId, &user.Realname)
	if err != nil {
		user.IsUser = err
	}
	wg.Done()
	return
}

// select carInfo from User table
func GetCarInfo(user *ValidateTest, wg *sync.WaitGroup) {
	sqlStr := fmt.Sprintf("select id from t_base_enterprise_vehicle where id=%d and status='%s'", user.CarId, "在线")
	err := Db.QueryRow(sqlStr).Scan(&user.CarId)
	if err != nil {
		user.IsCar = err
	}
	wg.Done()
	return
}

// select id,enterpriseId,retailId and realname from User table
func GetInsNo(user *ValidateTest, wg *sync.WaitGroup) {
	sqlStr := fmt.Sprintf("select id from `t_base_enterprise_vehicle_insurance` where insuranceNo='%s'", user.InsuranceNo)
	err := Db.QueryRow(sqlStr).Scan(&user.InsuranceId)
	if err == sql.ErrNoRows {
		err = nil
	} else if err == nil{
		err = sql.ErrNoRows
	}
	wg.Done()
	user.IsHaveNo = err

	return
}

func InitDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/zhongzhong"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = Db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() error {
	err := Db.Close()
	if err != nil {
		return err
	}
	return nil
}

type Obs struct {
	Id         int
	OrderNo    string
	RecordType string
	Data       struct {
		Id         int
		OrderNo    string
		RecordType string
	}
}

func QueryRowStruct() Obs {
	sqlStr := "select id,orderNo,recordType from t_base_enterprise_orderBill where id=?"
	var orderBill t_base_enterprise_orderBill
	var o Obs
	err := Db.QueryRow(sqlStr, 1).Scan(&orderBill.Id, &orderBill.OrderNo, &orderBill.RecordType)
	if err != nil {
		fmt.Printf("error: %v", err)
		return o
	}

	fmt.Printf("id: %d\norderNo: %v\nrecordType: %v", orderBill.Id, orderBill.OrderNo, orderBill.RecordType)

	var Obs Obs
	Obs.Id = orderBill.Id
	Obs.OrderNo = orderBill.OrderNo
	Obs.RecordType = orderBill.RecordType
	Obs.Data.Id = Obs.Id
	Obs.Data.OrderNo = Obs.OrderNo
	Obs.Data.RecordType = Obs.RecordType
	return Obs
}

// select 1 row
func QueryRow() map[string]interface{} {
	sqlStr := "select id,orderNo,recordType from t_base_enterprise_orderBill where id=?"
	var orderBill t_base_enterprise_orderBill

	err := Db.QueryRow(sqlStr, 1).Scan(&orderBill.Id, &orderBill.OrderNo, &orderBill.RecordType)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	fmt.Printf("id: %d\norderNo: %v\nrecordType: %v", orderBill.Id, orderBill.OrderNo, orderBill.RecordType)
	ob := make(map[string]interface{})
	ob["id"] = orderBill.Id
	ob["orderNo"] = orderBill.OrderNo
	ob["recordType"] = orderBill.RecordType

	return ob
}

// select rows
func QueryRows() {
	sqlStr := "select id,orderNo,recordType from t_base_enterprise_orderBill where id > ?"

	rows, err := Db.Query(sqlStr, 90)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	defer rows.Close()
	//return rows
	for rows.Next() {
		var orderBill t_base_enterprise_orderBill
		err := rows.Scan(&orderBill.Id, &orderBill.OrderNo, &orderBill.RecordType)
		if err != nil {
			fmt.Printf("error : %s", err)
		}
		fmt.Printf("id:%d orderNo:%s recordType:%s\n", orderBill.Id, orderBill.OrderNo, orderBill.RecordType)
	}
}

func RunSql() interface{} {
	err := InitDB()
	if err != nil {
		return fmt.Sprintf("init db failed,err: %v\n", err)
	}

	orderBill := QueryRow()
	return orderBill
}
