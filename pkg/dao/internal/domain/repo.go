package domain

import "gorm.io/gorm"

type Repo interface {
	CreateOne(dest interface{}) error
	UpdateOne(dest interface{}, conds ...interface{}) error
	RetrieveOne(dest interface{}, conds ...interface{}) error
	RetrieveAll(dest interface{}, conds ...interface{}) error
	RetrieveWithPaging(dest interface{}, conds ...interface{}) error
	CreateOrUpdate(dest interface{}) error
}

type GormRepo struct {
	*gorm.DB
}

func (r *GormRepo) CreateOne(dest interface{}) error {
	return r.Create(dest).Error
}

func (r *GormRepo) UpdateOne(dest interface{}) error {
	return r.Update(dest).Error
}

func (r *GormRepo) RetrieveOne(dest interface{}, conds ...interface{}) error {
	return r.Take(dest, conds...).Error
}

func (r *GormRepo) RetrieveAll(dest interface{}, conds ...interface{}) error {
	return r.Find(dest, conds...).Error
}

func (r *GormRepo) RetrieveWithPaging(dest interface{}, conds ...interface{}) error {
	return r.Find(dest, conds...).Error
}

func (r *GormRepo) CreateOrUpdate(dest interface{}) error {
	return r.Create(dest).Error
}
