package controllers

import (
	"../models"
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetOrderStatus(c *gin.Context) {

	created_at := c.Query("created_at")
	id_mst_branch := c.Query("id_mst_branch")
	id_mst_outlet := c.Query("id_mst_outlet")
	id_order_mst_status := c.Query("id_order_mst_status")
	limit := c.Query("limit")
	offset := c.Query("offset")
	first_name := c.Query("first_name")
	status := c.Query("status")

	responses := structs.JsonResponse{}

	var id_mst_branch_ary []int

	if id_mst_branch != "" {
		strs := strings.Split(id_mst_branch, ",")
		id_mst_branch_ary = make([]int, len(strs))
		for i := range id_mst_branch_ary {
			id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
		}
	}

	responses = models.GetOrderStatus(id_order_mst_status, id_mst_outlet, id_mst_branch_ary, created_at, offset, limit, first_name, status)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetOrderDetail(c *gin.Context) {

	id_order := c.Query("id_order")

	responses := structs.JsonResponse{}

	responses = models.OrderDetail(id_order)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func UpdateOrderStatus(c *gin.Context) {

	id := c.PostForm("id")
	id_order_mst_status := c.PostForm("id_order_mst_status")
	id_order_mst_reason := c.PostForm("id_order_mst_reason")
	updated_by := c.PostForm("updated_by")
	plafond := c.PostForm("plafond")
	down_payment := c.PostForm("down_payment")
	installment := c.PostForm("installment")
	tenor := c.PostForm("tenor")

	fmt.Println("id_order_mst_reason", id_order_mst_reason)

	response := structs.JsonResponse{}

	if id == "" && updated_by == "" {
		response.ApiStatus = 0
		response.ApiMessage = "field kosong"
		response.Data = nil

	} else {
		update_order := models.UpdateOrderStatus(id, id_order_mst_status, id_order_mst_reason, updated_by)

		if update_order.ApiStatus == 1 {
			response = models.UpdateOrderLoan(id, plafond, down_payment, installment, tenor, updated_by)
		}

	}

	c.JSON(http.StatusOK, response)

}

func GetOrderUsers(c *gin.Context) {

	created_at := c.Query("created_at")
	id_mst_branch := c.Query("id_mst_branch")
	id_mst_outlet := c.Query("id_mst_outlet")
	limit := c.Query("limit")
	offset := c.Query("offset")

	responses := structs.JsonResponse{}
	var id_mst_branch_ary []int

	if id_mst_branch != "" {
		strs := strings.Split(id_mst_branch, ",")
		id_mst_branch_ary = make([]int, len(strs))
		for i := range id_mst_branch_ary {
			id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
		}
	}

	responses = models.OrderUsers(id_mst_outlet, id_mst_branch_ary, created_at, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetDealsDetails(c *gin.Context) {

	created_at := c.Query("created_at")
	id_mst_branch := c.Query("id_mst_branch")
	id_cms_users := c.Query("id_cms_users")
	limit := c.Query("limit")
	offset := c.Query("offset")

	responses := structs.JsonResponse{}

	responses = models.GetDealsDetail(id_cms_users, id_mst_branch, created_at, offset, limit)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func DealsStatus(c *gin.Context) {

	created_at := c.Query("created_at")
	id_mst_branch := c.Query("id_mst_branch")

	responses := structs.JsonResponse{}

	var id_mst_branch_ary []int

	if id_mst_branch != "" {
		strs := strings.Split(id_mst_branch, ",")
		id_mst_branch_ary = make([]int, len(strs))
		for i := range id_mst_branch_ary {
			id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
		}
	}

	responses = models.DealStatus(id_mst_branch_ary, created_at)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func OrderBranch(c *gin.Context) {

	created_at := c.Query("created_at")
	id_mst_branch := c.Query("id_mst_branch")

	responses := structs.JsonResponse{}

	responses = models.GetOrderBranch(id_mst_branch, created_at)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func OrderDownload(c *gin.Context) {

	created_at := c.Query("created_at")
	//created_at2 := c.Query("created_at2")
	id_mst_branch := c.Query("id_mst_branch")
	id_mst_outlet := c.Query("id_mst_outlet")
	limit := c.Query("limit")
	offset := c.Query("offset")

	responses := structs.JsonResponse{}

	var id_mst_branch_ary []int

	if id_mst_branch != "" {
		strs := strings.Split(id_mst_branch, ",")
		id_mst_branch_ary = make([]int, len(strs))
		for i := range id_mst_branch_ary {
			id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
		}
	}

	responses = models.OrderDownload(id_mst_outlet, id_mst_branch_ary, created_at, limit, offset)

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetOrderPhoto(c *gin.Context) {

	id_order := c.Query("id_order")

	responses := structs.JsonResponse{}

	if id_order == "" {
		responses.ApiMessage = "Required id_order"
	} else {
		responses = models.GetPhotoOrder(id_order)
	}

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}
