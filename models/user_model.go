package models

import (
	"../config"
	"../structs"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/dgrijalva/jwt-go"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)
func GetUsers(id_mst_outlet string, id_mst_branch []int, id_cms_privileges string, status string,
	id_oh string, id_spv string, offset string, limit string, name string, privileges_name string) structs.JsonResponse {

	var (
		users []structs.GetUsers
		t     structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("cms_users").Select("cms_users.id, to_char(cms_users.created_at, 'YYYY Mon DD HH24:MI')  as created_at, cms_users.name," +
		"cms_users.npm ,mst_outlet.outlet_name , mst_branch.branch_name ,cms_users.email, cms_privileges.name as privileges_name," +
		"cms_users.id_mst_outlet , mst_outlet.id_mst_branch , cms_users.id_cms_privileges," +
		"(CASE WHEN cms_users.status = 'Y' then 'Active'  ELSE 'No Active' end ) as status," +
		"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +

		"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh," +

		"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
		"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv")

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")
	err = err.Joins("left join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}

	if name != "" {
		err = err.Where("cms_users.name::text ILIKE ?", "%"+name+"%")
	}

	if privileges_name != "" {
		err = err.Where("cms_privileges.name::text ILIKE ?", "%"+privileges_name+"%")
	}
	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_cms_privileges != "" {
		err = err.Where("cms_users.id_cms_privileges = ?", id_cms_privileges)
	}
	if id_oh != "" {
		err = err.Where("id_cms_users_oh = ?", id_oh)
	}

	if id_spv != "" {
		err = err.Where("id_cms_users_spv = ?", id_spv)
	}

	if status != "" {
		err = err.Where("cms_users.status LIKE ?", "%"+status+"%")
	}

	err = err.Order("cms_users.name asc")

	err = err.Find(&users)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = users
	}

	return response
}

func GetUsers2(id_mst_outlet string, id_mst_branch []int, id_cms_privileges string, status string,
	id_oh string, id_spv string, offset string, limit string, name string, privileges_name string) structs.JsonResponse {

	var (
		users []structs.GetUsers
		t     structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("cms_users").Select("cms_users.id, to_char(cms_users.created_at, 'YYYY Mon DD HH24:MI')  as created_at, cms_users.name," +
		"cms_users.npm ,mst_outlet.outlet_name , mst_branch.branch_name ,cms_users.email, cms_privileges.name as privileges_name," +
		"cms_users.id_mst_outlet , mst_outlet.id_mst_branch , cms_users.id_cms_privileges," +
		"(CASE WHEN cms_users.status = 'Y' then 'Active'  ELSE 'No Active' end ) as status," +
		"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +

		"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh," +

		"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
		"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv")

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")
	err = err.Joins("left join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")
	err = err.Where("cms_privileges.id in ('3', '6') and cms_users.status in ('Y')")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}

	if name != "" {
		err = err.Where("cms_users.name::text LIKE ?", "%"+name+"%")
	}

	if privileges_name != "" {
		err = err.Where("cms_privileges.name::text ILIKE ?", "%"+privileges_name+"%")
	}
	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_cms_privileges != "" {
		err = err.Where("cms_users.id_cms_privileges = ?", id_cms_privileges)
	}
	if id_oh != "" {
		err = err.Where("id_cms_users_oh = ?", id_oh)
	}

	if id_spv != "" {
		err = err.Where("id_cms_users_spv = ?", id_spv)
	}

	if status != "" {
		err = err.Where("cms_users.status LIKE ?", "%"+status+"%")
	}

	err = err.Order("cms_users.name asc")

	err = err.Find(&users)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = users
	}

	return response
}



func DownloadUsers(id_mst_branch []int, offset string, limit string) structs.JsonResponse {

	var (
		users []structs.GetUsers
		t     structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("cms_users").Select("cms_users.id,cms_users.created_at, cms_users.name," +
		"cms_users.npm ,mst_outlet.outlet_name , mst_branch.branch_name ,cms_users.email, cms_privileges.name as privileges_name," +
		"cms_users.id_mst_outlet , mst_outlet.id_mst_branch , cms_users.id_cms_privileges," +
		"(CASE WHEN cms_users.status = 'Y' then 'Active'  ELSE 'No Active' end ) as status," +
		"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +

		"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh," +

		"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
		"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv")

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")
	err = err.Joins("left join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}

	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch  IN (?)", id_mst_branch)
	}

	err = err.Order("cms_users.id desc")

	err = err.Find(&users)
	errx := err.Error

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "Created at")
	f.SetCellValue("Sheet1", "B1", "NPM")
	f.SetCellValue("Sheet1", "C1", "Name")
	f.SetCellValue("Sheet1", "D1", "Email")
	f.SetCellValue("Sheet1", "E1", "PrivilegesName")
	f.SetCellValue("Sheet1", "F1", "OutletName")
	f.SetCellValue("Sheet1", "G1", "BranchNAme")
	f.SetCellValue("Sheet1", "H1", "Status")
	f.SetCellValue("Sheet1", "I1", "Spv Name")
	f.SetCellValue("Sheet1", "J1", "Oh Name")

	for i, _ := range users {

		rows := 2

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+rows), users[i].CreatedAt)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+rows), users[i].Npm)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+rows), users[i].Name)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+rows), users[i].Email)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+rows), users[i].PrivilegesName)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+rows), users[i].OutletName)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+rows), users[i].BranchName)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+rows), users[i].Status)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+rows), users[i].Spv)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+rows), users[i].Oh)

	}

	f.SetActiveSheet(index)
	folderPath := "./files/"

	filename := "users" + time.Now().Format("20060102150405") + ".xlsx"

	os.MkdirAll(folderPath, os.ModePerm)

	// Save xlsx file by the given path.
	errs := f.SaveAs(folderPath + filename)
	if errs != nil {
		fmt.Println(errs)
	}

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	}

	if len(users) <= 0 {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = filename

	}

	return response

}

