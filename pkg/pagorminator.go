package pkg

import (
	"github.com/manuelarte/pagormginator/pkg/model"
	"gorm.io/gorm"
)

var _ gorm.Plugin = new(Pagorminator)

type Pagorminator struct {
}

func (p Pagorminator) Name() string {
	return "Pagorminator"
}

func (p Pagorminator) Initialize(db *gorm.DB) error {
	db.Callback().Query().Before("gorm:query").Register("pagorminator:isPagination", p.isPagination)
	//TODO implement me
	//panic("implement me")
	return nil
}

func (p Pagorminator) isPagination(db *gorm.DB) {
	if db.Statement.Schema != nil {
		if page, ok := db.Statement.Model.(model.Page[any]); ok {
			print(page)
		}
	}
}
