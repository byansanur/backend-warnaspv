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

func Perfomance(c *gin.Context) {

	created_at := c.Query("tahun")
	id_mst_branch := c.Query("id_mst_branch")

	var id_mst_branch_ary []int

	if id_mst_branch != "" {
		strs := strings.Split(id_mst_branch, ",")
		id_mst_branch_ary = make([]int, len(strs))
		for i := range id_mst_branch_ary {
			id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
		}
	}

	//created_at01 := created_at+"-01"

	responses := structs.JsonResponse{}

	jan := models.GetOrderChart(id_mst_branch_ary, created_at+"-02")
	feb := models.GetOrderChart(id_mst_branch_ary, created_at+"-02")
	mar := models.GetOrderChart(id_mst_branch_ary, created_at+"-03")
	apr := models.GetOrderChart(id_mst_branch_ary, created_at+"-04")
	mei := models.GetOrderChart(id_mst_branch_ary, created_at+"-05")
	jun := models.GetOrderChart(id_mst_branch_ary, created_at+"-06")
	jul := models.GetOrderChart(id_mst_branch_ary, created_at+"-07")
	ags := models.GetOrderChart(id_mst_branch_ary, created_at+"-08")
	sep := models.GetOrderChart(id_mst_branch_ary, created_at+"-09")
	okt := models.GetOrderChart(id_mst_branch_ary, created_at+"-10")
	nov := models.GetOrderChart(id_mst_branch_ary, created_at+"-11")
	des := models.GetOrderChart(id_mst_branch_ary, created_at+"-12")

	leadjan := models.GetLeadChart(id_mst_branch_ary, created_at+"-01")
	leadfeb := models.GetLeadChart(id_mst_branch_ary, created_at+"-02")
	leadmar := models.GetLeadChart(id_mst_branch_ary, created_at+"-03")
	leadapr := models.GetLeadChart(id_mst_branch_ary, created_at+"-04")
	leadmei := models.GetLeadChart(id_mst_branch_ary, created_at+"-05")
	leadjun := models.GetLeadChart(id_mst_branch_ary, created_at+"-06")
	leadjul := models.GetLeadChart(id_mst_branch_ary, created_at+"-07")
	leadags := models.GetLeadChart(id_mst_branch_ary, created_at+"-08")
	leadsep := models.GetLeadChart(id_mst_branch_ary, created_at+"-09")
	leadokt := models.GetLeadChart(id_mst_branch_ary, created_at+"-10")
	leadnov := models.GetLeadChart(id_mst_branch_ary, created_at+"-11")
	leaddes := models.GetLeadChart(id_mst_branch_ary, created_at+"-12")

	getfieldjan := jan.Data.(structs.GetOrderChart)
	getfieldfeb := feb.Data.(structs.GetOrderChart)
	getfieldmar := mar.Data.(structs.GetOrderChart)
	getfieldapr := apr.Data.(structs.GetOrderChart)
	getfieldmei := mei.Data.(structs.GetOrderChart)
	getfieldjun := jun.Data.(structs.GetOrderChart)
	getfieldjul := jul.Data.(structs.GetOrderChart)
	getfieldags := ags.Data.(structs.GetOrderChart)
	getfieldsep := sep.Data.(structs.GetOrderChart)
	getfieldokt := okt.Data.(structs.GetOrderChart)
	getfieldnov := nov.Data.(structs.GetOrderChart)
	getfielddesc := des.Data.(structs.GetOrderChart)

	fmt.Println("ordersep ", getfieldsep)
	fmt.Println("sep ", sep)

	getfieldleadjan := leadjan.Data.(structs.GetLeadChart)
	getfieldleadfeb := leadfeb.Data.(structs.GetLeadChart)
	getfieldleadmar := leadmar.Data.(structs.GetLeadChart)
	getfieldleadapr := leadapr.Data.(structs.GetLeadChart)
	getfieldleadmei := leadmei.Data.(structs.GetLeadChart)
	getfieldleadjun := leadjun.Data.(structs.GetLeadChart)
	getfieldleadjul := leadjul.Data.(structs.GetLeadChart)
	getfieldleadags := leadags.Data.(structs.GetLeadChart)
	getfieldleadsep := leadsep.Data.(structs.GetLeadChart)
	getfieldleadokt := leadokt.Data.(structs.GetLeadChart)
	getfieldleadnov := leadnov.Data.(structs.GetLeadChart)
	getfieldleaddes := leaddes.Data.(structs.GetLeadChart)

	convStruct := [12]structs.Chart{}

	convStruct[0].Month = "Jan"
	convStruct[0].Booking = getfieldjan.Booking
	convStruct[0].Lead = getfieldleadjan.Leads
	convStruct[1].Month = "Feb"
	convStruct[1].Booking = getfieldfeb.Booking
	convStruct[1].Lead = getfieldleadfeb.Leads
	convStruct[2].Month = "Mar"
	convStruct[2].Booking = getfieldmar.Booking
	convStruct[2].Lead = getfieldleadmar.Leads
	convStruct[3].Month = "Apr"
	convStruct[3].Booking = getfieldapr.Booking
	convStruct[3].Lead = getfieldleadapr.Leads
	convStruct[4].Month = "Mei"
	convStruct[4].Booking = getfieldmei.Booking
	convStruct[4].Lead = getfieldleadmei.Leads

	convStruct[5].Month = "Jun"
	convStruct[5].Booking = getfieldjun.Booking
	convStruct[5].Lead = getfieldleadjun.Leads

	convStruct[6].Month = "Jul"
	convStruct[6].Booking = getfieldjul.Booking
	convStruct[6].Lead = getfieldleadjul.Leads

	convStruct[7].Month = "Ags"
	convStruct[7].Booking = getfieldags.Booking
	convStruct[7].Lead = getfieldleadags.Leads

	convStruct[8].Month = "Sep"
	convStruct[8].Booking = getfieldsep.Booking
	convStruct[8].Lead = getfieldleadsep.Leads

	convStruct[9].Month = "Okt"
	convStruct[9].Booking = getfieldokt.Booking
	convStruct[9].Lead = getfieldleadokt.Leads

	convStruct[10].Month = "Nov"
	convStruct[10].Booking = getfieldnov.Booking
	convStruct[10].Lead = getfieldleadnov.Leads

	convStruct[11].Month = "Des"
	convStruct[11].Booking = getfielddesc.Booking
	convStruct[11].Lead = getfieldleaddes.Leads

	fmt.Println("creatd_at ", created_at)

	responses.ApiStatus = 1
	responses.ApiMessage = "succ"
	responses.Data = convStruct

	//c.JSON(http.StatusOK, convStruct)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func VisumStatus(c *gin.Context) {

	created_at := c.Query("created_at")
	id_mst_branch := c.Query("id_mst_branch")
	//id_mst_visum_status := c.Query("id_mst_visum_status")

	//created_at01 := created_at+"-01"

	responses := structs.JsonResponse{}

	var id_mst_branch_ary []int

	if id_mst_branch != "" {
		strs := strings.Split(id_mst_branch, ",")
		id_mst_branch_ary = make([]int, len(strs))
		for i := range id_mst_branch_ary {
			id_mst_branch_ary[i], _ = strconv.Atoi(strs[i])
		}
	}

	jan := models.TargetVisumChart("1", id_mst_branch_ary, created_at)
	feb := models.TargetVisumChart("2", id_mst_branch_ary, created_at)
	mar := models.TargetVisumChart("3", id_mst_branch_ary, created_at)
	apr := models.TargetVisumChart("4", id_mst_branch_ary, created_at)
	mei := models.TargetVisumChart("5", id_mst_branch_ary, created_at)
	jun := models.TargetVisumChart("6", id_mst_branch_ary, created_at)
	jul := models.TargetVisumChart("7", id_mst_branch_ary, created_at)
	ags := models.TargetVisumChart("8", id_mst_branch_ary, created_at)
	sep := models.TargetVisumChart("9", id_mst_branch_ary, created_at)
	//okt :=  models.TargetVisumChart(id_mst_visum_status,id_mst_branch_ary,created_at)
	//nov :=  models.TargetVisumChart(id_mst_visum_status,id_mst_branch_ary,created_at)
	//des :=  models.TargetVisumChart(id_mst_visum_status,id_mst_branch_ary,created_at)

	leadjan := models.LeadVisumChart("1", id_mst_branch_ary, created_at)
	leadfeb := models.LeadVisumChart("2", id_mst_branch_ary, created_at)
	leadmar := models.LeadVisumChart("3", id_mst_branch_ary, created_at)
	leadapr := models.LeadVisumChart("4", id_mst_branch_ary, created_at)
	leadmei := models.LeadVisumChart("5", id_mst_branch_ary, created_at)
	leadjun := models.LeadVisumChart("6", id_mst_branch_ary, created_at)
	leadjul := models.LeadVisumChart("7", id_mst_branch_ary, created_at)
	leadags := models.LeadVisumChart("8", id_mst_branch_ary, created_at)
	leadsep := models.LeadVisumChart("9", id_mst_branch_ary, created_at)
	//leadokt := models.LeadVisumChart(id_mst_visum_status,id_mst_branch, created_at)
	//leadnov := models.LeadVisumChart(id_mst_visum_status,id_mst_branch, created_at)
	//leaddes := models.LeadVisumChart(id_mst_visum_status,id_mst_branch, created_at)

	getfieldjan := jan.Data.(structs.TargetVisumChart)
	getfieldfeb := feb.Data.(structs.TargetVisumChart)
	getfieldmar := mar.Data.(structs.TargetVisumChart)
	getfieldapr := apr.Data.(structs.TargetVisumChart)
	getfieldmei := mei.Data.(structs.TargetVisumChart)
	getfieldjun := jun.Data.(structs.TargetVisumChart)
	getfieldjul := jul.Data.(structs.TargetVisumChart)
	getfieldags := ags.Data.(structs.TargetVisumChart)
	getfieldsep := sep.Data.(structs.TargetVisumChart)
	//getfieldokt := okt.Data.(structs.TargetVisumChart)
	//getfieldnov := nov.Data.(structs.TargetVisumChart)
	//getfielddesc := des.Data.(structs.TargetVisumChart)

	fmt.Println("ordersep ", getfieldsep)
	fmt.Println("sep ", sep)

	getfieldleadjan := leadjan.Data.(structs.LeadVisumChart)
	getfieldleadfeb := leadfeb.Data.(structs.LeadVisumChart)
	getfieldleadmar := leadmar.Data.(structs.LeadVisumChart)
	getfieldleadapr := leadapr.Data.(structs.LeadVisumChart)
	getfieldleadmei := leadmei.Data.(structs.LeadVisumChart)
	getfieldleadjun := leadjun.Data.(structs.LeadVisumChart)
	getfieldleadjul := leadjul.Data.(structs.LeadVisumChart)
	getfieldleadags := leadags.Data.(structs.LeadVisumChart)
	getfieldleadsep := leadsep.Data.(structs.LeadVisumChart)
	//getfieldleadokt := leadokt.Data.(structs.LeadVisumChart)
	//getfieldleadnov := leadnov.Data.(structs.LeadVisumChart)
	//getfieldleaddes := leaddes.Data.(structs.LeadVisumChart)

	convStruct := [9]structs.ChartVisum{}

	convStruct[0].Status = "RumahKosong"
	convStruct[0].Lead = getfieldleadjan.LeadVisum
	convStruct[0].Target = getfieldjan.TargetVisum
	convStruct[1].Status = "AlamatTidakSesuai"
	convStruct[1].Lead = getfieldleadfeb.LeadVisum
	convStruct[1].Target = getfieldfeb.TargetVisum

	convStruct[2].Status = "Ditolak/Diusir"
	convStruct[2].Lead = getfieldleadmar.LeadVisum
	convStruct[2].Target = getfieldmar.TargetVisum

	convStruct[3].Status = "ButuhWaktu"
	convStruct[3].Lead = getfieldleadapr.LeadVisum
	convStruct[3].Target = getfieldapr.TargetVisum

	convStruct[4].Status = "TanyaPasangan"
	convStruct[4].Lead = getfieldleadmei.LeadVisum
	convStruct[4].Target = getfieldmei.TargetVisum

	convStruct[5].Status = "Pikirpikir"
	convStruct[5].Lead = getfieldleadjun.LeadVisum
	convStruct[5].Target = getfieldjun.TargetVisum

	convStruct[6].Status = "Berminat"
	convStruct[6].Lead = getfieldleadjul.LeadVisum
	convStruct[6].Target = getfieldjul.TargetVisum

	convStruct[7].Status = "RumahTidakKetemu"
	convStruct[7].Lead = getfieldleadags.LeadVisum
	convStruct[7].Target = getfieldags.TargetVisum

	convStruct[8].Status = "TidakBerminat"
	convStruct[8].Lead = getfieldleadsep.LeadVisum
	convStruct[8].Target = getfieldsep.TargetVisum

	//convStruct[9].Month = "Okt"
	//convStruct[9].Lead = getfieldleadokt.LeadVisum
	//convStruct[9].Target = getfieldokt.TargetVisum
	//
	//convStruct[10].Month = "Nov"
	//convStruct[10].Lead = getfieldleadnov.LeadVisum
	//convStruct[10].Target = getfieldnov.TargetVisum
	//
	//convStruct[11].Month = "Des"
	//convStruct[11].Lead = getfieldleaddes.LeadVisum
	//convStruct[11].Target = getfielddesc.TargetVisum

	fmt.Println("creatd_at ", created_at)

	responses.ApiStatus = 1
	responses.ApiMessage = "succ"
	responses.Data = convStruct

	//c.JSON(http.StatusOK, convStruct)

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}
