package response

type GetTableRecords struct {
	Data   []map[string]any
	Total  int
	Limit  int
	Offset int
}
