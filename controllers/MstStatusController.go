package controllers

import (
	"../models"
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
)

func GetOrderMstStatus(c *gin.Context) {

	responses := structs.JsonResponse{}
	responses = models.GetOrderMstStatus()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetOrderMstReason(c *gin.Context) {

	responses := structs.JsonResponse{}
	responses = models.GetOrderMstReason()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetCmsPrivileges(c *gin.Context) {

	responses := structs.JsonResponse{}
	responses = models.GetCmsPrivilegs()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetMstOutlet(c *gin.Context) {

	responses := structs.JsonResponse{}

	limit := c.Query("limit")
	offset := c.Query("offset")
	id_mst_branch := c.Query("id_mst_branch")

	responses = models.GetMstOutlet(id_mst_branch, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetMstOutletDetail(c *gin.Context) {

	responses := structs.JsonResponse{}

	id := c.Query("id")

	if id == "" {
		responses.ApiMessage = "Required Id"
	} else {
		responses = models.GetMstOutletDetail(id)
	}

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(400, responses)
	}

}

func GetMstBranch(c *gin.Context) {

	responses := structs.JsonResponse{}
	responses = models.GetMstBranch()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetMstBizType(c *gin.Context) {

	responses := structs.JsonResponse{}
	responses = models.GetMstBizType()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetMstLogDesc(c *gin.Context) {

	responses := structs.JsonResponse{}
	responses = models.GetMstLogDesc()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetTargetMstStatus(c *gin.Context) {

	responses := structs.JsonResponse{}
	responses = models.GetTargetMstStatus()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetVueMenusPrivileges(c *gin.Context) {

	responses := structs.JsonResponse{}

	id_cms_privileges := c.Query("id_cms_privileges")
	responses = models.GetVueMenusPrivileges(id_cms_privileges)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetMstUnits(c *gin.Context) {

	responses := structs.JsonResponse{}

	limit := c.Query("limit")
	offset := c.Query("offset")
	id_mst_branch := c.Query("id_mst_branch")
	merk := c.Query("merk")
	otr := c.Query("otr")
	types := c.Query("type")
	model := c.Query("models")
	kode_unit := c.Query("kode_unit")

	responses = models.GetMstUnits(id_mst_branch, merk, types, model, kode_unit, otr, limit, offset)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func GetMstUnit(c *gin.Context) {

	responses := structs.JsonResponse{}

	id := c.Query("id")

	responses = models.GetMstUnit(id)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func UpdateMstUnit(c *gin.Context) {

	responses := structs.JsonResponse{}

	id := c.PostForm("id")
	merk := c.PostForm("merk")
	otr := c.PostForm("otr")
	types := c.PostForm("type")
	model := c.PostForm("models")
	kode_unit := c.PostForm("kode_unit")

	responses = models.UpdateMstUnits(id, merk, types, model, kode_unit, otr)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func CreateMstUnit(c *gin.Context) {

	responses := structs.JsonResponse{}

	id_mst_branch := c.PostForm("id_mst_branch")
	id_cms_users := c.PostForm("id_cms_users")
	merk := c.PostForm("merk")
	otr := c.PostForm("otr")
	types := c.PostForm("type")
	model := c.PostForm("models")
	kode_unit := c.PostForm("kode_unit")
	year := c.PostForm("year")

	responses = models.CreateMstUnits(id_mst_branch, year, merk, types, model, kode_unit, otr, id_cms_users)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func UpdateMstOutlet(c *gin.Context) {

	responses := structs.JsonResponse{}

	id_mst_branch := c.PostForm("id_mst_branch")
	id_cms_users := c.PostForm("id_cms_users")
	id := c.PostForm("id")
	outlet_id := c.PostForm("outlet_id")
	outlet_sys_id := c.PostForm("outlet_sys_id")
	outlet_name := c.PostForm("outlet_name")
	id_mst_biz_type := c.PostForm("id_mst_biz_type")
	outlet_location := c.PostForm("outlet_location")
	outlet_Address := c.PostForm("outlet_Address")
	id_mst_address := c.PostForm("id_mst_address")
	outlet_fif_code := c.PostForm("outlet_fif_code")
	outlet_desc := c.PostForm("outlet_desc")
	outlet_status := c.PostForm("outlet_status")
	outlet_lat := c.PostForm("outlet_lat")
	outlet_lng := c.PostForm("outlet_lng")
	updated_by := c.PostForm("updated_by")

	responses = models.UpdateMstOutlet(id, outlet_id, outlet_sys_id, outlet_name, id_mst_branch,
		id_mst_biz_type, outlet_location, outlet_Address, id_mst_address,
		outlet_fif_code, outlet_desc, outlet_status,
		outlet_lat, outlet_lng, id_cms_users, updated_by)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func CreateMstOutlet(c *gin.Context) {

	responses := structs.JsonResponse{}

	id_mst_branch := c.PostForm("id_mst_branch")
	id_cms_users := c.PostForm("id_cms_users")
	outlet_id := c.PostForm("outlet_id")
	outlet_sys_id := c.PostForm("outlet_sys_id")
	outlet_name := c.PostForm("outlet_name")
	id_mst_biz_type := c.PostForm("id_mst_biz_type")
	outlet_location := c.PostForm("outlet_location")
	outlet_Address := c.PostForm("outlet_Address")
	id_mst_address := c.PostForm("id_mst_address")
	outlet_fif_code := c.PostForm("outlet_fif_code")
	outlet_desc := c.PostForm("outlet_desc")
	outlet_status := "Y"
	outlet_lat := c.PostForm("outlet_lat")
	outlet_lng := c.PostForm("outlet_lng")

	responses = models.CreateMstOutlet(outlet_id, outlet_sys_id, outlet_name, id_mst_branch,
		id_mst_biz_type, outlet_location, outlet_Address, id_mst_address,
		outlet_fif_code, outlet_desc, outlet_status,
		outlet_lat, outlet_lng, id_cms_users)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}
