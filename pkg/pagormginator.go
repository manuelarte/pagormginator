package pkg

import (
	"github.com/manuelarte/pagormginator/pkg/model"
	"gorm.io/gorm"
)

var _ gorm.Plugin = new(Pagormginator)

type Pagormginator struct {
}

func (p Pagormginator) Name() string {
	return "Pagormginator"
}

func (p Pagormginator) Initialize(db *gorm.DB) error {
	db.Callback().Query().Before("gorm:query").Register("pagorminator:isPagination", p.isPagination)
	//TODO implement me
	//panic("implement me")
	return nil
}

func (p Pagormginator) isPagination(db *gorm.DB) {
	if db.Statement.Schema != nil {
		queryModel := db.Statement.Model
		if page, ok := queryModel.(*model.Page[any]); ok {
			println(page.Page)
		}
	}
}
