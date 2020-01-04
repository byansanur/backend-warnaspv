package structs

type GetActivityList struct {
	Id                		*int    `json:"id" form:"id"`
	CreatedAt         		*string `json:"created_at" form:"created_at" binding:"-"`
	IdActivitySchedule 		*int `json:"id_activity_schedule" form:"id_activity_schedule"`
	IdCmsUsers 				*int `json:"id_cms_users" form:"id_cms_users_activity"`
	Name 					*string `json:"name" form:"name"`
	IdCmsPrivileges 		*int `json:"id_cms_privileges" form:"id_cms_privileges"`
	Privileges 				*string `json:"privileges" form:"privileges"`
	IdActivityMstStatus 	*int `json:"id_activity_mst_status" form:"id_activity_mst_status"`
	Status 					*string `json:"status" form:"status"`
	IdActivity 				*int 	`json:"id_activity" form:"id_activity"`
	Location 				*string 	`json:"location" form:"location"`
	IdActivityMstType 		*int 	`json:"id_activity_mst_type" form:"id_activity_mst_type"`
	Type 					*string 	`json:"type" form:"type"`
	IdMstOutlet				*int `json:"id_mst_outlet" form:"id_mst_outlet"`
	OutletName				*string `json:"outlet_name" form:"outlet_name"`
}

//
//type GetActivityBind struct {
//	Id int	`json:"id" form:"id"`
//	CreatedAt string `json:"created_at" form:"created_at"`
//	IdActivitySchedule int `json:"id_activity_schedule" form:"-"`
//	IdCmsUsers []int64 `json:"id_cms_users" form:"id_cms_users_activities'"`
//	NameUser *string `json:"name_user" form:"name_user"`
//	IdActivity int `json:"id_activity" form:"id_activity"`
//	IdCmsUsersSchedule int `json:"id_cms_users_schedule" form:"id_cms_users_schedule"`
//	StartDate string `json:"start_date" form:"start_date"`
//	EndDate string `json:"end_date" form:"end_date"`
//	Started string `json:"started" form:"started"`
//	Ended string `json:"ended" form:"ended"`
//	Note string `json:"note" form:"note"`
//	IdMstOutlet int `json:"id_mst_outlet" form:"id_mst_outlet"`
//	IdActivityMstType int `json:"id_activity_mst_type" form:"id_activity_mst_type" binding:"-"`
//	Type string `json:"type" form:"type"`
//	Location string `json:"location" form:"location"`
//	IdCmsUserSpv int `json:"id_cms_user_spv" form:"id_cms_users_spv"`
//	OutletName string `json:"outlet_name" form:"outlet_name"`
//}

type ActivityList struct {
	DataUser []interface{} `json:"data_user"`
	IdActivity int `json:"id_activity"`
	IdActivitySchedule int64 `json:"id_activity_schedule"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Started string `json:"started"`
	Ended string `json:"ended"`
	Note string `json:"note"`
	IdMstOutlet *int64 `json:"id_mst_outlet" form:"id_mst_outlet"`
	IdActivityMstType *int64 `json:"id_activity_mst_type" form:"id_mst_outlet"`
	Type string `json:"type"`
	Location string `json:"location"`
	OutletName string `json:"outlet_name"`
	IdMstBranch *int `json:"id_mst_branch"`
	BranchName string `json:"branch_name"`
}

type AnakAyamList struct {
	Id int64 `json:"id" form:"id" binding:"-"`
	CreatedAt string `json:"created_at" form:"created_at"`
	IdActivitySchedule int64 `json:"id_activity_schedule" form:"id_activity_schedule" binding:"-"`
	IdCmsUsersActivity int64 `json:"id_cms_users_activity" form:"id_cms_users_activity"`
	NameUser string `json:"name_user" form:"name_user"`
}

type ListBinds struct {
	Id int64 `json:"id" form:"id" binding:"-"`
	CreatedAt string `json:"created_at" form:"created_at"`
	IdActivitySchedule int64 `json:"id_activity_schedule" form:"id_activity_schedule_user" binding:"-"`
	IdCmsUsersActivity int64 `json:"id_cms_users_activity" form:"id_cms_users_activity"`
	NameUser string `json:"name_user" form:"name_user"`
}


type GetActivity struct {
	Id *int `json:"id" form:"id"`
	CreatedAt string `json:"created_at" form:"created_at"`
	IdActivitySchedule []int64 `json:"id_activity_schedule" form:"-"`
	IdCmsUsers *int64 `json:"id_cms_users" form:"-"`
	NameUser string `json:"name_user" form:"name_user"`
	IdActivity int `json:"id_activity" form:"-"`
	IdCmsUsersSchedule int `json:"id_cms_users_schedule" form:"-" binding:"-"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate string `json:"end_date" form:"end_date"`
	Started string `json:"started" form:"started"`
	Ended string `json:"ended" form:"ended"`
	Note string `json:"note" form:"note"`
	IdMstOutlet int `json:"id_mst_outlet" form:"-" binding:"-"`
	IdActivityMstType int `json:"id_activity_mst_type" form:"-" binding:"-"`
	Type string `json:"type" form:"type"`
	Location string `json:"location" form:"location"`
	IdCmsUserSpv int `json:"id_cms_user_spv" form:"-" binding:"-"`
	OutletName string `json:"outlet_name" form:"outlet_name"`
}

