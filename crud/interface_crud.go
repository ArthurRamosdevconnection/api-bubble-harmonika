package crud

import "gorm.io/gorm"

type BaseModel struct {
	gorm.Model
}

func (model BaseModel) GetId() int {
	return int(model.ID)
}

type Iminima interface {
	GetId() int
}

type InterfaceCrud[T Iminima] interface {
	Create(toCreate T) (created T, err error)
	GetById(id int) (result T, err error)
	GetAll() (result []T, err error)
	Update(toUpdate T) (updated T, err error)
	Delete(id int) (deleted T, err error)
}

// Esse struct é feito na hora
type gormHandler[T Iminima] struct {
	InterfaceCrud[T]
	db *gorm.DB
}

// NewGormRepository é o construtor que cria uma instância do nosso repositório.
func NewGormRepository[T Iminima](db *gorm.DB) InterfaceCrud[T] {
	return &gormHandler[T]{
		db: db,
	}
}
