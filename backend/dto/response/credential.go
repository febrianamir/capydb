package response

import "capydb/backend/model"

type GetCredentials struct {
	Data  []model.Credential
	Total int
}