func GetUsersStatus(id_mst_outlet string, id_mst_branch string, offset string, limit string) structs.JsonResponse {

	var (
		users []structs.UsersStatus
		t     structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("cms_users").Select("mst_branch.id as id_branch, mst_branch.branch_name, mst_outlet.id as id_outlet," +
		"mst_outlet.outlet_name," +
		"count(cms_users.id) as total," +
		"(select count(cms_users.id) from cms_users where cms_users.id_mst_outlet = mst_outlet.id and cms_users.status = 'Y' ) as active," +
		"(select count(cms_users.id) from cms_users where cms_users.id_mst_outlet = mst_outlet.id and cms_users.status = 'N' ) as no_active")

	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	err = err.Group("mst_branch.id, mst_outlet.id")
	//err = err.Order("cms_users.created_at desc")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}

	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}
	if id_mst_branch != "" {
		err = err.Where("mst_outlet.id_mst_branch = ?", id_mst_branch)
	}

	err = err.Find(&users)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = users
	}

	return response
}

func CreateUsers(npm string, name string, email string, password string,
	files multipart.File, header *multipart.FileHeader, status string, id_cms_privileges string, id_mst_outlet string, created_by string,
	id_cms_users_oh string, id_cms_users_spv string, id_cms_users_sub_dept string,
	id_mst_branch []int) structs.JsonResponse {

	var (
		users       structs.CreateUsers
		CheckNpm    structs.CheckNpm
		UsersCabang structs.CreateCmsUsersCabang
		UsersArea   structs.CreateCmsUsersArea
		t           structs.Component
	)

	response := structs.JsonResponse{}

	check := idb.DB.Table("cms_users").Select("count(cms_users.id) as count_npm")

	check = check.Where("cms_users.npm = ?", npm)
	check = check.First(&CheckNpm)
	checkx := check.Error

	if checkx == nil {
		fmt.Println("cek npm ga err")
		if CheckNpm.CountNpm == 0 {

			encrptPassword, _ := EncryptPassword(password)
			created_by_conv, _ := strconv.Atoi(created_by)
			id_mst_outlet_conv, _ := strconv.Atoi(id_mst_outlet)
			id_cms_privileges_conv, _ := strconv.Atoi(id_cms_privileges)

			url := UploadImage("user", fmt.Sprint(npm), files, header)

			if url != "" {

				fmt.Println("Foto Ga null")

				users.CreatedAt = t.GetTimeNow()
				users.Npm = npm
				users.Name = name
				users.Status = status
				users.Password = encrptPassword
				users.Photo = url
				users.Email = email
				users.CreatedBy = created_by_conv
				users.IdCmsPrivileges = id_cms_privileges_conv
				users.IdMstOutlet = id_mst_outlet_conv

				err := idb.DB.Table("cms_users").Create(&users)

				errx := err.Error

				if errx == nil {
					id_cms_users_oh_conv, _ := strconv.Atoi(id_cms_users_oh)
					id_cms_users_spv_conv, _ := strconv.Atoi(id_cms_users_spv)
					id_cms_users_sub_dept_conv, _ := strconv.Atoi(id_cms_users_sub_dept)

					UsersCabang.CreatedAt = t.GetTimeNow()
					UsersCabang.IdCmsUsers = users.Id
					UsersCabang.IdCmsUsersOh = id_cms_users_oh_conv
					UsersCabang.IdCmsUsersSpv = id_cms_users_spv_conv
					UsersCabang.IdCmsUsersSubDept = id_cms_users_sub_dept_conv
					UsersCabang.CreatedBy = created_by_conv

					errs := idb.DB.Table("cms_users_cabang").Create(&UsersCabang)

					cabang := errs.Error

					if id_cms_privileges_conv == 10 || id_cms_privileges_conv == 11 || id_cms_privileges_conv == 12 {

						//var id_int = []int{}
						for _, i := range id_mst_branch {

							//id_int = append(id_int, i)

							UsersArea.IdCmsUsers = users.Id
							UsersArea.IdMstBranch = i
							errs_area := idb.DB.Table("cms_users_area").Create(&UsersArea)
							if errs_area != nil {
								fmt.Println("done")
							}
						}

					}

					if cabang != nil {

						response.ApiStatus = 0
						response.ApiMessage = errx.Error()
						response.Data = nil

					} else {
						response.ApiStatus = 1
						response.ApiMessage = t.GetMessageSucc()
						response.Data = users
					}
				}

			}

		} else {
			response.ApiStatus = 0
			response.ApiMessage = "Npm Already Used"
			response.Data = users
		}

	} else {
		fmt.Println(checkx)
	}

	return response
}

