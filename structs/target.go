package structs

type TelemarketingUsers struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Npm              string `json:"npm"`
	OutletName       string `json:"outlet_name"`
	BranchName       string `json:"branch_name"`
	PrivilegesName   string `json:"privileges_name"`
	Total            int    `json:"total"`
	Berminat         int    `json:"berminat"`
	TidakBerminat    int    `json:"tidak_berminat"`
	PikirPikir       int    `json:"pikir_pikir"`
	TanyaPasangan    int    `json:"tanya_pasangan"`
	ButuhWaktu       int    `json:"butuh_waktu"`
	Reject           int    `json:"reject"`
	TidakDiangkat    int    `json:"tidak_diangkat"`
	NoTidakTerdaftar int    `json:"no_tidak_terdaftar"`
	NoTidakAktif     int    `json:"no_tidak_aktif"`
	Cancel           int    `json:"cancel"`
	SalahSambung     int    `json:"salah_sambung"`
}

type TeleDownload struct {
	Id              int64  `json:"id"`
	IdTarget        int64  `json:"id_target"`
	CreatedAt       string `json:"created_at"`
	Duration        string `json:"duration"`
	Description     string `json:"description"`
	Recall          string `json:"recall"`
	DateUpload      string `json:"date_upload"`
	Datasource      string `json:"datasource"`
	Priority        string `json:"priority"`
	Nopol           string `json:"nopol"`
	NoContract      string `json:"no_contract"`
	ProviderSatu    string `json:"provider_satu"`
	ProviderDua     string `json:"provider_dua"`
	Status          string `json:"status"`
	FirstName       string `json:"first_name"`
	Lastname        string `json:"lastname"`
	Alamat          string `json:"alamat"`
	Kelurahan       string `json:"kelurahan"`
	Kecamatan       string `json:"kecamatan"`
	Kabupaten       string `json:"kabupaten"`
	Provinsi        string `json:"provinsi"`
	Name            string `json:"name"`
	Npm             string `json:"npm"`
	PrivilegesName  string `json:"privileges_name"`
	UsersOutletName string `json:"users_outlet_name"`
	UsersBranchName string `json:"users_branch_name"`
	Spv             string `json:"spv"`
	Oh              string `json:"oh"`
}

type GetDataAssignment struct {
	Id           int    `json:"id"`
	CreatedAt    string `json:"created_at"`
	Description  string `json:"description"`
	NoContract   string `json:"no_contract"`
	ProviderSatu string `json:"provider_satu"`
	ProviderDua  string `json:"provider_dua"`
	Status       string `json:"status"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	//Alamat string `json:"alamat"`
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
	Provinsi  string `json:"provinsi"`
	Name      string `json:"name"`
}

type UpdateAssignment struct {
	Id                int    `json:"id"`
	IdCmsUsers        int    `json:"id_cms_users"`
	UpdatedBy         int    `json:"updated_by"`
	IdTargetMstStatus int    `json:"id_target_mst_status"`
	UpdatedAt         string `json:"updated_at"`
}
type CreateTargetAssignmentLog struct {
	Id          int    `json:"id"`
	IdCmsUsers  int    `json:"id_cms_users"`
	TargetUsers int    `json:"target_users"`
	IdTarget    string `json:"id_target"`
	Total       int    `json:"total"`
	//UpdatedAt int `json:"updated_at"`
}

type TargetAssignmentLog struct {
	Id               int    `json:"id"`
	CreatedAt        string `json:"created_at"`
	Name             string `json:"name"`
	Npm              string `json:"npm"`
	OutletName       string `json:"outlet_name"`
	BranchName       string `json:"branch_name"`
	PrivilegesName   string `json:"privileges_name"`
	OhName           string `json:"oh_name"`
	SpvName          string `json:"spv_name"`
	Total            int    `json:"total"`
	TargetName       string `json:"target_name"`
	TargetNpm        string `json:"target_npm"`
	TargetPrivileges string `json:"target_privileges"`
}

type GetTargetUpload struct {
	Id          int    `json:"id"`
	IdCmsUsers  int    `json:"id_cms_users"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Total       string `json:"total"`
	Success     int    `json:"success"`
	Status      string `json:"status"`
	Result      string `json:"result"`
	Url         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetTargetStatus struct {
	Target           int `json:"target"`
	TargetNew        int `json:"target_new"`
	TargetHot        int `json:"target_hot"`
	TargetWorking    int `json:"target_working"`
	TargetVisit      int `json:"target_visit"`
	TargetUnqualifed int `json:"target_unqualifed"`
	TargetConverted  int `json:"target_converted"`
}
