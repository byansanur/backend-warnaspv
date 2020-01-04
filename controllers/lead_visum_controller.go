package controllers

import (
	"../models"
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func LeadVisumDownload(c *gin.Context) {

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

	responses = models.LeadVisumDownload(id_mst_outlet, id_mst_branch_ary, created_at, limit, offset)

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func LeadVisumUsers(c *gin.Context) {

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

	responses = models.LeadVisumUsers(id_mst_outlet, id_mst_branch_ary, created_at, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func LeadVisumDetail(c *gin.Context) {

	id_cms_users := c.Query("id_cms_users")
	created_at := c.Query("created_at")

	responses := structs.JsonResponse{}

	responses = models.LeadVisumDetail(id_cms_users, created_at)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetLeadVisumPhoto(c *gin.Context) {

	id_lead_visum := c.Query("id_lead_visum")

	responses := structs.JsonResponse{}

	responses = models.GetPhotoVisumLead(id_lead_visum)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}
