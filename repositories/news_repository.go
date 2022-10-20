package repositories

import (
	"crud/config"
	"crud/entities"
	"crud/request"

	"github.com/jinzhu/gorm"
)

type NewsRepository interface {
	Fetch(param request.FindParameter) interface{}
	Find(ID int) map[string]interface{}
	Create(request request.NewsRequest) map[string]interface{}
	Update(request request.NewsRequest) map[string]interface{}
	Delete(ID int) string
}

type newsRepository struct {
	db *gorm.DB
}

func New() *newsRepository {
	db := config.SetupDB()

	return &newsRepository{db}
}

func (r *newsRepository) Fetch(param request.FindParameter) interface{} {
	var news []entities.News

	query := r.db.Begin()

	if param.Title != "" {
		query = query.Where("title LIKE ?", "%"+param.Title+"%")
	}

	if param.Description != "" {
		query = query.Where("description LIKE ?", "%"+param.Description+"%")
	}

	query.Find(&news)

	thisMap := []map[string]interface{}{}
	for _, v := range news {
		aMap := map[string]interface{}{
			"news_id":          v.ID,
			"news_title":       v.Title,
			"news_description": v.Description,
			"created_at":       v.CreatedAt.Format("02-01-2006"),
		}
		thisMap = append(thisMap, aMap)
	}

	return thisMap
}

func (r *newsRepository) Find(ID int) map[string]interface{} {
	newsTable := entities.News{}
	r.db.Where("id = ?", ID).First(&newsTable)
	aMap := map[string]interface{}{
		"news_id":          newsTable.ID,
		"news_title":       newsTable.Title,
		"news_description": newsTable.Description,
		"created_at":       newsTable.CreatedAt.Format("02-01-2006"),
	}
	return aMap
}

func (r *newsRepository) Create(request request.NewsRequest) map[string]interface{} {
	newsTable := entities.News{Title: request.Title, Description: request.Description}
	r.db.Create(&newsTable)
	aMap := map[string]interface{}{
		"news_id":          newsTable.ID,
		"news_title":       newsTable.Title,
		"news_description": newsTable.Description,
		"created_at":       newsTable.CreatedAt.Format("02-01-2006"),
	}
	return aMap
}

func (r *newsRepository) Update(request request.NewsRequest, ID int) map[string]interface{} {
	newsTable := entities.News{Title: request.Title, Description: request.Description}
	r.db.Model(&newsTable).Where("id = ?", ID).Updates(&newsTable)
	aMap := map[string]interface{}{
		"news_id":          ID,
		"news_title":       newsTable.Title,
		"news_description": newsTable.Description,
		"created_at":       newsTable.CreatedAt.Format("02-01-2006"),
	}
	return aMap
}

func (r *newsRepository) Delete(ID int) string {
	newsTable := entities.News{}
	r.db.Where("id = ?", ID).Delete(&newsTable)
	return "Delete"
}
