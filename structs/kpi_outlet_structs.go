package structs

// query for kpi lead per outlet
// select count(lead.id) from lead where lead.id_mst_outlet = 1241 and lead.created_at::text like '%2020-01%'

//SELECT
//mo.id_mst_branch,
//mo.outlet_name,
//COUNT(od.id) as total_order
//FROM mst_outlet mo
//INNER JOIN "order" od ON mo.id = od.id_mst_outlet
//WHERE mo.id_mst_branch = 109
//GROUP BY mo.id_mst_branch, mo.outlet_name
//ORDER BY mo.outlet_name

type KpiOutletLead struct {
	Id           *int64 `json:"id" form:"id" binding:"-"`
	IdMstBranch  *int64 `json:"id_mst_branch"  form:"id_mst_branch" binding:"-"`
	OutletName   string `json:"outlet_name"`
	TotalOrder   int64  `json:"total_order"`
	TotalLead    int64  `json:"total_lead"`
	OutletStatus string `json:"outlet_status"`
}
