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

func TelemarketingDownload(c *gin.Context) {

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

	responses = models.TelemarketingDownload(id_mst_outlet, id_mst_branch_ary, created_at, limit, offset)

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func TelemarketingUsers(c *gin.Context) {

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

	responses = models.TelemarketingUsers(id_mst_outlet, id_mst_branch_ary, created_at, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetDataAssignment(c *gin.Context) {

	created_at := c.Query("created_at")
	id_cms_users := c.Query("id_cms_users")
	id_mst_branch := c.Query("id_mst_branch")
	id_target_mst_status := c.Query("id_target_mst_status")
	category := c.Query("category")
	no_contract := c.Query("no_contract")
	provider_1 := c.Query("provider_1")
	provider_2 := c.Query("provider_2")
	kelurahan := c.Query("kelurahan")
	kecamatan := c.Query("kecamatan")
	kabupaten := c.Query("kabupaten")
	provinsi := c.Query("provinsi")
	limit := c.Query("limit")
	offset := c.Query("offset")
	id_mst_log_desc := c.Query("id_mst_log_desc")

	//limit := c.Query("limit")
	//offset := c.Query("offset")

	responses := structs.JsonResponse{}

	responses = models.GetDataAssignmnet(created_at, id_mst_branch, id_cms_users, id_target_mst_status, category,
		no_contract, provider_1, provider_2, kelurahan, kecamatan, kabupaten, provinsi, id_mst_log_desc, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func Assignment(c *gin.Context) {
	responses := structs.JsonResponse{}

	id_cms_users := c.PostForm("id_cms_users")
	id_target_mst_status := c.PostForm("id_target_mst_status")
	//id := c.PostFormArray("id")
	id := c.PostForm("id")
	total := c.PostForm("total")
	updated_by := c.PostForm("updated_by")
	fmt.Println("id ", id)

	if id == "" {
		responses.ApiMessage = "required ID"
	} else if id_cms_users == "" {
		responses.ApiMessage = "Required id_cms_users"
	} else if id_target_mst_status == "" {
		responses.ApiMessage = "Required id_target_mst_status"
	} else {
		strs := strings.Split(id, ",")
		ary := make([]int, len(strs))
		for i := range ary {
			ary[i], _ = strconv.Atoi(strs[i])
		}

		responses = models.UpdateAssignment(ary, id_target_mst_status,
			id_cms_users, updated_by)
		responses = models.TargetLogAssignment(updated_by, id_cms_users, total, id)
	}

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(200, responses)
	}

}

func DownloadTargetAssignmentLog(c *gin.Context) {

	created_at := c.Query("created_at")
	id_mst_branch := c.Query("id_mst_branch")
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

	responses = models.DownloadTargetAssignmentLog(id_mst_branch_ary, created_at, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetTargetUpload(c *gin.Context) {

	id_mst_branch := c.Query("id_mst_branch")
	limit := c.Query("limit")
	offset := c.Query("offset")

	responses := structs.JsonResponse{}

	responses = models.GetTargetUpload(id_mst_branch, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetTargetStatus(c *gin.Context) {

	id_mst_branch := c.Query("id_mst_branch")
	created_at := c.Query("created_at")

	responses := structs.JsonResponse{}
	var id_mst_branch_ary []int

	if id_mst_branch != "" {
		strs := strings.Split(id_mst_branch, ",")
		id_mst_branch_ary = make([]int, len(strs))
		for i := range id_mst_branch_ary {
			id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
		}
	}

	responses = models.TargetStatus(id_mst_branch_ary, created_at)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}
