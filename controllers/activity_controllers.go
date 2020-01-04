package controllers

import (
	"../models"
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"../utils"
)

func GetActivity(c *gin.Context) {
	activity := structs.GetActivity{}
	user := structs.GetActivityUser{}

	t := structs.Component{}

	response := structs.JsonResponse{}

	err := c.BindQuery(&activity)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}
		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {
		data, err_user := models.GetListActivity(activity)

		var id_activity_schedule []int64
		for k := range data{
			id_activity_schedule = append(id_activity_schedule, data[k].IdActivitySchedule)
		}

		user.IdActivitySchedule = id_activity_schedule
		dataUser, user_err := models.GetListActivity(activity)

		fmt.Println("data", dataUser)
		fmt.Println("err", user_err)
		response.Data = data

		if err_user != nil {
			response.ApiMessage = err_user.Error()
			c.JSON(400, response)
		}

		response.ApiMessage = t.GetMessageSucc()
		response.ApiStatus = 1
		response.Data = activity
		c.JSON(http.StatusOK, response)
		//for k := range data {
			//	id = append(id, data[k].IdPembeli, data[k].IdPenjual)
			//}
	}
}

func GetActivitys(c *gin.Context) {
	ayams := structs.ActivityList{}
	list := structs.ListBinds{}

	limit := c.Query("limit")
	offset := c.Query("offset")

	t := structs.Component{}
	response := structs.JsonResponse{}

	err := c.BindQuery(&ayams)
	err2 := c.BindQuery(&list)
	if err != nil || err2 != nil {
		var m string
		if err != nil {
			m = m + err.Error()
		}
		if err2 != nil {
			m = m + err2.Error()
		}
		response.ApiMessage = "validate " + m
		c.JSON(400, response)
	} else {
		data, errc := models.GetsActivitys(ayams, limit, offset, list)

		//list.IdActivitySchedule = int64(len(id_schedule))
		//list.UsersActivity = int64(len(user_activity))
		//datas, err_user := models.GetUserAyam(list)
		//
		//fmt.Println("data_user", datas)
		//fmt.Println("err_user", err_user)

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

func GetActivityList2(c *gin.Context) {
	listActivity := structs.GetActivityList{}
	//listActivityBind := structs.CreateActivityUserBind{}
	t := structs.Component{}

	limit := c.Query("limit")
	offset := c.Query("offset")

	response := structs.JsonResponse{}

	err := c.BindQuery(&listActivity)
	//err2 := c.BindQuery(&listActivityBind)

	//|| err2 != nil
	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}
		//if err2 != nil {
		//	mess = mess + err2.Error()
		//}
		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {
		data, errc := models.GetActivityList2(listActivity, limit, offset)
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

//func GetActivityList(c *gin.Context) {
//	id_mst_outlet := c.Query("id_mst_outlet")
//	created_at := c.Query("created_at")
//	id_cms_users := c.Query("id_cms_users")
//	limit := c.Query("limit")
//	offset := c.Query("offset")
//
//	response := structs.JsonResponse{}
//	var id_cms_users_ary []int
//
//	if id_cms_users != "" {
//		strs := strings.Split(id_cms_users, ",")
//		id_cms_users_ary = make([]int, len(strs))
//		for i := range id_cms_users_ary {
//			id_cms_users_ary[i], _ = strconv.Atoi(strs[i])
//		}
//	}
//	response = models.GetActivityList(created_at, id_cms_users, id_mst_outlet, limit, offset)
//
//	if response.ApiStatus == 1 {
//		c.JSON(http.StatusOK, response)
//	} else {
//		c.JSON(500, response)
//	}
//}

func GetDetail(c *gin.Context) {
	id_activity_detail := c.Query("id_activity_detail")

	response := structs.JsonResponse{}

	response = models.GetDetail(id_activity_detail)

	if response.ApiStatus == 1 {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}


func GetActivityStatus(c *gin.Context) {

	responses := structs.JsonResponse{}
	responses = models.GetactivityStatus()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}
}

func GetActivityType(c *gin.Context) {

	response := structs.JsonResponse{}
	response = models.GetActivityType()

	if response.ApiStatus == 1 {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func CreateActivity(c *gin.Context) {

	activity := structs.CreateActivity{}
	activity_sec := structs.CreateActivitySchedule{}
	activity_user := structs.CreateActivityUserBind{}

	var t  = structs.Component{}

	response := structs.JsonResponse{}

	err := c.ShouldBind(&activity)
	err2 := c.ShouldBind(&activity_sec)
	err3 := c.ShouldBind(&activity_user)

	if err != nil || err2 != nil || err3 != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}
		if err2 != nil {
			mess = mess + err2.Error()
		}
		if err3 != nil {
			mess = mess + err3.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {
		data, err_order := models.CreateActivitys2(activity, activity_sec, activity_user)

		response.Data = data

		if err_order != nil {
			response.ApiMessage = t.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageSucc()
			c.JSON(http.StatusOK, response)
		}
	}

}

//func GetUserActivity(c *gin.Context){
//
//	dataUser := structs.GetUserForActivity{}
//	response := structs.JsonResponse{}
//	t := structs.Component{}
//	err := c.BindQuery(&dataUser)
//
//	fmt.Println(dataUser)
//
//	if err != nil {
//		fmt.Println(err.Error())
//		response.ApiMessage = err.Error()
//		c.JSON(400, response)
//	} else {
//		data := models.GetUsersActivity(dataUser)
//		response.Data = data
//		if err != nil {
//			response.ApiMessage = err.Error()
//			c.JSON(400, response)
//		} else {
//			if len(data) >= 1 {fmt.Println("ada ", data[0].Id)} else {fmt.Println("gak ada", response.Data)}
//			response.ApiMessage = t.GetMessageSucc()
//			c.JSON(http.StatusOK, response)
//		}
//	}
//
//	//response := structs.JsonResponse{}
//	//response =
//	//if response.ApiStatus == 1 {
//	//	c.JSON(http.StatusOK, response)
//	//} else {c.JSON(400, response)}
//}