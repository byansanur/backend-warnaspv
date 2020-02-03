package models

import (
	"../structs"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"time"
)

func TelemarketingUsers(id_mst_outlet string, id_mst_branch []int, created_at string, limit string, offset string) structs.JsonResponse {

	var (
		tele []structs.TelemarketingUsers
		t    structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	//date := "%"+created_at+"%"

	err := idb.DB.Table("cms_users").Select(`
 cms_users.id, cms_users.npm, cms_users.name , mst_outlet.outlet_name,
		cms_users.status,cms_privileges.name as privileges_name, mst_branch.branch_name,
		(select count(target_log.id) from target_log where target_log.id_cms_users = cms_users.id and target_log.created_at::text LIKE '%` + created_at + `%' ) as total,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 1 and target_log.created_at::text LIKE '%` + created_at + `%')  as berminat,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 2  and target_log.created_at::text LIKE '%` + created_at + `%')  as tidak_berminat,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 3  and target_log.created_at::text LIKE '%` + created_at + `%')  as pikir_pikir,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 4  and target_log.created_at::text LIKE'%` + created_at + `%')  as tanya_pasangan,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 5  and target_log.created_at::text LIKE '%` + created_at + `%')  as butuh_waktu,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 6 and target_log.created_at::text LIKE '%` + created_at + `%')  as reject,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 7  and target_log.created_at::text LIKE'%` + created_at + `%')  as tidak_diangkat,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 8  and target_log.created_at::text LIKE '%` + created_at + `%')  as no_tidak_terdaftar,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 9  and target_log.created_at::text LIKE '%` + created_at + `%')  as no_tidak_aktif,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 10  and target_log.created_at::text LIKE '%` + created_at + `%') as cancel,
       (select count(target_log.id) from target_log
           join mst_log_desc on mst_log_desc.id = target_log.id_mst_log_desc
       where target_log.id_cms_users = cms_users.id and mst_log_desc.id_mst_log_status = 11  and target_log.created_at::text LIKE '%` + created_at + `%') as salah_sambung
`)

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}
	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}

	err = err.Find(&tele)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	}

	if len(tele) <= 0 {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = tele

	}

	return response
}

