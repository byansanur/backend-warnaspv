package controllers

import (
	"../models"
	"../structs"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func GetUsers(c *gin.Context) {

	id_mst_outlet := c.Query("id_mst_outlet")
	id_mst_branch := c.Query("id_mst_branch")
	id_cms_privileges := c.Query("id_cms_privileges")
	status := c.Query("status")
	id_spv := c.Query("id_spv")
	id_oh := c.Query("id_oh")
	limit := c.Query("limit")
	offset := c.Query("offset")
	name := c.Query("name")
	privileges_name := c.Query("privileges_name")

	responses := structs.JsonResponse{}

	var id_mst_branch_ary []int

	if id_mst_branch != "" {
		strs := strings.Split(id_mst_branch, ",")
		id_mst_branch_ary = make([]int, len(strs))
		for i := range id_mst_branch_ary {
			id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
		}
	}

	responses = models.GetUsers(id_mst_outlet, id_mst_branch_ary, id_cms_privileges, status,
		id_oh, id_spv, offset, limit, name, privileges_name)

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetUsers2(c *gin.Context) {

	id_mst_outlet := c.Query("id_mst_outlet")
	id_mst_branch := c.Query("id_mst_branch")
	id_cms_privileges := c.Query("id_cms_privileges")
	status := c.Query("status")
	id_spv := c.Query("id_spv")
	id_oh := c.Query("id_oh")
	limit := c.Query("limit")
	offset := c.Query("offset")
	name := c.Query("name")
	privileges_name := c.Query("privileges_name")

	responses := structs.JsonResponse{}

	var id_mst_branch_ary []int

	if id_mst_branch != "" {
		strs := strings.Split(id_mst_branch, ",")
		id_mst_branch_ary = make([]int, len(strs))
		for i := range id_mst_branch_ary {
			id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
		}
	}

	responses = models.GetUsers2(id_mst_outlet, id_mst_branch_ary, id_cms_privileges, status,
		id_oh, id_spv, offset, limit, name, privileges_name)

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func Login(c *gin.Context) {

	npm := c.PostForm("npm")
	password := c.PostForm("password")

	responses := structs.JsonResponse{}

	if npm == "" {
		responses.ApiMessage = "Required Npm"
	} else if password == "" {
		responses.ApiMessage = "Required Password"
	} else {
		responses = models.Login(npm, password)
	}

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(200, responses)
	}

}

func DownloadUsers(c *gin.Context) {

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
	responses = models.DownloadUsers(id_mst_branch_ary, offset, limit)

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetUsersStatus(c *gin.Context) {

	id_mst_outlet := c.Query("id_mst_outlet")
	id_mst_branch := c.Query("id_mst_branch")
	limit := c.Query("limit")
	offset := c.Query("offset")

	responses := structs.JsonResponse{}

	responses = models.GetUsersStatus(id_mst_outlet, id_mst_branch, offset, limit)

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func CreateUsers(c *gin.Context) {

	npm := c.PostForm("npm")
	name := c.PostForm("name")
	id_mst_outlet := c.PostForm("id_mst_outlet")
	email := c.PostForm("email")
	password := c.PostForm("password")
	id_cms_privileges := c.PostForm("id_cms_privileges")
	//photo := c.PostForm("photo")
	photo, header, _ := c.Request.FormFile("photo")
	status := c.PostForm("status")
	created_by := c.PostForm("created_by")
	id_cms_users_oh := c.PostForm("id_cms_users_oh")
	id_cms_users_spv := c.PostForm("id_cms_users_spv")
	id_cms_users_sub_dept := c.PostForm("id_cms_users_sub_dept")
	id_mst_branch := c.PostForm("id_mst_branch")

	responses := structs.JsonResponse{}

	if npm == "" {
		responses.ApiMessage = "Required Npm"
	} else if name == "" {
		responses.ApiMessage = "Required Name"
	} else if id_cms_privileges == "" {
		responses.ApiMessage = "Required id_cms_privileges"
	} else if status == "" {
		responses.ApiMessage = "Required status"
	} else if password == "" {
		responses.ApiMessage = "Required password"
	} else if created_by == "" {
		responses.ApiMessage = "Required created_by"
	} else {
		strs := strings.Split(id_mst_branch, ",")
		ary := make([]int, len(strs))
		for i := range ary {
			ary[i], _ = strconv.Atoi(strs[i])
		}
		fmt.Println("arry ", ary)

		responses = models.CreateUsers(npm, name, email, password, photo, header, status, id_cms_privileges, id_mst_outlet, created_by,
			id_cms_users_oh, id_cms_users_spv, id_cms_users_sub_dept, ary)
	}

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func UpdateUsers(c *gin.Context) {

	id := c.PostForm("id")
	name := c.PostForm("name")
	id_mst_outlet := c.PostForm("id_mst_outlet")
	email := c.PostForm("email")
	npm := c.PostForm("npm")
	password := c.PostForm("password")
	id_cms_privileges := c.PostForm("id_cms_privileges")
	//photo := c.PostForm("photo")
	photo, header, _ := c.Request.FormFile("photo")
	status := c.PostForm("status")
	id_cms_users_oh := c.PostForm("id_cms_users_oh")
	id_cms_users_spv := c.PostForm("id_cms_users_spv")
	id_cms_users_sub_dept := c.PostForm("id_cms_users_sub_dept")
	id_mst_branch := c.PostForm("id_mst_branch")

	responses := structs.JsonResponse{}

	if id == "" {
		responses.ApiMessage = "Required Id"
	} else {
		strs := strings.Split(id_mst_branch, ",")
		ary := make([]int, len(strs))
		for i := range ary {
			ary[i], _ = strconv.Atoi(strs[i])
		}

		responses = models.UpdateUsers(id, name, npm, email, password, photo, header, status, id_cms_privileges, id_mst_outlet,
			id_cms_users_oh, id_cms_users_spv, ary, id_cms_users_sub_dept)
	}

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetUserDetail(c *gin.Context) {

	id := c.Query("id")

	responses := structs.JsonResponse{}

	if id == "" {
		responses.ApiStatus = 0
		responses.ApiMessage = "Required Id"
		responses.Data = nil

	} else {
		responses = models.GetUserDetail(id)

	}
	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(400, responses)
	}

}

func GetKPIUsers(c *gin.Context) {

	created_at1 := c.Query("created_at1")
	created_at2 := c.Query("created_at2")
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

	responses = models.GetKPIUsers(id_mst_outlet, id_mst_branch_ary, created_at1, created_at2, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func KPIUsersDownload(c *gin.Context) {

	created_at1 := c.Query("created_at1")
	created_at2 := c.Query("created_at2")
	id_mst_branch := c.Query("id_mst_branch")
	id_mst_outlet := c.Query("id_mst_outlet")
	limit := c.Query("limit")
	offset := c.Query("offset")

	responses := structs.JsonResponse{}

	strs := strings.Split(id_mst_branch, ",")
	id_mst_branch_ary := make([]int, len(strs))
	for i := range id_mst_branch_ary {
		id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
	}

	responses = models.KPIUsersDownload(id_mst_outlet, id_mst_branch_ary, created_at1, created_at2, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func Dashboard(c *gin.Context) {

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

	responses = models.Dashboard(id_mst_branch_ary, created_at)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetCmsUsersArea(c *gin.Context) {

	id_cms_users := c.Query("id_cms_users")
	id_mst_branch := c.Query("id_mst_branch")

	responses := structs.JsonResponse{}

	responses = models.GetUsersArea(id_cms_users, id_mst_branch)

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}