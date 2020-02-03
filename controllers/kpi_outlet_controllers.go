package controllers

import (
	"../models"
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetKpiOutletLeads(c *gin.Context) {
	lead := structs.KpiOutletLead{}
	t := structs.Component{}
	response := structs.JsonResponse{}
	limit := c.Query("limit")
	offset := c.Query("offset")

	err := c.BindQuery(&lead)
	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}
		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {
		data, errc := models.GetKpiOutlet(lead, limit, offset)
		response.Data = data

		if errc != nil {
			response.ApiMessage = t.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageSucc()
			c.JSON(http.StatusOK, response)
		}
	}
}