type GetActivityBinds struct {
	Id *int `json:"id" form:"id"`
	CreatedAt *string `json:"created_at" form:"created_at"`
	IdActivitySchedule int64 `json:"id_activity_schedule" form:"id_schedules"`
	IdCmsUsers int64 `json:"id_cms_users" form:"id_cms_activity_users"`
	NameUser string `json:"name_user" form:"name_user"`
	IdActivity int `json:"id_activity" form:"-"`
	IdCmsUsersSchedule int `json:"id_cms_users_schedule" form:"-" binding:"-"`
	StartDate *string `json:"start_date" form:"start_date"`
	EndDate *string `json:"end_date" form:"end_date"`
	Started *string `json:"started" form:"started"`
	Ended *string `json:"ended" form:"ended"`
	Note *string `json:"note" form:"note"`
	IdMstOutlet *int `json:"id_mst_outlet" form:"-" binding:"-"`
	IdActivityMstType *int `json:"id_activity_mst_type" form:"-" binding:"-"`
	Type *string `json:"type" form:"type"`
	Location *string `json:"location" form:"location"`
	IdCmsUserSpv int `json:"id_cms_user_spv" form:"-" binding:"-"`
	OutletName *string `json:"outlet_name" form:"outlet_name"`
}

type GetActivityUser struct {
	Id *[]int64 `json:"id" form:"id"`
	IdActivitySchedule []int64 `json:"id_activity_schedule" form:"id_activity_schedule" binding:"-"`
	IdCmsUsers *[]int64 `json:"id_cms_users" form:"id_cms_users_act" binding:"-"`
	NameUser *string `json:"name_user" form:"name_user"`
}


