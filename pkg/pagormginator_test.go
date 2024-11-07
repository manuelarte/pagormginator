package pkg

import (
	"github.com/manuelarte/pagormginator/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type TestStruct struct {
	gorm.Model
	Code  string
	Price uint
}

func TestPaginationScopeMetadata(t *testing.T) {
	tests := map[string]struct {
		toMigrate    []*TestStruct
		pageRequest  model.PageRequest
		expectedPage *model.Page[*TestStruct]
	}{
		"unpaged": {
			toMigrate: []*TestStruct{
				{Code: "1"}, {Price: 1},
			},
			pageRequest: model.UnPaged,
			expectedPage: &model.Page[*TestStruct]{
				Page:          0,
				Size:          1,
				TotalElements: 1,
				TotalPages:    0,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
			if err != nil {
				panic("failed to connect database")
			}

			// Migrate the schema
			db.AutoMigrate(&TestStruct{})
			db.CreateInBatches(&test.toMigrate, len(test.toMigrate))
			db.Use(Pagormginator{})

			// Read
			var product TestStruct
			db.First(&product, 1) // find product with integer primary key
		})
	}
}
