package usecase

import (
	"capydb/backend/dto/request"
	"capydb/backend/dto/response"
	"capydb/backend/model"
	"errors"
	"fmt"
)

func (u *Usecase) GetTables() ([]string, error) {
	if u.dbConn == nil {
		return []string{}, errors.New("no db connection")
	}

	query := "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name ASC"

	var tableNames = make([]string, 0)
	err := u.dbConn.Raw(query).Scan(&tableNames).Error
	if err != nil {
		return []string{}, err
	}

	return tableNames, nil
}

func (u *Usecase) GetTableColumns(tableName string) ([]map[string]any, error) {
	if u.dbConn == nil {
		return []map[string]any{}, errors.New("no db connection")
	}

	query := fmt.Sprintf("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = '%v' ORDER BY ordinal_position", tableName)

	var columns = make([]map[string]any, 0)
	err := u.dbConn.Raw(query).Scan(&columns).Error
	if err != nil {
		return []map[string]any{}, err
	}

	return columns, nil
}

func (u *Usecase) GetTableRecords(req request.GetTableRecords) (response.GetTableRecords, error) {
	if u.dbConn == nil {
		return response.GetTableRecords{}, errors.New("no db connection")
	}

	var res response.GetTableRecords
	query := fmt.Sprintf("SELECT * FROM %v", req.TableName)
	queryCount := fmt.Sprintf("SELECT COUNT(*) FROM %v", req.TableName)

	if len(req.Conditions) > 0 {
		whereQuery := buildWhereQuery(req.Conditions)
		query += whereQuery
		queryCount += whereQuery
	}

	if req.SortBy != "" {
		query += buildOrderQuery(model.OrderQueryParam{
			SortBy:  req.SortBy,
			OrderBy: req.OrderBy,
		})
	}

	if req.Limit > 0 {
		limitQuery := buildLimitQuery(req.Limit)
		query += limitQuery
	}

	if req.Offset > 0 {
		offsetQuery := buildOffsetQuery(req.Offset)
		query += offsetQuery
	}

	var records = make([]map[string]any, 0)
	err := u.dbConn.Raw(query).Scan(&records).Error
	if err != nil {
		return res, err
	}

	var count int64
	err = u.dbConn.Raw(queryCount).Count(&count).Error
	if err != nil {
		return res, err
	}

	res.Data = records
	res.Limit = req.Limit
	res.Offset = req.Offset
	res.Total = int(count)
	return res, nil
}

func buildOrderQuery(req model.OrderQueryParam) string {
	orderQuery := ""
	if req.SortBy != "" {
		orderQuery = req.SortBy + " DESC"
		if req.OrderBy == "asc" || req.OrderBy == "ASC" {
			orderQuery = req.SortBy + " ASC"
		}
	}
	return fmt.Sprintf(" ORDER BY %s", orderQuery)
}

func buildWhereQuery(conditions []model.Condition) string {
	whereQuery := ""
	for i, condition := range conditions {
		if i == 0 {
			whereQuery += fmt.Sprintf(" WHERE %v %v", condition.Field, condition.Operator)
		} else {
			whereQuery += fmt.Sprintf(" AND %v %v", condition.Field, condition.Operator)
		}
		if condition.FirstValue != "" {
			whereQuery += fmt.Sprintf(" '%v'", condition.FirstValue)
		}
		if condition.SecondValue != "" { // for 'BETWEEN' operator
			whereQuery += fmt.Sprintf(" AND '%v'", condition.SecondValue)
		}
	}
	return whereQuery
}

func buildLimitQuery(limit int) string {
	return fmt.Sprintf(" LIMIT %d", limit)
}

func buildOffsetQuery(offset int) string {
	return fmt.Sprintf(" OFFSET %d", offset)
}
