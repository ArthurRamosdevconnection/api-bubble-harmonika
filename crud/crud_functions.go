package crud

import (
	"fmt"
)

func (r *GormHandler[T]) Create(toCreate T) (created T, err error) {
	err = r.db.Create(&toCreate).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(toCreate)
	fmt.Println(toCreate.GetId())
	return toCreate, nil
}

func (r *GormHandler[T]) GetById(id int) (result T, err error) {
	err = r.db.First(&result, id).Error
	if err != nil {
		return
	}
	return result, nil
}

func (r *GormHandler[T]) Update(toUpdate T) (updated T, err error) {
	err = r.db.Save(&toUpdate).Error
	if err != nil {
		return
	}
	return toUpdate, nil
}

func (r *GormHandler[T]) Delete(id int) (deleted T, err error) {
	err = r.db.Delete(&deleted, id).Error
	if err != nil {
		return
	}
	return deleted, nil
}

func (r *GormHandler[T]) GetAll() (result []T, err error) {
	err = r.db.Find(&result).Error
	if err != nil {
		return
	}
	return result, nil
}
