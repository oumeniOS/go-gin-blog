package article_service

import (
	"github.com/oumeniOS/go-gin-blog/models"
	"github.com/oumeniOS/go-gin-blog/service/cache_service"
	"github.com/oumeniOS/go-gin-blog/pkg/gredis"
	"github.com/oumeniOS/go-gin-blog/pkg/logging"
	"encoding/json"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}
	return nil
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	})
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	gredis.Set(key,article,3600)
	return article, nil

}

func (a *Article) GetAll()([]*models.Article, error) {
	var (
		articles, cacheActicles []*models.Article
	)
	cache := cache_service.Article{
		TagID:a.TagID,
		State:a.State,
		PageSize:a.PageSize,
		PageNum:a.PageNum,
	}
	key := cache.GetArticlesKey()
	if gredis.Exists(key){
		data, err := gredis.Get(key)
		if err != nil {
			return nil, err
		}else {
			json.Unmarshal(data,&cacheActicles)
			return cacheActicles, nil
		}
	}
	articles, err := models.ArticleList(a.PageNum,a.PageSize,a.getMaps())
	if err != nil{
		return nil, err
	}
	gredis.Set(key,articles,3600)
	return articles, nil

}

func (a *Article)Delete() error  {
	return models.DeleteArticle(a.ID)
}

func (a *Article)ExistByID()(bool, error)  {
	return models.IsArticleExistById(a.ID)
}

func (a *Article)Count()(int, error)  {
	return models.GetArticleTotal(a.getMaps())
}

func (a *Article)getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	
	maps["delete_at"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}
	return maps
}
































