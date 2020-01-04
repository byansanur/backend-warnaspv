package structs


type Response struct {
	ApiStatus  int         `json:"api_status,omitempty"`
	ApiMessage string      `json:"api_message"`
	Data       interface{} `json:"data"`
}

// type JsonResponseFirebaseData struct {
// 	Identitas interface{} `json:"identitas"`
// }

type JsonResponseClaimsFB struct {
	Auth_Time    int64                    `json:"auth_time"`
	Phone_Number string                   `json:"phone_number"`
	User_Id      string                   `json:"user_id"`
}

type ResponseTbPeserta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
