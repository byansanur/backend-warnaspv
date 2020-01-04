package structs

type GetOrderMstStatus struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type GetOrderMstReason struct {
	Id               int    `json:"id"`
	Reason           string `json:"reason"`
	IdOrderMstStatus int    `json:"id_order_mst_status"`
}

type GetMstOutlet struct {
	Id            int    `json:"id"`
	OutletName    string `json:"outlet_name"`
	IdMstBranch   int    `json:"id_mst_branch"`
	BranchName    string `json:"branch_name"`
	OutletId      string `json:"outlet_id"`
	OutletSysId   string `json:"outlet_sys_id"`
	BizType       string `json:"biz_type"`
	OutletAddress string `json:"outlet_address"`
	OutletFifCode string `json:"outlet_fif_code"`
	OutletDesc    string `json:"outlet_desc"`
	OutletStatus  string `json:"outlet_status"`
}

type GetMstOutletDetail struct {
	Id             int    `json:"id"`
	OutletName     string `json:"outlet_name"`
	IdMstBranch    int    `json:"id_mst_branch"`
	BranchName     string `json:"branch_name"`
	OutletId       string `json:"outlet_id"`
	OutletSysId    string `json:"outlet_sys_id"`
	IdMstBizType   int    `json:"id_mst_biz_type"`
	BizType        string `json:"biz_type"`
	OutletAddress  string `json:"outlet_address"`
	OutletFifCode  string `json:"outlet_fif_code"`
	OutletDesc     string `json:"outlet_desc"`
	OutletLocation string `json:"outlet_location"`
	OutletStatus   string `json:"outlet_status"`
}

type UpdateMstOutlet struct {
	Id             int    `json:"id"`
	OutletName     string `json:"outlet_name"`
	IdMstBranch    int    `json:"id_mst_branch"`
	OutletId       string `json:"outlet_id"`
	OutletSysId    string `json:"outlet_sys_id"`
	IdMstBizType   int    `json:"id_mst_biz_type"`
	OutletLocation string `json:"outlet_location"`
	OutletAddress  string `json:"outlet_address"`
	IdMstAddress   int    `json:"id_mst_address"`
	OutletFifCode  string `json:"outlet_fif_code"`
	OutletDesc     string `json:"outlet_desc"`
	OutletStatus   string `json:"outlet_status"`
	OutletLat      string `json:"outlet_lat"`
	OutletLng      string `json:"outlet_lng"`
	IdCmsUsers     int    `json:"id_cms_users"`
	UpdatedBy      int    `json:"updated_by"`
	UpdatedAt      string `json:"updated_at"`
}

type CreateMstOutlet struct {
	Id             int    `json:"id"`
	OutletName     string `json:"outlet_name"`
	IdMstBranch    int    `json:"id_mst_branch"`
	OutletId       string `json:"outlet_id"`
	OutletSysId    string `json:"outlet_sys_id"`
	IdMstBizType   int    `json:"id_mst_biz_type"`
	OutletLocation string `json:"outlet_location"`
	OutletAddress  string `json:"outlet_address"`
	IdMstAddress   int    `json:"id_mst_address"`
	OutletFifCode  string `json:"outlet_fif_code"`
	OutletDesc     string `json:"outlet_desc"`
	OutletStatus   string `json:"outlet_status"`
	OutletLat      string `json:"outlet_lat"`
	OutletLng      string `json:"outlet_lng"`
	IdCmsUsers     int    `json:"id_cms_users"`
}

type GetMstBranch struct {
	Id         int    `json:"id"`
	BranchName string `json:"branch_name"`
}

type GetPrivileges struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetTargetMstStatus struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type GetMstLogDesc struct {
	Id             int    `json:"id"`
	Description    string `json:"description"`
	IdMstLogStatus int    `json:"id_mst_log_status"`
}
type VueMenusPrivileges struct {
	Id              int    `json:"id"`
	IdVueMenus      int    `json:"id_vue_menus"`
	IdCmsPrivileges int    `json:"id_cms_privileges"`
	CreatedAt       string `json:"created_at"`
	Title           string `json:"title"`
	Name            string `json:"name"`
	Group           string `json:"group"`
	Icon            string `json:"icon"`
	Componet        string `json:"componet"`
}

type MstUnits struct {
	Id         int    `json:"id"`
	BranchName string `json:"branch_name"`
	Year       int    `json:"year"`
	KodeUnit   string `json:"kode_unit"`
	Merk       string `json:"merk"`
	Type       string `json:"type"`
	Model      string `json:"model"`
	Otr        int    `json:"otr"`
}
type MstUnit struct {
	Id          int    `json:"id"`
	IdMstBranch int    `json:"id_mst_branch"`
	BranchName  string `json:"branch_name"`
	Year        int    `json:"year"`
	KodeUnit    string `json:"kode_unit"`
	Merk        string `json:"merk"`
	Type        string `json:"type"`
	Model       string `json:"model"`
	Otr         int    `json:"otr"`
}

type UpdateMstUnits struct {
	Id         int    `json:"id"`
	BranchName string `json:"branch_name"`
	Year       int    `json:"year"`
	KodeUnit   string `json:"kode_unit"`
	Merk       string `json:"merk"`
	Type       string `json:"type"`
	Model      string `json:"model"`
	Otr        int    `json:"otr"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateMstUnits struct {
	Id          int    `json:"id"`
	IdMstBranch int    `json:"id_mst_branch"`
	IdCmsUsers  int    `json:"id_cms_users"`
	Year        int    `json:"year"`
	KodeUnit    string `json:"kode_unit"`
	Merk        string `json:"merk"`
	Type        string `json:"type"`
	Model       string `json:"model"`
	Otr         int    `json:"otr"`
}

type GetMstBizType struct {
	Id      int    `json:"id"`
	BizCode string `json:"biz_code"`
	BizType string `json:"biz_type"`
}
