package usecase

func (u *Usecase) GetTables() ([]string, error) {
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name ASC"

	var tableNames = make([]string, 0)
	err := u.dbConn.Raw(query).Scan(&tableNames).Error
	if err != nil {
		return []string{}, err
	}

	return tableNames, nil
}

