package tag_service

import (
	"github.com/oumeniOS/go-gin-blog/models"
	"github.com/oumeniOS/go-gin-blog/service/cache_service"
	"github.com/oumeniOS/go-gin-blog/pkg/gredis"
	"encoding/json"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

/*Name:name,
		State:state,
		CreatedBy:createdBy,*/

func (t *Tag) Add() error {
	err := models.AddTag(t.Name, t.State, t.CreatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tag)Exist()(bool, error)  {
	return models.ExistTagById(t.ID)
}

func (t *Tag) Edit() error {
	return models.EditArticle(t.ID, t.getMap())
}

func (t *Tag) GetAll() ([]*models.Tag, error) {
	var (
		tags, cacheTags []*models.Tag
		err error
	)

	cache := cache_service.Tag{
		ID:       t.ID,
		Name:     t.Name,
		State:    t.State,
		PageSize: t.PageSize,
		PageNum:  t.PageNum,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		var data []byte
		data, err = gredis.Get(key)
		if err != nil {
			return nil, err
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	} else {
		tags, err = models.GetTags(t.PageNum, t.PageSize, t.getMap())
		if err != nil {
			return nil, err
		} else {
			gredis.Set(key,tags,3600)
			return tags, nil
		}
	}
}

func (t *Tag) Delete() (error) {
	return models.DeleteTagById(t.ID)
}

func (t *Tag) TotalCount() (int, error) {
	return models.GetTagTotal(t.getMap())
}

func (t *Tag) getMap() (map[string]interface{}) {
	maps := make(map[string]interface{})
	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.CreatedBy != "" {
		maps["modified_by"] = t.CreatedBy
	}
	if t.State != -1 {
		maps["state"] = t.State
	}
	return maps
}
