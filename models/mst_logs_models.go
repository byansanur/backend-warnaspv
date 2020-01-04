package models

import (
	"../structs"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"time"
)

func GetRiwayat(id_mst_outlet string, id_mst_branch []int, id_cms_users string, created_at string, limit string, offset string) structs.JsonResponse {

	var (
		users []structs.Riwayat
		t     structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()
	first_name := ` (case when id_modul = 2 then
	(select first_name from lead where lead.id = id_data limit 1)
		when id_modul = 3 then (select first_name from target where target.id = id_data limit 1)
			when id_modul = 4 then (select first_name from contact where contact.id = id_data limit 1)
			when id_modul = 5 then (select first_name from contact where (select id_contact from "order"  orders where orders.id = id_data limit 1) = contact.id)
			when id_modul = 6 then (select first_name from target where target.id = id_data)
			when id_modul = 7 then (select first_name from lead where lead.id = id_data)
			ELSE 'Opps there is something wrong' end )as first_name,`

	last_name := ` (case when id_modul = 2 then
             (select last_name from lead where lead.id = id_data limit 1)
             when id_modul = 3 then (select last_name from target where target.id = id_data limit 1)
             when id_modul = 4 then (select last_name from contact where contact.id = id_data limit 1)
             when id_modul = 5 then (select last_name from contact where (select id_contact from "order" orders where orders.id = id_data limit 1) = contact.id)
             when id_modul = 6 then (select last_name from target where target.id = id_data)
             when id_modul = 7 then (select last_name from lead where lead.id = id_data)
             ELSE 'Opps there is something wrong' end)as last_name,`

	modul := `(case when id_modul = 1 and jenis = 'create'
             then 'Menambahkan Aktivitas'
             when id_modul = 1 and jenis = 'call' then 'Menghubungi Data Aktivitas'
             when id_modul = 1 and jenis = 'delete' then 'Menghapus Data Aktivitas'
             when id_modul = 2 and jenis = 'create' then 'Menambahkan Lead'
             when id_modul = 2 and jenis = 'call' then 'Menghubungi Data Lead'
             when id_modul = 2 and jenis = 'delete' then 'Menghapus Data Lead'
             when id_modul = 3 and jenis = 'create' then 'Menghubungi Target'
             when id_modul = 3 and jenis = 'call' then 'Menghubungi Data Target'
             when id_modul = 3 and jenis = 'delete' then 'Menghapus Data Target'
             when id_modul = 4 and jenis = 'create' then 'Menambahkan Contact'
             when id_modul = 4 and jenis = 'call' then 'Menghubungi Data Contact'
             when id_modul = 4 and jenis = 'delete' then 'Menghapus Data Contact'
             when id_modul = 5 and jenis = 'create' then 'Menambahkan Order'
             when id_modul = 6 and jenis = 'visit' then 'Visum Target'
             when id_modul = 7 and jenis = 'visit' then 'Visum Lead'
             ELSE 'Opps there is something wrong' end )as modul,`

	description := ` (case when id_modul = 1 and jenis = 'create' then 'Menambahkan Aktivitas'
             when id_modul = 1 and jenis = 'call' then 'Menghubungi Data Aktivitas'
             when id_modul = 1 and jenis = 'delete' then 'Menghapus Data Aktivitas'
             when id_modul = 2 and jenis = 'create' then 'Menambahkan Lead'
             when id_modul = 2 and jenis = 'call' then 'Menghubungi Data Lead'
             when id_modul = 2 and jenis = 'delete' then 'Menghapus Data Lead'
             when id_modul = 3 and jenis = 'create' then 'Menghubungi Target'
             when id_modul = 3 and jenis = 'call' then 'Menghubungi Data Target'
             when id_modul = 3 and jenis = 'delete' then 'Menghapus Data Target'
             when id_modul = 4 and jenis = 'create' then 'Menambahkan Contact'
             when id_modul = 4 and jenis = 'call' then 'Menghubungi Data Contact'
             when id_modul = 4 and jenis = 'delete' then 'Menghapus Data Contact'
             when id_modul = 5 and jenis = 'create' then 'Menambahkan Order'
             when id_modul = 6 and jenis = 'visit' then 'Visum Target'
             when id_modul = 7 and jenis = 'visit' then 'Visum Lead'
             ELSE 'Opps there is something wrong' end )as description`

	err := idb.DB.Table("mst_logs").Select(
		"mst_logs.id,mst_logs.created_at, cms_users.npm, cms_users.name, cms_users.status," +
			"cms_users.id_cms_privileges, cms_privileges.name as privileges_name, cms_users.id_mst_outlet," +
			"mst_outlet.outlet_name, mst_outlet.id_mst_branch, mst_branch.branch_name,mst_logs.id_cms_users," +
			first_name + last_name + modul + description)

	err = err.Joins("join  cms_users on cms_users.id = mst_logs.id_cms_users")
	err = err.Joins("join  cms_privileges on cms_privileges.id = cms_users.id_cms_privileges")
	err = err.Joins("join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("join mst_branch on mst_branch.id = mst_outlet.id_mst_branch")

	if limit != "" {
		err = err.Limit(limit)
	}
	if id_cms_users != "" {
		err = err.Where("mst_logs.id_cms_users = ?", id_cms_users)
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if created_at != "" {
		err = err.Where("mst_logs.created_at::text LIKE ?", "%"+created_at+"%")
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

func DownloadRiwayat(id_mst_outlet string, id_mst_branch []int, created_at string, limit string, offset string) structs.JsonResponse {

	var (
		riwayat []structs.Riwayat
		t       structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()
	first_name := ` (case when id_modul = 2 then
	(select first_name from lead where lead.id = id_data )
		when id_modul = 3 then (select first_name from target where target.id = id_data )
			when id_modul = 4 then (select first_name from contact where contact.id = id_data )
			when id_modul = 5 then (select first_name from contact where (select id_contact from "order"  orders where orders.id = id_data ) = contact.id)
			when id_modul = 6 then (select first_name from target where target.id = id_data)
			when id_modul = 7 then (select first_name from lead where lead.id = id_data)
			ELSE 'Opps there is something wrong' end )as first_name,`

	last_name := ` (case when id_modul = 2 then
             (select last_name from lead where lead.id = id_data )
             when id_modul = 3 then (select last_name from target where target.id = id_data )
             when id_modul = 4 then (select last_name from contact where contact.id = id_data )
             when id_modul = 5 then (select last_name from contact where (select id_contact from "order" orders where orders.id = id_data ) = contact.id)
             when id_modul = 6 then (select last_name from target where target.id = id_data)
             when id_modul = 7 then (select last_name from lead where lead.id = id_data)
             ELSE 'Opps there is something wrong' end)as last_name,`

	modul := `(case when id_modul = 1 and jenis = 'create'
             then 'Menambahkan Aktivitas'
             when id_modul = 1 and jenis = 'call' then 'Menghubungi Data Aktivitas'
             when id_modul = 1 and jenis = 'delete' then 'Menghapus Data Aktivitas'
             when id_modul = 2 and jenis = 'create' then 'Menambahkan Lead'
             when id_modul = 2 and jenis = 'call' then 'Menghubungi Data Lead'
             when id_modul = 2 and jenis = 'delete' then 'Menghapus Data Lead'
             when id_modul = 3 and jenis = 'create' then 'Menghubungi Target'
             when id_modul = 3 and jenis = 'call' then 'Menghubungi Data Target'
             when id_modul = 3 and jenis = 'delete' then 'Menghapus Data Target'
             when id_modul = 4 and jenis = 'create' then 'Menambahkan Contact'
             when id_modul = 4 and jenis = 'call' then 'Menghubungi Data Contact'
             when id_modul = 4 and jenis = 'delete' then 'Menghapus Data Contact'
             when id_modul = 5 and jenis = 'create' then 'Menambahkan Order'
             when id_modul = 6 and jenis = 'visit' then 'Visum Target'
             when id_modul = 7 and jenis = 'visit' then 'Visum Lead'
             ELSE 'Opps there is something wrong' end )as modul,`

	description := ` (case when id_modul = 1 and jenis = 'create' then 'Menambahkan Aktivitas'
             when id_modul = 1 and jenis = 'call' then 'Menghubungi Data Aktivitas'
             when id_modul = 1 and jenis = 'delete' then 'Menghapus Data Aktivitas'
             when id_modul = 2 and jenis = 'create' then 'Menambahkan Lead'
             when id_modul = 2 and jenis = 'call' then 'Menghubungi Data Lead'
             when id_modul = 2 and jenis = 'delete' then 'Menghapus Data Lead'
             when id_modul = 3 and jenis = 'create' then 'Menghubungi Target'
             when id_modul = 3 and jenis = 'call' then 'Menghubungi Data Target'
             when id_modul = 3 and jenis = 'delete' then 'Menghapus Data Target'
             when id_modul = 4 and jenis = 'create' then 'Menambahkan Contact'
             when id_modul = 4 and jenis = 'call' then 'Menghubungi Data Contact'
             when id_modul = 4 and jenis = 'delete' then 'Menghapus Data Contact'
             when id_modul = 5 and jenis = 'create' then 'Menambahkan Order'
             when id_modul = 6 and jenis = 'visit' then 'Visum Target'
             when id_modul = 7 and jenis = 'visit' then 'Visum Lead'
             ELSE 'Opps there is something wrong' end )as description,`

	err := idb.DB.Table("mst_logs").Select(
		"mst_logs.id, mst_logs.created_at,mst_logs.id_data, cms_users.npm, cms_users.name, cms_users.status," +
			"cms_users.id_cms_privileges, cms_privileges.name as privileges_name, cms_users.id_mst_outlet," +
			"mst_outlet.outlet_name, mst_outlet.id_mst_branch, mst_branch.branch_name,mst_logs.id_cms_users," +
			first_name + last_name + modul + description +
			"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +
			"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh_name," +
			"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
			"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv_name")

	err = err.Joins("left join  cms_users on cms_users.id = mst_logs.id_cms_users")
	err = err.Joins("left join  cms_privileges on cms_privileges.id = cms_users.id_cms_privileges")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id = mst_outlet.id_mst_branch")
	err = err.Joins("left join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")

	if limit != "" {
		err = err.Limit(limit)
	}

	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if created_at != "" {
		err = err.Where("mst_logs.created_at::text LIKE ?", "%"+created_at+"%")
	}

	err = err.Find(&riwayat)
	errx := err.Error

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "Id data")
	f.SetCellValue("Sheet1", "B1", "Created_at")
	f.SetCellValue("Sheet1", "C1", "NPM")
	f.SetCellValue("Sheet1", "D1", "Name")
	f.SetCellValue("Sheet1", "E1", "Privileges Name")
	f.SetCellValue("Sheet1", "F1", "Outlet Name")
	f.SetCellValue("Sheet1", "G1", "Branch Name")
	f.SetCellValue("Sheet1", "H1", "FirstName")
	f.SetCellValue("Sheet1", "I1", "LastName")
	f.SetCellValue("Sheet1", "J1", "Modul")
	f.SetCellValue("Sheet1", "K1", "Description")
	f.SetCellValue("Sheet1", "L1", "Oh name")
	f.SetCellValue("Sheet1", "M1", "Spv name")

	for i, _ := range riwayat {

		rows := 2

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+rows), riwayat[i].IdData)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+rows), riwayat[i].CreatedAt)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+rows), riwayat[i].Npm)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+rows), riwayat[i].Name)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+rows), riwayat[i].PrivilegesName)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+rows), riwayat[i].OutletName)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+rows), riwayat[i].BranchName)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+rows), riwayat[i].FirstName)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+rows), riwayat[i].LastName)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+rows), riwayat[i].Modul)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+rows), riwayat[i].Description)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+rows), riwayat[i].OhName)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(i+rows), riwayat[i].SpvName)

	}

	f.SetActiveSheet(index)
	folderPath := "./files/"

	filename := "Riwayat" + time.Now().Format("20060102150405") + ".xlsx"

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

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = filename
	}

	return response
}
