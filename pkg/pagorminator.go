package pkg

import "gorm.io/gorm"

var _ gorm.Plugin = new(Pagorminator)

type Pagorminator struct {
}

func (p Pagorminator) Name() string {
	return "Pagorminator"
}

func (p Pagorminator) Initialize(db *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}
