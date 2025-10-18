package usecase

import (
	"capydb/backend/dto/request"
	"capydb/backend/dto/response"
	"capydb/backend/model"
	"errors"
)

func (u *Usecase) SaveCredential(req request.SaveCredential) error {
	if u.dbDataConn == nil {
		return errors.New("no db connection")
	}

	err := u.dbDataConn.Create(&model.Credential{
		Title:        req.Title,
		DBVendor:     req.DBVendor,
		Host:         req.Host,
		Port:         req.Port,
		User:         req.User,
		Password:     req.Password,
		DatabaseName: req.DatabaseName,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetCredentials(req request.GetCredentials) (response.GetCredentials, error) {
	if u.dbDataConn == nil {
		return response.GetCredentials{}, errors.New("no db connection")
	}

	var res response.GetCredentials
	var data []model.Credential
	statement := u.dbDataConn.Model(&model.Credential{})

	if req.Search != "" {
		statement = statement.Where("title ILIKE ?", "%"+req.Search+"%")
	}

	statement.Order("id DESC")
	err := statement.Find(&data).Error
	if err != nil {
		return res, err
	}

	res.Data = data
	return res, nil
}

func (u *Usecase) DeleteCredential(req request.DeleteCredential) error {
	if u.dbDataConn == nil {
		return errors.New("no db connection")
	}

	err := u.dbDataConn.Delete(&model.Credential{}, req.CredentialId).Error
	if err != nil {
		return err
	}

	return nil
}

//
// func (repo *Repository) GetBlogTotal(ctx context.Context, query request.GetBlogQuery) (uint, error) {
// 	ctx, span := sentry.StartSpan(ctx, "repository.GetBlogTotal")
// 	defer span.Finish()
//
// 	var total int64
// 	var blog model.Blog
//
// 	statement := repo.db.Model(&blog)
//
// 	if query.Search != "" {
// 		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
// 		statement = statement.Where("title ILIKE ?", querySearch)
// 	}
// 	if query.Active == 1 {
// 		statement = statement.Where("status = ?", true)
// 	}
//
// 	order := query.GetOrderQuery()
// 	if order != "" {
// 		statement = statement.Order(order)
// 	}
//
// 	err := statement.Count(&total).Error
//
// 	if err != nil {
// 		elk.LogError(ctx, "Error GetBlogTotal", []zap.Field{
// 			zap.Error(err),
// 			zap.Strings("tags", []string{"postgres", "blog", "repo"}),
// 		}...)
//
// 		return uint(total), err
// 	}
//
// 	return uint(total), nil
// }
//
// func (repo *Repository) CreateBlog(ctx context.Context, blog model.Blog) (model.Blog, error) {
// }
//
// func (repo *Repository) GetBlogByID(ctx context.Context, id uint) (blog model.Blog, err error) {
// 	ctx, span := sentry.StartSpan(ctx, "repository.GetBlogByID")
// 	defer span.Finish()
//
// 	tx, ok := ctx.Value("Trx").(*lib.Database)
// 	if !ok {
// 		tx = repo.db
// 	}
//
// 	err = tx.Debug().Preload("BlogTags").First(&blog, id).Error
// 	if err != nil {
// 		elk.LogError(ctx, "Error GetBlogByID", []zap.Field{
// 			zap.Error(err),
// 			zap.Strings("tags", []string{"postgres", "blog", "repo"}),
// 		}...)
//
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return blog, lib.ErrorBlogNotFound
// 		}
//
// 		return blog, err
// 	}
//
// 	return blog, nil
// }
//
// func (repo *Repository) UpdateBlog(ctx context.Context, blog model.Blog) (model.Blog, error) {
// 	ctx, span := sentry.StartSpan(ctx, "repository.UpdateBlog")
// 	defer span.Finish()
//
// 	err := repo.db.Save(&blog).Error
// 	if err != nil {
// 		elk.LogError(ctx, "Error UpdateBlog", []zap.Field{
// 			zap.Error(err),
// 			zap.Strings("tags", []string{"postgres", "blog", "repo"}),
// 		}...)
//
// 		return blog, err
// 	}
//
// 	return blog, nil
// }
//