func UpdateUsers(id string, name string, npm, email string, password string,
	files multipart.File, header *multipart.FileHeader, status string, id_cms_privileges string, id_mst_outlet string,
	id_cms_users_oh string, id_cms_users_spv string, id_mst_branch []int, id_cms_users_sub_dept string) structs.JsonResponse {

	var (
		users structs.UpdateUsers
		//CheckNpm    structs.CheckNpm
		UsersCabang structs.CreateCmsUsersCabang
		DeleteArea  structs.DeleteCmsUsersArea
		UsersArea   structs.CreateCmsUsersArea
		t           structs.Component
	)

	response := structs.JsonResponse{}

	encrptPassword, _ := EncryptPassword(password)
	id_mst_outlet_conv, _ := strconv.Atoi(id_mst_outlet)
	id_conv, _ := strconv.Atoi(id)
	id_cms_privileges_conv, _ := strconv.Atoi(id_cms_privileges)

	if files != nil && header != nil {
		url := UploadImage("user", fmt.Sprint(name), files, header)
		users.Photo = url
	}
	if password != "" {
		users.Password = encrptPassword
	}

	fmt.Println("Foto Ga null")

	users.Status = status
	users.Name = name
	users.Npm = npm
	users.Email = email
	users.IdCmsPrivileges = id_cms_privileges_conv
	users.IdMstOutlet = id_mst_outlet_conv

	err := idb.DB.Table("cms_users").Where("cms_users.id = ?", id_conv).Updates(&users)

	errx := err.Error

	fmt.Println("id_mst_branch ", id_mst_branch)
	fmt.Println("id_mst_branch len ", len(id_mst_branch))

	if len(id_mst_branch) >= 1 {

		errdel := idb.DB.Table("cms_users_area").Unscoped().Where("id_cms_users = ? ", id_conv).Delete(&DeleteArea).Error

		if errdel == nil {

			for _, i := range id_mst_branch {

				UsersArea.IdCmsUsers = id_conv
				UsersArea.IdMstBranch = i
				errs_area := idb.DB.Table("cms_users_area").Create(&UsersArea)
				if errs_area != nil {
					fmt.Println("done")
				}

			}
		}

	}

	if errx == nil {
		id_cms_users_oh_conv, _ := strconv.Atoi(id_cms_users_oh)
		id_cms_users_spv_conv, _ := strconv.Atoi(id_cms_users_spv)
		id_cms_users_sub_dept_conv, _ := strconv.Atoi(id_cms_users_sub_dept)

		UsersCabang.IdCmsUsers = id_conv
		UsersCabang.IdCmsUsersOh = id_cms_users_oh_conv
		UsersCabang.IdCmsUsersSpv = id_cms_users_spv_conv
		UsersCabang.IdCmsUsersSubDept = id_cms_users_sub_dept_conv

		errs := idb.DB.Table("cms_users_cabang").Where("id_cms_users = ?", id_conv).Updates(&UsersCabang)

		cabang := errs.Error
		if cabang != nil {

			response.ApiStatus = 0
			response.ApiMessage = errx.Error()
			response.Data = nil

		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageSucc()
			response.Data = users
		}

	}

	//} else {
	//	response.ApiStatus = 0
	//	response.ApiMessage = "Npm Already Used"
	//	response.Data = users
	//
	//}

	//}else {
	//	fmt.Println(checkx)
	//}

	return response
}

