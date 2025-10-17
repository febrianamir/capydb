package request

import "capydb/backend/model"

type GetTables struct {
	ConnectionId string `json:"connection_id"`
}

type GetTableColumns struct {
	ConnectionId string `json:"connection_id"`
	TableName    string `json:"table_name"`
}

type GetTableRecords struct {
	ConnectionId string            `json:"connection_id"`
	TableName    string            `json:"table_name"`
	Limit        int               `json:"limit"`
	Offset       int               `json:"offset"`
	SortBy       string            `json:"sort_by"`
	OrderBy      string            `json:"order_by"`
	Conditions   []model.Condition `json:"conditions"`
}
