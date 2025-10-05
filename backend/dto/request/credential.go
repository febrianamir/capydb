package request

type SaveCredential struct {
	Title        string `json:"title"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

type GetCredentials struct {
	Search string `json:"search"`
}