func GetUserDetail(id string) structs.JsonResponse {

	var (
		users structs.GetUsers
		t     structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table("cms_users").Select("cms_users.id, cms_users.created_at, cms_users.name," +
		"cms_users.npm ,mst_outlet.outlet_name , mst_branch.branch_name , cms_privileges.name as privileges_name," +
		"cms_users.id_mst_outlet , mst_outlet.id_mst_branch , cms_users.id_cms_privileges," +
		"cms_users.status," +
	//"(CASE WHEN cms_users.status = 'Y' then 'Active'  ELSE 'No Active' end ) as status," +
		"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +
		"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh," +
		"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
		"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv," +
		"(select cms_users_cabang.id_cms_users_sub_dept from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_sub_dept," +
		"(select name from cms_users where cms_users.id = id_cms_users_sub_dept )  as sub_dept")

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")
	err = err.Joins("left join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")

	err = err.Where("cms_users.id = ?", id)

	err = err.First(&users)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = users
	}

	return response
}

func Login(npm string, password string) structs.JsonResponse {

	var (
		userslogin structs.CheckUsersLogin
		users      structs.GetUsersLogin
		//t     structs.Component
	)

	response := structs.JsonResponse{}

	checknpm := idb.DB.Table("cms_users").Select("id,npm,password")
	checknpm = checknpm.Where("cms_users.status = 'Y' ")
	checknpm = checknpm.Where("npm = ?", npm)
	checknpm = checknpm.Where("cms_users.id_cms_privileges != 3")
	checknpm = checknpm.Where("cms_users.id_cms_privileges != 6")

	checknpm = checknpm.Scan(&userslogin)
	checknpms := checknpm.RecordNotFound()

	fmt.Println(userslogin.Id)

	if checknpms {
		fmt.Println("Npm Ga ketemu")
		response.ApiMessage = "NPM tidak ditemukan"
	} else {

		checkPassword, errPass := DecryptPassword(userslogin.Password)

		if errPass != nil {
			fmt.Println("Password salah")
			response.ApiMessage = "Password salah"
		} else {
			fmt.Println("adsaasd")
			if checkPassword == password {

				fmt.Println("Pass sama")

				err := idb.DB.Table("cms_users").Select("cms_users.id, cms_users.created_at, cms_users.name," +
					"cms_users.npm ,mst_outlet.outlet_name , mst_branch.branch_name , cms_privileges.name as privileges_name," +
					"cms_users.id_mst_outlet , mst_outlet.id_mst_branch , cms_users.id_cms_privileges," +
					"(CASE WHEN cms_users.status = 'Y' then 'Active'  ELSE 'No Active' end ) as status," +
					"cms_users.password," +
					"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +
					"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh," +
					"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
					"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv")

				err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
				err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
				err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")
				err = err.Joins("left join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")

				err = err.Where("cms_users.id = ?", userslogin.Id)

				err = err.First(&users)
				errx := err.Error

				//
				sign := jwt.New(jwt.SigningMethodHS256)
				claims := sign.Claims.(jwt.MapClaims)

				claims["authorized"] = true
				claims["client"] = users.Id
				claims["exp"] = time.Now().Add(time.Minute * 360).Unix()

				token, _ := sign.SignedString(config.JwtKey())
				users.Token = token
				//fmt.Println("token", token)

				if errx == nil {
					response.ApiStatus = 1
					response.ApiMessage = "success login"
					response.Data = users
				} else {
					response.ApiMessage = " err"
				}
			} else {
				response.ApiMessage = "pass salah"
			}
		}

	}

	return response
}

