package structs

type TargetVisumUsers struct {
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

type TargetVisumDetail struct {
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

type TargetVisumDownload struct {
	Id              int64  `json:"id"`
	IdTarget        int64  `json:"id_target"`
	CreatedAt       string `json:"created_at"`
	Datasource      string `json:"datasource"`
	Revisit         string `json:"revisit"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Alamat          string `json:"alamat"`
	Kelurahan       string `json:"kelurahan"`
	Kecamatan       string `json:"kecamatan"`
	Kabupaten       string `json:"kabupaten"`
	Provinsi        string `json:"provinsi"`
	VisumStatus     string `json:"visum_status"`
	Name            string `json:"name"`
	Npm             string `json:"npm"`
	PrivilegesName  string `json:"privileges_name"`
	UsersOutletName string `json:"users_outlet_name"`
	UsersBranchName string `json:"users_branch_name"`
	HpSatu          string `json:"hp_satu"`
	HpDua           string `json:"hp_dua"`
	OhName          string `json:"oh_name"`
	SpvName         string `json:"spv_name"`
}

type TargetVisumPhoto struct {
	Id            int    `json:"id"`
	IdTargetVisum int    `json:"id_target_visum"`
	Photo         string `json:"photo"`
}

type TargetVisumChart struct {
	Id          int `json:"id"`
	TargetVisum int `json:"target_visum"`
}
