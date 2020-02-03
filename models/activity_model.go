package models

import (
	"../structs"
	"errors"
	"fmt"
)

func GetActivityList2(activityList structs.GetActivityList, limit string, offset string) ([]structs.GetActivityList, error) {
	fmt.Println(activityList)
	data := []structs.GetActivityList{}
	get := idb.DB.Table("activity_user").
		Select("activity_user.id, to_char(activity_user.created_at, 'YYYY-MM-DD HH24:MI') as created_at, " +
			"activity_user.id_activity_schedule, activity_user.id_cms_users," +
			"cms_users.name, cms_users.id_cms_privileges," + "cms_privileges.name as privileges, " +
			"activity_report.id_activity_mst_status, activity_mst_status.status, activity.id_mst_outlet, mst_outlet.outlet_name," +
			"activity_schedule.id_activity, activity.location," +
			"activity.id_activity_mst_type, activity_mst_type.type").
		Joins("join activity_schedule on activity_user.id_activity_schedule = activity_schedule.id").
		Joins("join cms_users on activity_user.id_cms_users = cms_users.id").
		Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id").
		Joins("join activity_report on activity_schedule.id = activity_report.id_activity_schedule").
		Joins("join activity_mst_status on activity_report.id_activity_mst_status = activity_mst_status.id").
		Joins("join activity on activity_schedule.id_activity = activity.id").
		Joins("join activity_mst_type on activity.id_activity_mst_type = activity_mst_type.id").
		Joins("join mst_outlet on activity.id_mst_outlet = mst_outlet.id")

	if limit != "" {
		get = get.Limit(limit)
	}
	if offset != "" {
		get = get.Offset(offset)
	}
	if activityList.Id != nil {
		get = get.Where("activity_user.id in (?)", int(*activityList.Id))
	}
	if activityList.IdMstOutlet != nil {
		get = get.Where("activity.id_mst_outlet in (?)", int(*activityList.IdMstOutlet))
	}
	if activityList.Name != nil {
		get = get.Where("activity_user.name LIKE", "%"+*activityList.Name+"%")
	}
	if activityList.CreatedAt != nil {
		get = get.Where("activity_user.created_at::text LIKE ?", "%"+*activityList.CreatedAt+"%")
	}

	err := get.Find(&data).Error

	fmt.Println("get orders", data)

	return data, err
}

func GetListActivity(listActivity structs.GetActivity) ([]structs.GetActivityBinds, error) {
	data := []structs.GetActivityBinds{}
	getData := idb.DB.Table("activity_user").Select("activity_user.id, activity_user.created_at, " +
		"activity_user.id_activity_schedule, activity_user.id_cms_users,a.name as name_user, " +
		"activity_schedule.id_activity, activity_schedule.id_cms_users as id_cms_users_schedule, " +
		"activity_schedule.start_date, activity_schedule.end_date, activity_schedule.started, " +
		"activity_schedule.ended, activity_schedule.note, activity.id_mst_outlet, activity.id_activity_mst_type, " +
		"activity_mst_type.type, activity.location, activity.id_cms_users as id_cms_users_spv," +
		" mst_outlet.outlet_name").
		Joins("join cms_users as a on activity_user.id_cms_users = a.id").
		Joins("join activity_schedule on activity_user.id_activity_schedule = activity_schedule.id").
		Joins("join activity on activity_schedule.id_activity = activity.id").
		Joins("join activity_mst_type on activity.id_activity_mst_type = activity_mst_type.id").
		Joins("join mst_outlet on activity.id_mst_outlet = mst_outlet.id")

	if listActivity.IdCmsUsers != nil {
		getData = getData.Where("activity_user.id_cms_users in (?)", int64(*listActivity.IdCmsUsers))
	}
	if listActivity.Id != nil {
		getData = getData.Where("activity_user.id in (?)", int(*listActivity.Id))
	}
	err := getData.Find(&data).Error

	fmt.Println("get s", data)

	return data, err
}

