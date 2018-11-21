package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/oumeniOS/go-gin-blog/pkg/logging"
)

func MarkErrors(errors []*validation.Error)  {
	for _, err := range errors{
		logging.Info(err.Key,err.Value)
	}
	return
}















































