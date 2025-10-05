package model

type Credential struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	DBVendor     string `json:"db_vendor"`
	HexColor     string `json:"hex_color"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (Credential) TableName() string {
	return "credentials"
}