//func GetActivityList(id_mst_outlet string, created_at string, id_cms_users string, limit string, offset string) structs.JsonResponse {
//
//	var (
//		activity []structs.GetActivityList
//		t structs.Component
//	)
//
//	response := structs.JsonResponse{}
//	err := idb.DB.Table("activity_user").Select("activity_user.id, to_char(activity_user.created_at, 'YYYY-MM-DD HH24:MI') as created_at, activity_user.id_activity_schedule, activity_user.id_cms_users," + "cms_users.name, cms_users.id_cms_privileges," + "cms_privileges.name as privileges, activity_report.id_activity_mst_status, activity_mst_status.status," + "activity_schedule.id_activity, activity.location," + "activity.id_activity_mst_type, activity_mst_type.type")
//
//	err = err.Joins("join activity_schedule on activity_user.id_activity_schedule = activity_schedule.id")
//	err = err.Joins("join cms_users on activity_user.id_cms_users = cms_users.id")
//	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
//	err = err.Joins("join activity_report on activity_schedule.id = activity_report.id_activity_schedule")
//	err = err.Joins("join activity_mst_status on activity_report.id_activity_mst_status = activity_mst_status.id")
//	err = err.Joins("join activity on activity_schedule.id_activity = activity.id")
//	err = err.Joins("join activity_mst_type on activity.id_activity_mst_type = activity_mst_type.id")
//
//	if created_at != "" {
//		err = err.Where("activity_user.created_at::text LIKE ?", "%"+created_at+"%")
//	}
//	if id_mst_outlet != "" {
//		err = err.Where("activity.id_mst_outlet = ?", id_mst_outlet)
//	}
//	if id_cms_users != "" {
//		err = err.Where("activity.id_cms_users = ?", id_cms_users)
//	}
//	if limit != "" {
//		err = err.Limit(limit)
//	}
//	if offset != "" {
//		err = err.Offset(offset)
//	}
//
//	err = err.Find(&activity)
//	errx := err.Error
//
//	if errx != nil {
//
//		response.ApiStatus = 0
//		response.ApiMessage = errx.Error()
//		response.Data = nil
//
//	} else {
//		response.ApiStatus = 1
//		response.ApiMessage = t.GetMessageSucc()
//		response.Data = activity
//	}
//	return response
//}

func GetDetail(id_activity_detail string) structs.JsonResponse {
	var (
		detail structs.GetDetail
		t      structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table("activity_user").Select("activity_user.id, to_char(activity_user.created_at, 'YYYY-MM-DD HH24:MI') as created_at, " +
		"activity_user.id_activity_schedule, activity_user.id_cms_users," + "activity_schedule.start_date, activity_schedule.end_date," +
		"cms_users.name, cms_users.id_cms_privileges, cms_privileges.name as privileges," + "activity_report.brosur," +
		" activity_report.note, activity_report.lat, activity_report.lng, activity_report.id_activity_mst_status, " +
		"activity_mst_status.status," + " activity_schedule.id_activity," + "activity.location, " +
		"activity.id_activity_mst_type, activity_mst_type.type," + "activity.id_mst_outlet, mst_outlet.outlet_id," +
		" mst_outlet.outlet_name, mst_outlet.outlet_address, mst_outlet.outlet_status")

	err = err.Joins("join activity_schedule on activity_user.id_activity_schedule = activity_schedule.id")
	err = err.Joins("join cms_users on activity_user.id_cms_users = cms_users.id")
	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("join activity_report on activity_schedule.id = activity_report.id_activity_schedule")
	err = err.Joins("join activity_mst_status on activity_report.id_activity_mst_status = activity_mst_status.id")
	err = err.Joins("join activity on activity_schedule.id_activity = activity.id")
	err = err.Joins("join activity_mst_type on activity.id_activity_mst_type = activity_mst_type.id")
	err = err.Joins("join mst_outlet on activity.id_mst_outlet = mst_outlet.id")

	err = err.Where("activity_user.id	 = ?", id_activity_detail)

	err = err.Find(&detail)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = detail
	}
	return response
}

