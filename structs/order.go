package structs

type GetOrderStatus struct {
	Id               int    `json:"id"`
	IdOrderMstStatus int    `json:"id_order_mst_status"`
	IdOrderMstReason string `gorm:"default:'nil'" json:"id_order_mst_reason"`
	Created_at       string `json:"created_at"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Merk             string `json:"merk"`
	Type             string `json:"type"`
	Year             int    `json:"year"`
	Plafond          int    `json:"plafond"`
	Installment      int    `json:"installment"`
	Tenor            int    `json:"tenor"`
	Status           string `json:"status"`
	Reason           string `json:"reason"`
}

type GetOrderDetail struct {
	Id               int    `json:"id"`
	IdOrderMstStatus int    `json:"id_order_mst_status"`
	IdOrderMstReason int    `json:"id_order_mst_reason"`
	Created_at       string `json:"created_at"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Name             string `json:"name"`
	//Status string `json:"status"`
	Datasource  string `json:"datasource"`
	Nopol       string `json:"nopol"`
	TaxStatus   string `json:"tax_status"`
	Owner       string `json:"owner"`
	Merk        string `json:"merk"`
	Type        string `json:"type"`
	Year        int    `json:"year"`
	Otr         int    `json:"otr"`
	Model       string `json:"model"`
	Plafond     int    `json:"plafond"`
	Installment int    `json:"installment"`
	DownPayment int    `json:"down_payment"`
	Tenor       int    `json:"tenor"`
	Need        string `json:"need"`
	OtrCustom   int    `json:"otr_custom"`
	SuretyName  string `json:"surety_name"`
	Status      string `json:"status"`
	Reason      string `json:"reason"`
}

type UpdateOrderStatus struct {
	Id               int `form:"id" json:"id" binding:"required"`
	IdOrderMstStatus int `json:"id_order_mst_status" form:"id_order_mst_status" binding:"required"`
	UpdatedBy        int `json:"updated_by" validate:"required"`
	IdOrderMstReason int `json:"id_order_mst_reason"`
}

type OrderUsers struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Npm            string `json:"npm"`
	OutletName     string `json:"outlet_name"`
	BranchName     string `json:"branch_name"`
	PrivilegesName string `json:"privileges_name"`
	Orders         int    `json:"orders"`
	Approve        int    `json:"approve"`
	Process        int    `json:"process"`
	Cancel         int    `json:"cancel"`
	Reject         int    `json:"reject"`
	Paid           int    `json:"paid"`
}

type OrderDownload struct {
	Id                     int    `json:"id"`
	CreatedAt              string `json:"created_at"`
	Category               string `json:"category"`
	StatusAddress          string `json:"status_address"`
	NoOrder                string `json:"no_order"`
	Survey                 string `json:"survey"`
	DownPayment            int    `json:"down_payment"`
	Installment            int    `json:"installment"`
	Need                   string `json:"need"`
	OtrCustom              int    `json:"otr_custom"`
	Plafond                int    `json:"plafond"`
	Tenor                  int    `json:"tenor"`
	Reason                 string `json:"reason"`
	OrderStatus            string `json:"order_status"`
	OrderNote              string `json:"order_note"`
	NomorTransaksi         string `json:"nomor_transaksi"`
	Nopol                  string `json:"nopol"`
	OtrTransaksi           string `json:"otr_transaksi"`
	Owner                  string `json:"owner"`
	TaxStatus              string `json:"tax_status"`
	FifBranch              string `json:"fif_branch"`
	CodePosFif             string `json:"code_pos_fif"`
	PosName                string `json:"pos_name"`
	KodeUnit               string `json:"kode_unit"`
	Merk                   string `json:"merk"`
	Type                   string `json:"type"`
	Model                  string `json:"model"`
	Year                   string `json:"year"`
	Otr                    string `json:"otr"`
	SuretyName             string `json:"surety_name"`
	SuretyBirthDate        string `json:"surety_birth_date"`
	SuretyBirthPlace       string `json:"surety_birth_place"`
	SuretyJob              string `json:"surety_job"`
	SuretyCompany          string `json:"surety_company"`
	SuretyPosition         string `json:"surety_position"`
	SuretyWorkingTime      int    `json:"surety_working_time"`
	SuretyIncome           int    `json:"surety_income"`
	SuretyOutlay           int    `json:"surety_outlay"`
	ContactFirstName       string `json:"contact_first_name"`
	ContactLastName        string `json:"contact_last_name"`
	ContactNik             string `json:"contact_nik"`
	ContactBirthDate       string `json:"contact_birth_date"`
	ContactBirthPlace      string `json:"contact_birth_place"`
	ContactGender          string `json:"contact_gender"`
	ContactReligion        string `json:"contact_religion"`
	ContactMaterial        string `json:"contact_material"`
	ContactJob             string `json:"contact_job"`
	ContactDatasource      string `json:"contact_datasource"`
	ContactCompany         string `json:"contact_company"`
	ContactMother          string `json:"contact_mother"`
	ContactFamily          int    `json:"contact_family"`
	ContactPlaceStatus     string `json:"contact_place_status"`
	ContactEmployeeStatus  string `json:"contact_employee_status"`
	ContactPosition        string `json:"contact_position"`
	ContactWorkingTime     int    `json:"contact_working_time"`
	ContactIncome          int    `json:"contact_income"`
	ContactOutlay          int    `json:"contact_outlay"`
	ContactNote            string `json:"contact_note"`
	ContactPhoneSatu       string `json:"contact_phone_satu"`
	ContactPhoneDua        string `json:"contact_phone_dua"`
	ContactAddress         string `json:"contact_address"`
	ContactRt              string `json:"contact_rt"`
	ContactRw              string `json:"contact_rw"`
	ContactKelurahan       string `json:"contact_kelurahan"`
	ContactKecamatan       string `json:"contact_kecamatan"`
	ContactKabupaten       string `json:"contact_kabupaten"`
	ContactProvinsi        string `json:"contact_provinsi"`
	ContactKodepos         string `json:"contact_kodepos"`
	ContactAddressCategory string `json:"contact_address_category"`

	UsersName      string `json:"users_name"`
	UsersNpm       string `json:"users_npm"`
	PrivilegesName string `json:"privileges_name"`
	UsersOutlet    string `json:"users_outlet"`
	UsersBranch    string `json:"users_branch"`
	IdCmsUsersOh   int    `json:"id_cms_users_oh"`
	IdCmsUsersSpv  int    `json:"id_cms_users_spv"`
	OhName         string `json:"oh_name"`
	SpvName        string `json:"spv_name"`
}

type DealStatus struct {
	Deals   int `json:"deals"`
	Approve int `json:"approve"`
	Process int `json:"process"`
	Cancel  int `json:"cancel"`
	Reject  int `json:"reject"`
	Paid    int `json:"paid"`
}

type Conv struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type Chart struct {
	Month   string `json:"month"`
	Booking int    `json:"booking"`
	Lead    int    `json:"lead"`
}

type ChartVisum struct {
	Status string `json:"status"`
	Target int    `json:"target"`
	Lead   int    `json:"lead"`
}
type GetOrderBranch struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

type GetOrderChart struct {
	Booking int `json:"booking"`
}
type GetLeadChart struct {
	Leads int `json:"leads"`
}

type GetOrderPhoto struct {
	Id      int    `json:"id"`
	idOrder int    `json:"id_order"`
	Photo   string `json:"photo"`
}
