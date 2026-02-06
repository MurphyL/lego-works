package domain

import "gorm.io/gorm"

type GormRepo struct {
	*gorm.DB
}

func (r *GormRepo) CreateOne(dest interface{}) error {
	return r.Create(dest).Error
}

func (r *GormRepo) UpdateOne(dest interface{}, args ...interface{}) error {
	return r.Update("", dest).Error
}

func (r *GormRepo) RetrieveOne(dest interface{}, args ...interface{}) error {
	return r.Take(dest, args...).Error
}

func (r *GormRepo) RetrieveAll(dest interface{}, args ...interface{}) error {
	return r.Find(dest, args...).Error
}

func (r *GormRepo) RetrieveWithPaging(dest interface{}, args ...interface{}) error {
	return r.Find(dest, args...).Error
}

func (r *GormRepo) CreateOrUpdate(dest interface{}) error {
	return r.Create(dest).Error
}