func GetKPIUsers(id_mst_outlet string, id_mst_branch []int, created_at1 string, created_at2 string, limit string, offset string) structs.JsonResponse {

	var (
		leadvisum []structs.GetKPI
		t         structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	//created_at1s := "%"+created_at1+"%"
	//created_at2s := "%"+created_at2+"%"

	var orders = `(select count(id) from "order" orders where orders.id_cms_users = cms_users.id and date(orders.created_at) between '` + created_at1 + `' and '` + created_at2 + `') as orders,`
	var booking = `(select count(id) from "order" orders where orders.id_cms_users = cms_users.id and orders.id_order_mst_status = 5 and date(orders.created_at) between '` + created_at1 + `' and '` + created_at2 + `') as booking,`
	var lead = `(select count(lead.id) from lead where lead.id_cms_users = cms_users.id  and date(lead.created_at) between '` + created_at1 + `' and '` + created_at2 + `') as lead,`
	var lead_visum = `(select count(lead_visum.id) from lead_visum where lead_visum.id_cms_users = cms_users.id  and date(lead_visum.created_at) between '` + created_at1 + `' and '` + created_at2 + `') as lead_visum,`

	err := idb.DB.Table("cms_users").Select("distinct (cms_users.id), cms_users.npm, cms_users.name , mst_outlet.outlet_name," +
		"cms_users.status,cms_privileges.name as privileges_name, mst_branch.branch_name," +
		orders + booking + lead + lead_visum +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "' ) as target," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 1 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_new," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 2 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_hot," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 3 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_working," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 4 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_visit," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 5 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_unqualifed," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 6 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_converted," +
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and date(target_visum.created_at) between '" + created_at1 + "' and '" + created_at2 + "' ) as target_visum," +
		"(select count(target_log.id) from target_log where target_log.id_cms_users = cms_users.id and date(target_log.created_at) between '" + created_at1 + "' and '" + created_at2 + "' ) as target_calls")

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}

	err = err.Find(&leadvisum)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	}

	if len(leadvisum) <= 0 {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = leadvisum

	}

	return response
}

