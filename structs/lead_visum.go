package structs

type LeadVisumUsers struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Npm               string `json:"npm"`
	OutletName        string `json:"outlet_name"`
	BranchName        string `json:"branch_name"`
	PrivilegesName    string `json:"privileges_name"`
	Visum             int    `json:"visum"`
	RumahKosong       int    `json:"rumah_kosong"`
	AlamatTidakSesuai int    `json:"alamat_tidak_sesuai"`
	DitolakAtauDiusir int    `json:"ditolak_atau_diusir"`
	ButuhWaktu        int    `json:"butuh_waktu"`
	TanyaPasangan     int    `json:"tanya_pasangan"`
	PikirPikir        int    `json:"pikir_pikir"`
	Berminat          int    `json:"berminat"`
	RumahTidakKetemu  int    `json:"rumah_tidak_ketemu"`
	TidakBerminat     int    `json:"tidak_berminat"`
}

type LeadVisumDetail struct {
	Id         int    `json:"id"`
	CreatedAt  string `json:"created_at"`
	Datasource string `json:"datasource"`
	Revisit    string `json:"revisit"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Alamat     string `json:"alamat"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	Kabupaten  string `json:"kabupaten"`
	Provinsi   string `json:"provinsi"`
	KodePost   int    `json:"kode_post"`
	Status     string `json:"status"`
}

type LeadVisumDownload struct {
	Id             int64  `json:"id"`
	IdLead         int64  `json:"id_lead"`
	CreatedAt      string `json:"created_at"`
	Datasource     string `json:"datasource"`
	Revisit        string `json:"revisit"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Alamat         string `json:"alamat"`
	Kelurahan      string `json:"kelurahan"`
	Kecamatan      string `json:"kecamatan"`
	Kabupaten      string `json:"kabupaten"`
	Provinsi       string `json:"provinsi"`
	VisumStatus    string `json:"visum_status"`
	Name           string `json:"name"`
	Npm            string `json:"npm"`
	PrivilegesName string `json:"privileges_name"`
	LeadOutletName string `json:"lead_outlet_name"`
	LeadBranchName string `json:"lead_branch_name"`
	Number         string `json:"number"`
	OhName         string `json:"oh_name"`
	SpvName        string `json:"spv_name"`
}

type LeadVisumPhoto struct {
	Id          int    `json:"id"`
	IdLeadVisum int    `json:"id_lead_visum"`
	Photo       string `json:"photo"`
}

type LeadVisumChart struct {
	Id        int `json:"id"`
	LeadVisum int `json:"lead_visum"`
}
