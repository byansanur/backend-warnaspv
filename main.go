package main

import (
	"./auth"
	"./controllers"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
)

func main() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("sedoooot %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.Use(cors.AllowAll())

	public := router.Group("/api_bian/v1")
	{
		//public.POST("/login", controllers.Login)
		public.POST("/auth/login", controllers.Login)
		public.Static("/files", "./files")
		public.GET("/vue_menus_privileges", controllers.GetVueMenusPrivileges)

	}
	v1 := router.Group("/api_bian/v1")
	v1.Use(auth.Auth)
	{
		v1.POST("/bian_activity", controllers.CreateActivity)
		v1.GET("/bian_activity_detail", controllers.GetDetail)
		v1.GET("/bian_activity_status", controllers.GetActivityStatus)
		v1.GET("/bian_activity_type", controllers.GetActivityType)
		v1.GET("/list_activity", controllers.GetActivitys)
		//v1.GET("/bian_activity_list_2", controllers.GetActivity)
		//v1.GET("/bian_activity_list", controllers.GetActivityList2)
		v1.GET("/bian_kpi_outled_lead", controllers.GetKpiOutletLeads)

		v1.GET("/lead/download", controllers.DownloadLead)
		v1.GET("/lead/users", controllers.LeadUsers)
		v1.GET("/lead/detail", controllers.LeadDetail)
		v1.GET("/lead/status", controllers.LeadStatus)

		v1.GET("/lead/visum/users", controllers.LeadVisumUsers)
		v1.GET("/lead/visum/photo", controllers.GetLeadVisumPhoto)
		v1.GET("/lead/visum/detail", controllers.LeadVisumDetail)
		v1.GET("/lead/visum/download", controllers.LeadVisumDownload)

		v1.GET("/target/visum/users", controllers.TargetVisumUsers)
		v1.GET("/target/visum/photo", controllers.GetTargetVisumPhoto)
		v1.GET("/target/visum/detail", controllers.TargetVisumDetail)
		v1.GET("/target/visum/download", controllers.TargetVisumDownload)
		v1.GET("/target/target_mst_status", controllers.GetTargetMstStatus)
		v1.GET("/target/tele/users", controllers.TelemarketingUsers)
		v1.GET("/target/tele/download", controllers.TelemarketingDownload)
		v1.GET("/target/assignment/download", controllers.DownloadTargetAssignmentLog)
		v1.GET("/target/target_upload", controllers.GetTargetUpload)
		v1.GET("/target/status", controllers.GetTargetStatus)

		v1.GET("/target/assignment", controllers.GetDataAssignment)
		v1.POST("/target/assignment", controllers.Assignment)

		v1.GET("/privileges", controllers.GetCmsPrivileges)

		v1.GET("/mst_branch", controllers.GetMstBranch)
		v1.GET("/mst_log_desc", controllers.GetMstLogDesc)

		v1.GET("/users", controllers.GetUsers)
		v1.GET("/users2", controllers.GetUsers2)
		v1.GET("/users/download", controllers.DownloadUsers)
		v1.GET("/user", controllers.GetUserDetail)
		v1.GET("/users/status", controllers.GetUsersStatus)
		v1.POST("/users", controllers.CreateUsers)
		v1.PUT("/user", controllers.UpdateUsers)
		v1.GET("/users/kpi", controllers.GetKPIUsers)
		v1.GET("/users/kpi/download", controllers.KPIUsersDownload)

		v1.GET("/dashboard", controllers.Dashboard)
		v1.GET("/dashboard/perfomance", controllers.Perfomance)
		v1.GET("/dashboard/visumstatus", controllers.VisumStatus)

		v1.GET("/riwayat", controllers.Riwayat)
		v1.GET("/riwayat/download", controllers.DownloadRiwayat)

		v1.GET("/order/users", controllers.GetOrderUsers)
		v1.GET("/order/branch", controllers.OrderBranch)
		v1.GET("/order/deal/detail", controllers.GetDealsDetails)
		v1.GET("/order/status", controllers.GetOrderStatus)
		v1.GET("/order/photo", controllers.GetOrderPhoto)
		v1.GET("/order/detail", controllers.GetOrderDetail)
		v1.GET("/order/order_mst_status", controllers.GetOrderMstStatus)
		v1.GET("/order/order_mst_reason", controllers.GetOrderMstReason)
		v1.PUT("/order/update_status", controllers.UpdateOrderStatus)
		v1.GET("/order/download", controllers.OrderDownload)
		v1.GET("/order/deals/status", controllers.DealsStatus)

		v1.GET("/mst_units", controllers.GetMstUnits)
		v1.GET("/mst_unit", controllers.GetMstUnit)
		v1.POST("/mst_unit", controllers.CreateMstUnit)
		v1.PUT("/mst_unit", controllers.UpdateMstUnit)

		v1.GET("/mst_outlet", controllers.GetMstOutlet)
		v1.GET("/mst_outlet_detail", controllers.GetMstOutletDetail)
		v1.PUT("/mst_outlet", controllers.UpdateMstOutlet)
		v1.POST("/mst_outlet", controllers.CreateMstOutlet)

		v1.GET("/mst_biz_type", controllers.GetMstBizType)

		v1.GET("/cms_users_area", controllers.GetCmsUsersArea)
	}
	router.Run(":3007")
}
