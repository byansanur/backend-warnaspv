package models

import (
	"../structs"
	"fmt"
	"strconv"
)

func UpdateOrderLoan(id_order string, plafond string, down_payment string, installment string, tenor string, updated_by string) structs.JsonResponse {

	var (
		updateorderloan structs.UpdateOrderLoan
	)

	response := structs.JsonResponse{}

	id_orders, _ := strconv.Atoi(id_order)
	plafonds, _ := strconv.Atoi(plafond)
	down_payments, _ := strconv.Atoi(down_payment)
	installments, _ := strconv.Atoi(installment)
	tenors, _ := strconv.Atoi(tenor)
	updated_bys, _ := strconv.Atoi(updated_by)

	updateorderloan.IdOrder = id_orders
	updateorderloan.Plafond = plafonds
	updateorderloan.DownPayment = down_payments
	updateorderloan.Tenor = tenors
	updateorderloan.Installment = installments
	updateorderloan.UpdatedBy = updated_bys

	err := idb.DB.Table("order_loan").Where("id_order = ?", id_order).Updates(updateorderloan).Error

	fmt.Println("err", err)
	if err != nil {

		response.ApiStatus = 0
		response.ApiMessage = err.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = "success"
		response.Data = updateorderloan
	}

	return response
}