type GetActivityList2 struct {
	Id int `json:"id" form:"id"`
	IdMstOutlet int64 `json:"id_mst_outlet" form:"id_mst_outlet"`
	IdActivityMstType int `json:"id_activity_mst_type" form:"id_activity_mst_type"`
	IdCmsUsers int64 `json:"id_cms_users" form:"id_cms_users"` // spv buat activity
	Location string `json:"location" form:"location"`
	Lat string `json:"lat" form:"lat"`
	Lng string `json:"lng" form:"lng"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate string `json:"end_date" form:"end_date"`
	Started string `json:"started" form:"started"`
	Ended string `json:"ended" form:"ended"`
	Note string `json:"note" form:"note"`
	IdCmsUsersActivity []int64 `json:"id_cms_users_activity" form:"id_cms_users_activity"`
	NameUsers string `json:"name_users" form:"name_users"`
}

type GetDetail struct {
	Id 						int 	`json:"id"`
	CreatedAt 				string 	`json:"created_at"`
	IdActivitySchedule 		int 	`json:"id_activity_schedule"`
	IdCmsUsers 				int 	`json:"id_cms_users"`
	StartDate 				string 	`json:"start_date"`
	EndDate 				string 	`json:"end_date"`
	Name 					string 	`json:"name"`
	IdCmsPrivileges 		int 	`json:"id_cms_privileges"`
	Privileges 				string 	`json:"privileges"`
	Brosur 					string 	`json:"brosur"`
	Note 					string 	`json:"note"`
	IdActivityMstStatus 	int 	`json:"id_activity_mst_status"`
	Status 					string 	`json:"status"`
	IdActivity 				int 	`json:"id_activity"`
	Location 				string 	`json:"location"`
	IdActivityMstType 		int 	`json:"id_activity_mst_type"`
	Type 					string 	`json:"type"`
	IdMstOutlet 			int 	`json:"id_mst_outlet"`
	OutletName 				string 	`json:"outlet_name"`
	OutletAddress 			string 	`json:"outlet_address"`
	OutletStatus 			string 	`json:"outlet_status"`
}

type CheckIdActivity struct {
	CountId int `json:"count_id"`
}

type CreateActivity struct {
	Id int `json:"id" form:"id"`
	CreatedAt string `json:"created_at" form:"created_at"`
	IdMstOutlet int `json:"id_mst_outlet" form:"id_mst_outlet"`
	IdActivityMstType int `json:"id_activity_mst_type" form:"id_activity_mst_type"`
	IdCmsUsers int `json:"id_cms_users" form:"id_cms_users" binding:"required"`
	Location string `json:"location" form:"location" binding:"required"`
	Lat string `json:"lat" form:"lat"`
	Lng string `json:"lng" form:"lng"`
	//UpdatedBy *int `json:"updated_by"`
}

type CreateActivitySchedule struct {
	Id int `json:"id" form:"id"`
	CreatedAt string `json:"created_at" form:"created_at"`
	IdActivity int `json:"id_activity" form:"-" binding:"-"`
	StartDate string `json:"start_date" form:"start_date" binding:"required"`
	EndDate string `json:"end_date" form:"end_date"`
	Started string `json:"started" form:"started"`
	Ended string `json:"ended" form:"ended"`
	Note string `json:"note" form:"note"`
	IdCmsUsers int `json:"id_cms_users" form:"id_cms_users"` // pembuat activity for cfa
	//UpdatedBy int `json:"updated_by"`
}

type CreateActivityUser struct {
	Id int `json:"id" form:"id"`
	CreatedAt string `json:"created_at" form:"created_at"`
	IdActivitySchedule int `json:"id_activity_schedule" form:"-" binding:"-"`
	IdCmsUsers int64 `json:"id_cms_users" form:"-"`
}

type CreateActivityUserBind struct {
	Id int `json:"id" form:"id"`
	CreatedAt string `json:"created_at" form:"created_at"`
	IdActivitySchedule int `json:"id_activity_schedule" form:"-" binding:"-"`
	IdCmsUsers []int64 `json:"id_cms_users" form:"id_cms_users_activity"` // id untuk kesi cfa, diambil dari create activityUser
}

type CreateActivityType struct {
	Id int64 `json:"id"`
	CreatedAt string `json:"created_at"`
	Type string `json:"type"`
	IdCmsUsers int `json:"id_cms_users"`
}

type CreateactivityStatus struct {
	Id int `json:"id"`
	CreatedAt string `json:"created_at"`
	Status string `json:"status"`
	IdCmsUsers int `json:"id_cms_users"`
}

type GetActivityStatus struct {
	Id int `json:"id"`
	Status string `json:"status"`
}

type GetActivityType struct {
	Id int `json:"id"`
	Type string `json:"type"`
}

//type GetUserForActivity struct {
//	Id int `json:"id" form:"id"`
//	Name *string `json:"name" form:"name" binding:"required"`
//	Npm *string `json:"npm" form:"npm" binding:"-"`
//	IdCmsPrivileges *int64 `json:"id_cms_privileges" form:"id_cms_privileges" binding:"-"`
//	NamePrivileges *string `json:"name_privileges" form:"name_privileges" binding:"-"`
//	IdMstOutlet *int64 `json:"id_mst_outlet" form:"id_mst_outlet" binding:"required"`
//	OutletName *string `json:"outlet_name" form:"outlet_name"`
//	IdMstBranch *int `json:"id_mst_branch" form:"id_mst_branch" binding:"-"`
//	BranchName *string `json:"branch_name" form:"branch_name"`
//	Status *string `json:"status" form:"status"`
//}