func Dashboard(id_mst_branch []int, created_at string) structs.JsonResponse {

	var (
		dashboard structs.Dashboard
		t         structs.Component
	)

	response := structs.JsonResponse{}

	//new
	user := idb.DB.Table("cms_users").Select("count(cms_users.id) as users")
	user = user.Joins("join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	user = user.Where("cms_users.status = 'Y' ")

	if len(id_mst_branch) >= 1 {
		user = user.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	user = user.Find(&dashboard)
	erruser := user.Error

	lead := idb.DB.Table("lead").Select("count(lead.id) as leads")
	lead = lead.Joins("join mst_outlet on mst_outlet.id = lead.id_mst_outlet")

	if created_at != "" {
		lead = lead.Where("lead.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		lead = lead.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	lead = lead.Find(&dashboard)
	errlead := lead.Error

	deals := idb.DB.Table(order + " orders").Select("count(orders.id) as deals")
	deals = deals.Joins("join mst_outlet on mst_outlet.id = orders.id_mst_outlet")

	if created_at != "" {
		deals = deals.Where("orders.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		deals = deals.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	deals = deals.Find(&dashboard)
	errdeals := deals.Error

	booking := idb.DB.Table(order + " orders").Select("count(orders.id) as booking")
	booking = booking.Joins("join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	booking = booking.Where("orders.id_order_mst_status = 5")

	if created_at != "" {
		booking = booking.Where("orders.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		booking = booking.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	booking = booking.Find(&dashboard)
	errbooking := booking.Error

	target := idb.DB.Table("target").Select("count(target.id) as target")

	if len(id_mst_branch) >= 1 {
		target = target.Where("target.id_mst_branch IN (?)", id_mst_branch)
	}
	target = target.Find(&dashboard)
	errtarget := target.Error

	target_call := idb.DB.Table("target_log").Select("count(target_log.id) as target_call")
	target_call = target_call.Joins("left join target on target.id = target_log.id_target")

	if created_at != "" {
		target_call = target_call.Where("target_log.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		target_call = target_call.Where("target.id_mst_branch IN (?)", id_mst_branch)
	}
	target_call = target_call.Find(&dashboard)
	errtarget_call := target_call.Error

	if erruser != nil && errlead != nil && errdeals != nil && errbooking != nil && errtarget != nil && errtarget_call != nil {

		response.ApiStatus = 0
		response.ApiMessage = erruser.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = dashboard
	}

	return response
}

func KPIUsersDownload(id_mst_outlet string, id_mst_branch []int, created_at1 string, created_at2 string, limit string, offset string) structs.JsonResponse {

	var (
		targetvisum []structs.GetKPI
		t           structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	var orders = `(select count(id) from "order" orders where orders.id_cms_users = cms_users.id and date(orders.created_at) between '` + created_at1 + `' and '` + created_at2 + `') as orders,`
	var booking = `(select count(id) from "order" orders where orders.id_cms_users = cms_users.id and orders.id_order_mst_status = 5 and date(orders.created_at) between '` + created_at1 + `' and '` + created_at2 + `') as booking,`
	var lead = `(select count(lead.id) from lead where lead.id_cms_users = cms_users.id  and date(lead.created_at) between '` + created_at1 + `' and '` + created_at2 + `') as lead,`
	var lead_visum = `(select count(lead_visum.id) from lead_visum where lead_visum.id_cms_users = cms_users.id  and date(lead_visum.created_at) between '` + created_at1 + `' and '` + created_at2 + `') as lead_visum,`

	err := idb.DB.Table("cms_users").Select("distinct (cms_users.id), cms_users.npm, cms_users.name , mst_outlet.outlet_name," +
		"cms_users.status,cms_privileges.name as privileges_name, mst_branch.branch_name," +
		orders + booking + lead + lead_visum +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "' ) as target," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 1 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_new," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 2 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_hot," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 3 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_working," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 4 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_visit," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 5 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_unqualifed," +
		"(select count(target.id) from target where target.id_cms_users = cms_users.id  and target.id_target_mst_status = 6 and date(target.updated_at) between '" + created_at1 + "' and '" + created_at2 + "') as target_converted," +
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and date(target_visum.created_at) between '" + created_at1 + "' and '" + created_at2 + "' ) as target_visum," +
		"(select count(target_log.id) from target_log where target_log.id_cms_users = cms_users.id and date(target_log.created_at) between '" + created_at1 + "' and '" + created_at2 + "' ) as target_calls")

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}

	err = err.Find(&targetvisum)
	errx := err.Error

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "NPM")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Privileges Name")
	f.SetCellValue("Sheet1", "D1", "Users Outlet")
	f.SetCellValue("Sheet1", "E1", "Users Branch")
	f.SetCellValue("Sheet1", "F1", "Lead")
	f.SetCellValue("Sheet1", "G1", "Lead Visum")
	f.SetCellValue("Sheet1", "H1", "Order")
	f.SetCellValue("Sheet1", "I1", "Booking")
	f.SetCellValue("Sheet1", "J1", "Target")
	f.SetCellValue("Sheet1", "K1", "Target New")
	f.SetCellValue("Sheet1", "L1", "Target Hot")
	f.SetCellValue("Sheet1", "M1", "Target Working")
	f.SetCellValue("Sheet1", "N1", "Target Visit")
	f.SetCellValue("Sheet1", "O1", "Target Unqualifed")
	f.SetCellValue("Sheet1", "P1", "Target Converted")
	f.SetCellValue("Sheet1", "Q1", "Target Visum")
	f.SetCellValue("Sheet1", "R1", "Target Calls")

	for i, _ := range targetvisum {

		rows := 2

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+rows), targetvisum[i].Npm)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+rows), targetvisum[i].Name)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+rows), targetvisum[i].PrivilegesName)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+rows), targetvisum[i].OutletName)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+rows), targetvisum[i].BranchName)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+rows), targetvisum[i].Lead)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+rows), targetvisum[i].LeadVisum)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+rows), targetvisum[i].Orders)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+rows), targetvisum[i].Booking)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+rows), targetvisum[i].Target)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+rows), targetvisum[i].TargetNew)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+rows), targetvisum[i].TargetHot)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(i+rows), targetvisum[i].TargetWorking)
		f.SetCellValue("Sheet1", "N"+strconv.Itoa(i+rows), targetvisum[i].TargetVisit)
		f.SetCellValue("Sheet1", "O"+strconv.Itoa(i+rows), targetvisum[i].TargetUnqualifed)
		f.SetCellValue("Sheet1", "P"+strconv.Itoa(i+rows), targetvisum[i].TargetConverted)
		f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i+rows), targetvisum[i].TargetVisum)
		f.SetCellValue("Sheet1", "R"+strconv.Itoa(i+rows), targetvisum[i].TargetCalls)

	}

	f.SetActiveSheet(index)
	folderPath := "./files/"

	filename := "KPI" + time.Now().Format("20060102150405") + ".xlsx"

	os.MkdirAll(folderPath, os.ModePerm)

	// Save xlsx file by the given path.
	errs := f.SaveAs(folderPath + filename)
	if errs != nil {
		fmt.Println(errs)
	}

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	}

	if len(targetvisum) <= 0 {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = filename

	}

	return response
}

func GetUsersArea(id_cms_users string, id_mst_branch string) structs.JsonResponse {

	var (
		users []structs.GetCmsUsersArea
		t     structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("cms_users_area").Select(`cms_users_area.id_cms_users,
cms_users_area.id_mst_branch as id, mst_branch.branch_name,cms_users_area.id `)

	err = err.Joins("join mst_branch on mst_branch.id = cms_users_area.id_mst_branch")

	if id_cms_users != "" {
		err = err.Where("cms_users_area.id_cms_users = ?", id_cms_users)
	}

	err = err.Find(&users)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = users
	}

	return response
}