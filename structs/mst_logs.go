package structs

type Riwayat struct {
	Id             int    `json:"id"`
	IdData         int    `json:"id_data"`
	CreatedAt      string `json:"created_at"`
	Npm            string `json:"npm"`
	Name           string `json:"name"`
	PrivilegesName string `json:"privileges_name"`
	OutletName     string `json:"outlet_name"`
	BranchName     string `json:"branch_name"`
	IdCmsUsers     int    `json:"id_cms_users"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Modul          string `json:"modul"`
	Description    string `json:"description"`
	IdCmsUsersOh   int    `json:"id_cms_users_oh"`
	IdCmsUsersSpv  int    `json:"id_cms_users_spv"`
	OhName         string `json:"oh_name"`
	SpvName        string `json:"spv_name"`
}
