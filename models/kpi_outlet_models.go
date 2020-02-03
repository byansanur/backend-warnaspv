package models

import (
	"../structs"
	"fmt"
)

func GetKpiOutlet(kpiLead structs.KpiOutletLead, limit string, offset string) ([]structs.KpiOutletLead, error) {

	fmt.Print(kpiLead)
	data := []structs.KpiOutletLead{}
	get := idb.DB.Table("mst_outlet").
		Select(`mst_outlet.id, mst_outlet.id_mst_branch, mst_outlet.outlet_name, mst_outlet.outlet_status,
			(select count(id) as total_order from "order" orders where orders.id_mst_outlet = mst_outlet.id), 
			(select count(id) as total_lead from lead where lead.id_mst_outlet = mst_outlet.id)`).
		Where(`mst_outlet.outlet_status = 'Y'`)

	fmt.Println("idbranch ", kpiLead.IdMstBranch)
	if kpiLead.Id != nil {
		get = get.Where(`mst_outlet.id = ?`, kpiLead.Id)
	}
	if kpiLead.IdMstBranch != nil {
		get = get.Where(`mst_outlet.id_mst_branch in (?)`, kpiLead.IdMstBranch)
	}
	if limit != "" {
		get = get.Limit(limit)
	}
	if offset != "" {
		get = get.Offset(offset)
	}

	err := get.Find(&data).Error
	fmt.Println("get orders", data)

	return data, err

}