func TelemarketingDownload(id_mst_outlet string, id_mst_branch []int, created_at string, limit string, offset string) structs.JsonResponse {

	var (
		tele []structs.TeleDownload
		t    structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("target_log").Select("target_log.id, target_log.created_at," +
		"target_log.id_target," +
		"target_log.duration,mst_log_desc.description,target_log.recall,cms_users.name,cms_users.npm," +
		"cms_privileges.name as privileges_name,target.address as alamat, target.last_name as lastname," +
		"mst_outlet.outlet_name as users_outlet_name,mst_branch.branch_name as users_branch_name,target_log.id_target,target.created_at as date_upload," +
		"target.id_mst_branch,mst_data_source.datasource,target.priority,target.nopol,target.first_name," +
		"target.no_contract,target.provider_1 as provider_satu,target.provider_2 as provider_dua,target.kabupaten,target.kecamatan," +
		"target.kelurahan,target.provinsi,target.job,mst_log_desc.description,target_mst_status.status," +
		"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +

		"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh," +

		"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
		"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv")

	err = err.Joins("left join target on target_log.id_target = target.id")
	err = err.Joins("left join mst_log_desc on target_log.id_mst_log_desc = mst_log_desc.id")
	err = err.Joins("left join cms_users on cms_users.id = target_log.id_cms_users")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id = mst_outlet.id_mst_branch")
	err = err.Joins("left join mst_data_source on mst_data_source.id = target.id_mst_data_source")
	err = err.Joins("left join target_mst_status on target.id_target_mst_status = target_mst_status.id")
	err = err.Joins("left join cms_privileges on cms_privileges.id = cms_users.id_cms_privileges")
	err = err.Joins("left join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")

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
	if created_at != "" {
		err = err.Where("target_log.created_at::text LIKE  ?", "%"+created_at+"%")
	}

	err = err.Find(&tele)
	errx := err.Error

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "ID Target")
	f.SetCellValue("Sheet1", "B1", "Created at")
	f.SetCellValue("Sheet1", "C1", "duration")
	f.SetCellValue("Sheet1", "D1", "description")
	f.SetCellValue("Sheet1", "E1", "recall")
	f.SetCellValue("Sheet1", "F1", "Tanggal Upload")
	f.SetCellValue("Sheet1", "G1", "datasource")
	f.SetCellValue("Sheet1", "H1", "priority")
	f.SetCellValue("Sheet1", "I1", "nopol")
	f.SetCellValue("Sheet1", "J1", "no_contract")
	f.SetCellValue("Sheet1", "K1", "provider_1")
	f.SetCellValue("Sheet1", "L1", "provider_2")
	f.SetCellValue("Sheet1", "M1", "status")
	f.SetCellValue("Sheet1", "N1", "first_name")
	f.SetCellValue("Sheet1", "O1", "last_name")
	f.SetCellValue("Sheet1", "P1", "alamat")
	f.SetCellValue("Sheet1", "Q1", "kelurahan")
	f.SetCellValue("Sheet1", "R1", "kecamatan")
	f.SetCellValue("Sheet1", "S1", "kabupaten")
	f.SetCellValue("Sheet1", "T1", "provinsi")
	f.SetCellValue("Sheet1", "U1", "name")
	f.SetCellValue("Sheet1", "V1", "NPM")
	f.SetCellValue("Sheet1", "W1", "privileges_name")
	f.SetCellValue("Sheet1", "X1", "users_outlet_name")
	f.SetCellValue("Sheet1", "Y1", "users_branch_name")
	f.SetCellValue("Sheet1", "Z1", "Outlet Head")
	f.SetCellValue("Sheet1", "AA1", "SPV Head")
	//f.SetCellValue("Sheet1", "R1", "oh_name")
	//f.SetCellValue("Sheet1", "S1", "spv_name")

	for i, _ := range tele {

		rows := 2

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+rows), tele[i].IdTarget)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+rows), tele[i].CreatedAt)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+rows), tele[i].Duration)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+rows), tele[i].Description)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+rows), tele[i].Recall)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+rows), tele[i].DateUpload)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+rows), tele[i].Datasource)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+rows), tele[i].Priority)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+rows), tele[i].Nopol)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+rows), tele[i].NoContract)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+rows), tele[i].ProviderSatu)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+rows), tele[i].ProviderDua)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(i+rows), tele[i].Status)
		f.SetCellValue("Sheet1", "N"+strconv.Itoa(i+rows), tele[i].FirstName)
		f.SetCellValue("Sheet1", "O"+strconv.Itoa(i+rows), tele[i].Lastname)
		f.SetCellValue("Sheet1", "P"+strconv.Itoa(i+rows), tele[i].Alamat)
		f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i+rows), tele[i].Kelurahan)
		f.SetCellValue("Sheet1", "R"+strconv.Itoa(i+rows), tele[i].Kecamatan)
		f.SetCellValue("Sheet1", "S"+strconv.Itoa(i+rows), tele[i].Kabupaten)
		f.SetCellValue("Sheet1", "T"+strconv.Itoa(i+rows), tele[i].Provinsi)
		f.SetCellValue("Sheet1", "U"+strconv.Itoa(i+rows), tele[i].Name)
		f.SetCellValue("Sheet1", "V"+strconv.Itoa(i+rows), tele[i].Npm)
		f.SetCellValue("Sheet1", "W"+strconv.Itoa(i+rows), tele[i].PrivilegesName)
		f.SetCellValue("Sheet1", "X"+strconv.Itoa(i+rows), tele[i].UsersOutletName)
		f.SetCellValue("Sheet1", "Y"+strconv.Itoa(i+rows), tele[i].UsersBranchName)
		f.SetCellValue("Sheet1", "Z"+strconv.Itoa(i+rows), tele[i].Oh)
		f.SetCellValue("Sheet1", "AA"+strconv.Itoa(i+rows), tele[i].Spv)

	}

	f.SetActiveSheet(index)
	folderPath := "./files/"

	filename := "tele" + time.Now().Format("20060102150405") + ".xlsx"

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

	if len(tele) <= 0 {

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

func DownloadTargetAssignmentLog(id_mst_branch []int, created_at string, limit string, offset string) structs.JsonResponse {

	var (
		tele []structs.TargetAssignmentLog
		t    structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("target_assignment_log").Select(`cms_users.id,
target_assignment_log.created_at,cms_users.name, cms_users.npm,
                mst_outlet.outlet_name, mst_branch.branch_name, cms_privileges.name as privileges,cms_users.id_mst_outlet,mst_outlet.id_mst_branch,
                cms_users.id_cms_privileges,
               (select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh, 
		(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh_name,
		(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv,
		(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv_name,
                target_assignment_log.total,
                (SELECT name from cms_users WHERE cms_users.id = target_assignment_log.target_users) as target_name,
                (SELECT npm from cms_users WHERE cms_users.id = target_assignment_log.target_users) as target_npm,
                (SELECT cms_privileges.name from cms_privileges join cms_users on cms_users.id_cms_privileges = cms_privileges.id where cms_users.id = target_assignment_log.target_users) as target_privileges
`)

	err = err.Joins("join cms_users on cms_users.id = target_assignment_log.id_cms_users")
	err = err.Joins("join cms_privileges on cms_privileges.id = cms_users.id_cms_privileges")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id = mst_outlet.id_mst_branch")
	err = err.Joins("join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if len(id_mst_branch) >= 1 {
		fmt.Println("asdasdasdasdasd")
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if created_at != "" {
		err = err.Where("target_assignment_log.created_at::text LIKE  ?", "%"+created_at+"%")
	}

	err = err.Find(&tele)
	errx := err.Error

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "Tanggal")
	f.SetCellValue("Sheet1", "B1", "Total")
	f.SetCellValue("Sheet1", "C1", "Npm")
	f.SetCellValue("Sheet1", "D1", "Name")
	f.SetCellValue("Sheet1", "E1", "Privileges")
	f.SetCellValue("Sheet1", "F1", "OutletName")
	f.SetCellValue("Sheet1", "G1", "BranchNAme")
	f.SetCellValue("Sheet1", "H1", "Spv Name")
	f.SetCellValue("Sheet1", "I1", "Oh Name")
	f.SetCellValue("Sheet1", "J1", "Target Name")
	f.SetCellValue("Sheet1", "K1", "Target NPM")
	f.SetCellValue("Sheet1", "L1", "Target Privileges")

	for i, _ := range tele {

		rows := 2

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+rows), tele[i].CreatedAt)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+rows), tele[i].Total)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+rows), tele[i].Npm)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+rows), tele[i].Name)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+rows), tele[i].PrivilegesName)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+rows), tele[i].OutletName)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+rows), tele[i].BranchName)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+rows), tele[i].SpvName)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+rows), tele[i].OhName)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+rows), tele[i].TargetName)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+rows), tele[i].TargetNpm)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+rows), tele[i].TargetPrivileges)

	}

	f.SetActiveSheet(index)
	folderPath := "./files/"

	filename := "assignmentLog" + time.Now().Format("20060102150405") + ".xlsx"

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

	if len(tele) <= 0 {

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

func GetDataAssignmnet(created_at string, id_mst_branch string, id_cms_users string, id_target_mst_status string, category string,
	no_contract string, provider_1 string, provider_2 string, kelurahan string, kecamatan string, kabupaten string,
	provinsi string, id_mst_log_desc string, limit string, offset string) structs.JsonResponse {

	var (
		assignment []structs.GetDataAssignment
		t          structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("target").Select("target.id ,to_char(target.created_at, 'YYYY-MM-DD HH24:MI') as created_at ,target.id_mst_branch," +
		"target.first_name ,target.last_name,target.no_contract,target.kabupaten," +
		"target.kecamatan,target.kelurahan,target.provider_1 as provider_satu, target.provider_2 as provider_dua," +
		"target.provinsi,cms_users.name,mst_log_desc.description,target_mst_status.status")

	err = err.Joins("left join target_log on target_log.id_target = target.id and target_log.id = (select max(a.id) from target_log a where a.id_target = target_log.id_target)")
	err = err.Joins("join cms_users on target.id_cms_users = cms_users.id")
	err = err.Joins("left join mst_log_desc on target_log.id_mst_log_desc = mst_log_desc.id")
	err = err.Joins("join target_mst_status on target.id_target_mst_status = target_mst_status.id")
	//err = err.Joins("left join target_visum on target_visum.id_target = target.id and target_visum.id = (select max (b.id) from target_visum b where b.id_target = target_visum.id_target)")

	//err = err.Where("target_visum.id is null")
	err = err.Where("target.id_target_mst_status != 6")

	err = err.Order("target.id desc")

	id_mst_branch_conv, _ := strconv.Atoi(id_mst_branch)

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if created_at != "" {
		err = err.Where("target.created_at::text LIKE ?", "%"+created_at+"%")
	}
	if id_cms_users != "" {
		err = err.Where("target.id_cms_users = ?", id_cms_users)
	}
	if id_mst_branch != "" {
		err = err.Where("target.id_mst_branch = ?", id_mst_branch_conv)
	}
	if id_target_mst_status != "" {
		err = err.Where("target.id_target_mst_status = ?", id_target_mst_status)
	}
	if id_mst_log_desc != "" {
		err = err.Where("target_log.id_mst_log_desc = ?", id_mst_log_desc)
	}
	if category != "" {
		err = err.Where("target.category::text ILIKE ?", "%"+category+"%")
	}
	if no_contract != "" {
		err = err.Where("target.no_contract::text ILIKE  ?", "%"+no_contract+"%")
	}
	if provider_1 != "" {
		err = err.Where("target.provider_1::text ILIKE  ?", "%"+provider_1+"%")
	}
	if provider_2 != "" {
		err = err.Where("target.provider_2::text ILIKE  ?", "%"+provider_2+"%")
	}
	if kelurahan != "" {
		err = err.Where("target.kelurahan::text ILIKE  ?", "%"+kelurahan+"%")
	}
	if kecamatan != "" {
		err = err.Where("target.kecamatan::text ILIKE  ?", "%"+kecamatan+"%")
	}
	if kabupaten != "" {
		err = err.Where("target.kabupaten::text ILIKE  ?", "%"+kabupaten+"%")
	}
	if provinsi != "" {
		err = err.Where("target.provinsi::text ILIKE  ?", "%"+provinsi+"%")
	}

	err = err.Find(&assignment)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = assignment
	}

	return response
}

func UpdateAssignment(id []int, id_target_mst_status string,
	id_cms_users string, updated_by string) structs.JsonResponse {

	var (
		updateAssignment structs.UpdateAssignment
		t                structs.Component
	)

	response := structs.JsonResponse{}

	//id_conv, _ := strconv.Atoi(id)
	id_cms_users_conv, _ := strconv.Atoi(id_cms_users)
	id_target_mst_status_conv, _ := strconv.Atoi(id_target_mst_status)
	updated_by_conv, _ := strconv.Atoi(updated_by)

	var id_int = []int{}
	for _, i := range id {
		//j, err := strconv.ParseInt(i, 10, 64)
		//if err != nil {
		//	//panic(err)
		//	fmt.Println(err)
		//}
		id_int = append(id_int, i)
	}
	//updateAssignment.Id = id_int
	updateAssignment.IdCmsUsers = id_cms_users_conv
	updateAssignment.IdTargetMstStatus = id_target_mst_status_conv
	updateAssignment.UpdatedBy = updated_by_conv
	updateAssignment.UpdatedAt = t.GetTimeNow()

	err := idb.DB.Table("target").Where("id IN  (?)", id_int).Updates(updateAssignment).Error

	fmt.Println("err", err)
	if err != nil {

		response.ApiStatus = 0
		response.ApiMessage = err.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = "success"
		response.Data = updateAssignment
	}

	return response
}

func TargetLogAssignment(id_cms_users string, target_users string, total string, id_target string) structs.JsonResponse {

	var (
		ass structs.CreateTargetAssignmentLog
	)

	response := structs.JsonResponse{}

	id_cms_users_conv, _ := strconv.Atoi(id_cms_users)
	target_users_conv, _ := strconv.Atoi(target_users)
	total_conv, _ := strconv.Atoi(total)
	//id_target_conv,_ := strconv.Atoi(id_target)
	//updated_at_conv,_ := strconv.Atoi(updated_at)

	ass.IdCmsUsers = id_cms_users_conv
	ass.TargetUsers = target_users_conv
	ass.Total = total_conv
	ass.IdTarget = id_target

	err := idb.DB.Table("target_assignment_log").Create(&ass).Error

	fmt.Println("err", err)
	if err != nil {

		response.ApiStatus = 0
		response.ApiMessage = err.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = "success"
		response.Data = ass
	}

	return response
}

//func TargetStatus(id_mst_branch string) structs.JsonResponse {
//
//	var (
//		targetvisumdetail structs.GetTargetStatus
//		t                 structs.Component
//	)
//
//	response := structs.JsonResponse{}
//	//tables := t.TblLead()
//
//	err := idb.DB.Table("target").Select(`count(target.id) as target,
//       (select count(target.id) from  target where target.id_target_mst_status = 1) as target_new,
//       (select count(target.id) from  target where target.id_target_mst_status = 2) as target_how,
//       (select count(target.id) from  target where target.id_target_mst_status = 3) as target_working,
//       (select count(target.id) from  target where target.id_target_mst_status = 4) as target_visit,
//       (select count(target.id) from  target where target.id_target_mst_status = 5) as target_unqualified,
//       (select count(target.id) from  target where target.id_target_mst_status = 6) as target_converted`)
//
//	if id_mst_branch != "" {
//		err = err.Where("target.id_mst_branch = ?", id_mst_branch)
//	}
//
//
//	err = err.Find(&targetvisumdetail)
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
//		response.Data = targetvisumdetail
//	}
//
//	return response
//}

func TargetStatus(id_mst_branch []int, created_at string) structs.JsonResponse {

	var (
		status structs.GetTargetStatus
		t      structs.Component
	)

	response := structs.JsonResponse{}

	//target_new
	target_new := idb.DB.Table("target").Select("count(target.id) as target_new")
	target_new = target_new.Where("target.id_target_mst_status = 1")
	if created_at != "" {
		target_new = target_new.Where("target.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		target_new = target_new.Where("target.id_mst_branch IN (?)", id_mst_branch)
	}
	target_new = target_new.Find(&status)
	errtarget_new := target_new.Error

	//target_hot
	target_hot := idb.DB.Table("target").Select("count(target.id) as target_hot")
	target_hot = target_hot.Where("target.id_target_mst_status = 2")
	if created_at != "" {
		target_hot = target_hot.Where("target.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		target_hot = target_hot.Where("target.id_mst_branch IN (?)", id_mst_branch)
	}
	target_hot = target_hot.Find(&status)
	errtarget_hot := target_hot.Error

	//target_working
	target_working := idb.DB.Table("target").Select("count(target.id) as target_working")
	target_working = target_working.Where("target.id_target_mst_status = 3")
	if created_at != "" {
		target_working = target_working.Where("target.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		target_working = target_working.Where("target.id_mst_branch IN (?)", id_mst_branch)
	}
	target_working = target_working.Find(&status)
	errtarget_working := target_working.Error

	// target visit
	target_visit := idb.DB.Table("target").Select("count(target.id) as target_visit")
	target_working = target_working.Where("target.id_target_mst_status = 4")
	if created_at != "" {
		target_visit = target_visit.Where("target.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		target_visit = target_visit.Where("target.id_mst_branch IN (?)", id_mst_branch)
	}
	target_visit = target_visit.Find(&status)
	errtarget_visit := target_visit.Error

	// target unqualified
	target_unqualified := idb.DB.Table("target").Select("count(target.id) as target_unqualified")
	target_unqualified = target_unqualified.Where("target.id_target_mst_status = 5")
	if created_at != "" {
		target_unqualified = target_unqualified.Where("target.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		target_unqualified = target_unqualified.Where("target.id_mst_branch IN (?)", id_mst_branch)
	}
	target_unqualified = target_unqualified.Find(&status)
	errtarget_unqualified := target_unqualified.Error

	// target converted
	target_converted := idb.DB.Table("target").Select("count(target.id) as target_converted")
	target_converted = target_converted.Where("target.id_target_mst_status = 6")
	if created_at != "" {
		target_converted = target_converted.Where("target.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		target_converted = target_converted.Where("target.id_mst_branch IN (?)", id_mst_branch)
	}
	target_converted = target_converted.Find(&status)
	errtarget_converted := target_converted.Error

	if errtarget_new != nil && errtarget_hot != nil && errtarget_working != nil && errtarget_visit != nil && errtarget_unqualified != nil && errtarget_converted != nil {

		response.ApiStatus = 0
		response.ApiMessage = errtarget_new.Error()
		response.Data = nil

	} else {

		convStruct := [6]structs.Conv{}

		convStruct[0].Name = "New"
		convStruct[0].Value = status.TargetNew
		convStruct[1].Name = "Hot"
		convStruct[1].Value = status.TargetHot
		convStruct[2].Name = "Working"
		convStruct[2].Value = status.TargetWorking
		convStruct[3].Name = "Visit"
		convStruct[3].Value = status.TargetVisit
		convStruct[4].Name = "Unqualified"
		convStruct[4].Value = status.TargetUnqualifed
		convStruct[5].Name = "Converted"
		convStruct[5].Value = status.TargetConverted

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = convStruct
	}

	return response
}
