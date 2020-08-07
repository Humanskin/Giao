package sqlite

import (
	"database/sql"
	"fmt"
)

// select id,enterpriseId,retailId and realname from User table
func GetUser2(user *User) error {
	sqlStr := fmt.Sprintf("select enterpriseId,retailId,realname from t_base_enterprise_user where id=%d", user.Id)
	err := Db.QueryRow(sqlStr).Scan(&user.EnterpriseId, &user.RetailId, &user.Realname)
	if err != nil {
		user.IsUser = err
		return err
	}
	return err
}

// select carInfo from User table
func GetCarInfo2(user *ValidateTest) error {
	sqlStr := fmt.Sprintf("select id from t_base_enterprise_vehicle where id=%d and status='%s'", user.CarId, "在线")
	err := Db.QueryRow(sqlStr).Scan(&user.CarId)
	if err != nil {
		user.IsCar = err
		return err
	}
	return err
}

// select id,enterpriseId,retailId and realname from User table
func GetInsNo2(user *ValidateTest) error {
	sqlStr := fmt.Sprintf("select id from `t_base_enterprise_vehicle_insurance` where insuranceNo='%s'", user.InsuranceNo)
	err := Db.QueryRow(sqlStr).Scan(&user.InsuranceId)
	if err == sql.ErrNoRows {
		err = nil
	} else if err == nil {
		err = sql.ErrNoRows
	}
	user.IsHaveNo = err

	return err
}