//func CreateActivitys(id_mst_outlet string, id_activity_mst_type string, id_cms_users string, id_cms_users_other string,  location string, lat string, lng string) structs.JsonResponse {
//
//	var(
//		activityCreate structs.CreateActivity
//		//cekIdActivity structs.CheckIdActivity
//		t structs.Component
//	)
//
//	response := structs.JsonResponse{}
//	check := idb.DB.Table("activity").Select("count(activity.id) as count_id")
//	//check = check.Where("activity.location = ?", location)
//	//check = check.Where()
//
//	//check = check.First(&cekIdActivity)
//
//	checkx := check.Error
//
//	if checkx == nil {
//		fmt.Println("cek id ga error")
//		//if cekIdActivity.CountId == 0 {
//			id_mst_outlet_conv, _ := strconv.Atoi(id_mst_outlet)
//			id_activity_mst_type_conv, _ := strconv.Atoi(id_activity_mst_type)
//			id_cms_users_conv, _ := strconv.Atoi(id_cms_users)
//			id_cms_users_other_conv, _ := strconv.Atoi(id_cms_users_other)
//
//			activityCreate.CreatedAt = t.GetTimeNow()
//			activityCreate.IdMstOutlet = id_mst_outlet_conv
//			activityCreate.IdActivityMstType = id_activity_mst_type_conv
//			activityCreate.IdCmsUsers = id_cms_users_conv
//			activityCreate.IdCmsUsersOther = id_cms_users_other_conv
//			activityCreate.Location = location
//			activityCreate.Lat = lat
//			activityCreate.Lng = lng
//
//			err := idb.DB.Table("activity").Create(&activityCreate)
//
//			errx := err.Error
//
//			if errx != nil {
//				fmt.Println("gagal buat rekomendasi")
//				response.ApiMessage = t.GetMessageErr()
//			} else {
//				response.ApiStatus = 1
//				response.ApiMessage = t.GetMessageSucc()
//				response.Data = activityCreate
//			}
//		//} else {
//		//	response.ApiMessage = "gagal"
//		//}
//	} else {
//		fmt.Println(checkx)
//	}
//	return response
//}

func GetactivityStatus() structs.JsonResponse {

	var (
		statusActivity []structs.GetActivityStatus
		t              structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table("activity_mst_status")

	err = err.Find(&statusActivity)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = statusActivity
	}
	return response
}

func GetActivityType() structs.JsonResponse {
	var (
		statusType []structs.GetActivityType
		t          structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("activity_mst_type")
	err = err.Find(&statusType)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = statusType
	}
	return response
}

