package tag_service

import (
	"github.com/oumeniOS/go-gin-blog/models"
	"github.com/oumeniOS/go-gin-blog/service/cache_service"
	"github.com/oumeniOS/go-gin-blog/pkg/gredis"
	"encoding/json"
	"strconv"
	"time"
	"github.com/oumeniOS/go-gin-blog/pkg/export"
	"github.com/tealeg/xlsx"
		"io"
	"github.com/xuri/excelize"
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
	existId := false
	existName := false
	existId, _ = models.ExistTagById(t.ID)
	existName, _ = models.ExistTagByName(t.Name)
	return existId || existName , nil
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

func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range tags {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
			v.CreatedAt.Format("2006年01月02日 15时04分05秒"),
			v.ModifiedBy,
			v.UpdatedAt.Format("2006年01月02日 15时04分05秒"),
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + time + ".xlsx"

	fullPath := export.GetExcelFullPath() + filename
	err = file.Save(fullPath)
	if err != nil {
		return "", err
	}

	return filename, nil
}


func (t *Tag) Import(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	rows := xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}

			models.AddTag(data[1], 1, data[2])
		}
	}

	return nil
}