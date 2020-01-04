package structs

type DownloadLead struct {
	Id             int    `json:"id"`
	IdLead         int64  `json:"id_lead"`
	LeadStatus     string `json:"lead_status"`
	Number         string `json:"number"`
	CreatedAt      string `json:"created_at"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Alamat         string `json:"alamat"`
	Kelurahan      string `json:"kelurahan"`
	Kecamatan      string `json:"kecamatan"`
	Kabupaten      string `json:"kabupaten"`
	Provinsi       string `json:"provinsi"`
	Name           string `json:"name"`
	Npm            string `json:"npm"`
	OutletName     string `json:"outlet_name"`
	BranchName     string `json:"branch_name"`
	PrivilegesName string `json:"privileges_name"`
	Job            string `json:"job"`
	Datasource     string `json:"datasource"`
	Duration       int    `json:"duration"`
	Description    string `json:"description"`
	Recall         string `json:"recall"`
	Note           string `json:"note"`
	NamaProduct    string `json:"nama_product"`
	Year           int    `json:"year"`
	Merk           string `json:"merk"`
	Type           string `json:"type"`
	Model          string `json:"model"`
	Otr            int    `json:"otr"`
	Nopol          string `json:"nopol"`
	TaxStatus      string `json:"tax_status"`
	Owner          string `json:"owner"`
}

type LeadDetail struct {
	Id         int    `json:"id"`
	LeadStatus string `json:"lead_status"`
	Number     string `json:"number"`
	CreatedAt  string `json:"created_at"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	DataSource string `json:"data_source"`
	Alamat     string `json:"alamat"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	Kabupaten  string `json:"kabupaten"`
	Provinsi   string `json:"provinsi"`
}

type LeadUsers struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Npm            string `json:"npm"`
	OutletName     string `json:"outlet_name"`
	BranchName     string `json:"branch_name"`
	PrivilegesName string `json:"privileges_name"`
	Lead           int    `json:"lead"`
	New            int    `json:"new"`
	Hot            int    `json:"hot"`
	Working        int    `json:"working"`
	Unqualified    int    `json:"unqualified"`
	Converted      int    `json:"converted"`
	Deleted        int    `json:"deleted"`
}

type LeadStatus struct {
	Lead        int `json:"lead"`
	New         int `json:"new"`
	Hot         int `json:"hot"`
	Working     int `json:"working"`
	Unqualified int `json:"unqualified"`
	Converted   int `json:"converted"`
	Deleted     int `json:"deleted"`
}
