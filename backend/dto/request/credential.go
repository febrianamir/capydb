package request

type SaveCredential struct {
	Title        string `json:"title"`
	DBVendor     string `json:"db_vendor"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type GetCredentials struct {
	Search string `json:"search"`
}

type DeleteCredential struct {
	CredentialId uint `json:"credential_id"`
}