func CreateActivitys2(activity structs.CreateActivity,
	activity_sec structs.CreateActivitySchedule,
	activity_user structs.CreateActivityUserBind) (structs.CreateActivity, error) {

	var err error
	var t = structs.Component{}
	tx := idb.DB.Begin()

	if err = tx.Error; err != nil {
		fmt.Println("err start tx", err.Error())
		return activity, err
	}

	// insert ke table activity
	// id cms users adalah pembuat activity
	activity.CreatedAt = t.GetTimeNow()
	if err = tx.Table("activity").Create(&activity).Error; err != nil {
		tx.Rollback()
		return activity, err
	}

	//
	activity_sec.IdActivity = activity.Id       // ambil id dari table activity
	activity_sec.CreatedAt = activity.CreatedAt // ambil tanggal dibuat dari tabel activity
	if err = tx.Table("activity_schedule").Create(&activity_sec).Error; err != nil {
		tx.Rollback()
		return activity, err
	}

	activity_user.IdActivitySchedule = activity_sec.Id // ambil id activity dari activity_schedule
	activity_user.CreatedAt = activity.CreatedAt       // ambil tanggal dibuat dari tabel activity
	for i := range activity_user.IdCmsUsers {          // perulangan untuk id cms users dari activity_user
		activity_user_create := structs.CreateActivityUser{}                       // pembuatan variable untuk ke dalam struct CreateActivityUsers
		activity_user_create.IdCmsUsers = activity_user.IdCmsUsers[i]              // perulangan input untuk id cms users atau cfa
		activity_user_create.CreatedAt = activity_user.CreatedAt                   // mengambil tanggal dibuat
		activity_user_create.IdActivitySchedule = activity_user.IdActivitySchedule // ambil dari struct CreateActivityUsers
		if err = tx.Table("activity_user").Create(&activity_user_create).Error; err != nil {
			tx.Rollback()
			return activity, err
		}
	}

	tx.Commit()

	return activity, err

}
func GetsActivitys(activity structs.ActivityList,
	limit string, offset string, list structs.ListBinds) ([]structs.ActivityList, error) {
	data := []structs.ActivityList{} // struct list data activity.
	data2 := []structs.ListBinds{}   // struct list data array users.

	getSechedulex := idb.DB.Table("activity_schedule").
		Select("activity_schedule.id as id_activity_schedule, " +
			"to_char(activity_schedule.created_at, 'YYYY-MM-DD') as created_at, " +
			"activity_schedule.id_activity, " +
			"to_char(activity_schedule.start_date, 'YYYY-MM-DD') as start_date, " +
			"to_char(activity_schedule.end_date, 'YYYY-MM-DD') as end_date, " +
			"to_char(activity_schedule.started, 'HH24:MI') as started," +
			"to_char(activity_schedule.ended, 'HH24:MI') as ended, " +
			"activity_schedule.note, " +
			"activity.id_mst_outlet, activity.id_activity_mst_type, activity_mst_type.type, activity.location, " +
			"activity.id_cms_users as id_cms_users_spv, mst_outlet.outlet_name, " +
			"mst_outlet.id_mst_branch, mst_branch.branch_name").
		Joins("join activity on activity_schedule.id_activity = activity.id").
		Joins("join activity_mst_type on activity.id_activity_mst_type = activity_mst_type.id").
		Joins("join mst_outlet on activity.id_mst_outlet = mst_outlet.id").
		Joins("join mst_branch on mst_outlet.id_mst_branch = mst_branch.id").
		Order("activity_schedule.id_activity desc")

	if activity.IdActivity != nil {
		getSechedulex = getSechedulex.Where("activity_schedule.id_activity in (?)", int(*activity.IdActivity))
	}
	if activity.CreatedAt != "" {
		getSechedulex = getSechedulex.Where("activity_schedule.created_at LIKE", "%"+activity.CreatedAt+"%")
	}
	if activity.IdMstOutlet != nil {
		getSechedulex = getSechedulex.Where("activity.id_mst_outlet in (?)", int(*activity.IdMstOutlet))
	}
	if activity.IdMstBranch != nil {
		getSechedulex = getSechedulex.Where("mst_outlet.id_mst_branch in (?)", int(*activity.IdMstBranch))
	}

	if limit != "" {
		getSechedulex = getSechedulex.Limit(limit)
	}
	if offset != "" {
		getSechedulex = getSechedulex.Offset(offset)
	}

	getSechedule := getSechedulex.Find(&data).Error

	// kesalahan query
	if getSechedule != nil {
		return data, errors.New("Kesalahan Server")
	}
	// data kosong
	if len(data) <= 0 {
		return data, errors.New("Tidak ada data")
	}

	var idSchedule []int64 // grouping by id_activity_schedule

	for i := range data { // looping 'i' base on data variable
		// idSchedule (arrayList,slice) base load data in id_activity_schedule
		idSchedule = append(idSchedule, data[i].IdActivitySchedule)
	}

	idb.DB.Table("activity_user").
		Select("activity_user.id as id_activity_users, to_char(activity_user.created_at, 'YYYY-MM-DD HH24:MI') as created_at, "+
			"activity_user.id_activity_schedule, "+"activity_user.id_cms_users , "+
			"cms_users.name as name_user").
		Joins("join cms_users on activity_user.id_cms_users = cms_users.id").
		Where("activity_user.id_activity_schedule in (?)", idSchedule).Find(&data2)

	for i := range data {

		for x := range data2 {
			// penyesuaian data antara struct
			if data[i].IdActivitySchedule == data2[x].IdActivitySchedule {

				data[i].DataUser = append(data[i].DataUser, data2[x])
			}
		}

	}
	var err error
	fmt.Println("get activity", data)
	return data, err

}
