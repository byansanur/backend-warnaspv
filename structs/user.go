package structs

type GetUsers struct {
	Id                int    `json:"id"`
	CreatedAt         string `json:"created_at"`
	Name              string `json:"name"`
	Npm               string `json:"npm"`
	Email             string `json:"email"`
	IdMstOutlet       int    `json:"id_mst_outlet"`
	OutletName        string `json:"outlet_name"`
	IdMstBranch       int    `json:"id_mst_branch"`
	BranchName        string `json:"branch_name"`
	IdCmsPrivileges   int    `json:"id_cms_privileges"`
	PrivilegesName    string `json:"privileges_name"`
	IdCmsUsersOh      int    `json:"id_cms_users_oh"`
	IdCmsUsersSpv     int    `json:"id_cms_users_spv"`
	IdCmsUsersSubDept int    `json:"id_cms_users_sub_dept"`
	Spv               string `json:"spv"`
	Oh                string `json:"oh"`
	SubDept           string `json:"sub_dept"`
	Status            string `json:"status"`
}

type GetUsersLogin struct {
	Id              int     `json:"id"`
	CreatedAt       *string `json:"created_at"`
	Name            *string `json:"name"`
	Npm             *string `json:"npm"`
	Email           *string `json:"email"`
	IdMstOutlet     *int    `json:"id_mst_outlet"`
	OutletName      *string `json:"outlet_name"`
	IdMstBranch     *int    `json:"id_mst_branch"`
	BranchName      *string `json:"branch_name"`
	IdCmsPrivileges int     `json:"id_cms_privileges"`
	PrivilegesName  *string `json:"privileges_name"`
	IdCmsUsersOh    *int    `json:"id_cms_users_oh"`
	IdCmsUsersSpv   *int    `json:"id_cms_users_spv"`
	Spv             *string `json:"spv"`
	Oh              *string `json:"oh"`
	Status          *string `json:"status"`
	Token           string  `json:"token"`
}

type CheckUsersLogin struct {
	Id       int    `json:"id"`
	Npm      string `json:"npm"`
	Password string `json:"password"`
}

type CekLogout struct {
	Id       int64  `gorm:"Default:'null'" json:"id"`
	Npm      string `gorm:"Default:'null'" json:"npm"`
	Password string `gorm:"Default:'null'" json:"password"`
}

type GetKPI struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Npm              string `json:"npm"`
	OutletName       string `json:"outlet_name"`
	BranchName       string `json:"branch_name"`
	PrivilegesName   string `json:"privileges_name"`
	Lead             int    `json:"lead"`
	LeadVisum        int    `json:"lead_visum"`
	Orders           int    `json:"orders"`
	Booking          int    `json:"booking"`
	Target           int    `json:"target"`
	TargetNew        int    `json:"target_new"`
	TargetHot        int    `json:"target_hot"`
	TargetWorking    int    `json:"target_working"`
	TargetVisit      int    `json:"target_visit"`
	TargetUnqualifed int    `json:"target_unqualifed"`
	TargetConverted  int    `json:"target_converted"`
	TargetVisum      int    `json:"target_visum"`
	TargetCalls      int    `json:"target_calls"`
}

type Dashboard struct {
	Users      int `json:"users"`
	Brosur     int `json:"brosur"`
	Leads      int `json:"leads"`
	TargetCall int `json:"target_call"`
	Deals      int `json:"deals"`
	Booking    int `json:"booking"`
	Target     int `json:"target"`
}

type UsersStatus struct {
	IdBranch   int    `json:"id_branch"`
	BranchName string `json:"branch_name"`
	IdOutlet   int    `json:"id_outlet"`
	OutletName string `json:"outlet_name"`
	Total      int    `json:"total"`
	Active     int    `json:"active"`
	NoActive   int    `json:"no_active"`
}

type CreateUsers struct {
	CreatedAt       string `json:"created_at"`
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Npm             string `json:"npm"`
	IdMstOutlet     int    `json:"id_mst_outlet"`
	IdCmsPrivileges int    `json:"id_cms_privileges"`
	Password        string `json:"password"`
	Photo           string `json:"photo"`
	Email           string `json:"email"`
	Status          string `json:"status"`
	CreatedBy       int    `json:"created_by"`
	//UpdatedAt string `json:"updated_at"`
}

type UpdateUsers struct {
	//CreatedAt string `json:"created_at"`
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Npm             string `json:"npm"`
	IdMstOutlet     int    `json:"id_mst_outlet"`
	IdCmsPrivileges int    `json:"id_cms_privileges"`
	Password        string `json:"password"`
	Photo           string `json:"photo"`
	Email           string `json:"email"`
	Status          string `json:"status"`
	CreatedBy       int    `json:"created_by"`
	UpdatedAt       string `json:"updated_at"`
}
type CreateCmsUsersCabang struct {
	IdCmsUsers        int    `json:"id_cms_users"`
	IdCmsUsersOh      int    `json:"id_cms_users_oh"`
	IdCmsUsersSpv     int    `json:"id_cms_users_spv"`
	IdCmsUsersSubDept int    `json:"id_cms_users_sub_dept"`
	CreatedAt         string `json:"created_at"`
	CreatedBy         int    `json:"created_by"`
}
type CheckNpm struct {
	CountNpm int `json:"count_npm"`
}

type CreateCmsUsersArea struct {
	IdCmsUsers  int `json:"id_cms_users"`
	IdMstBranch int `json:"id_mst_branch"`
}

type GetCmsUsersArea struct {
	Id         int    `json:"id"`
	IdCmsUsers int    `json:"id_cms_users"`
	BranchName string `json:"branch_name"`
}

type DeleteCmsUsersArea struct {
	Id int `json:"id"`
}
