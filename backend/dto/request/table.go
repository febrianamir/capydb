package request

import "capydb/backend/model"

type GetTableRecords struct {
	TableName  string            `json:"table_name"`
	Limit      int               `json:"limit"`
	Offset     int               `json:"offset"`
	SortBy     string            `json:"sort_by"`
	OrderBy    string            `json:"order_by"`
	Conditions []model.Condition `json:"conditions"`
}
