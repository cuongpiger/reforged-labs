package repository

func NewRepository() IRepository {
	return &repository{}
}

type IRepository interface {
}

type repository struct {
}
