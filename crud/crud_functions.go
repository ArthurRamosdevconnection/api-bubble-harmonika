package crud

func (r *gormHandler[T]) Create(toCreate T) (created T, err error) {
	err = r.db.Create(&toCreate).Error
	if err != nil {
		return
	}
	return toCreate, nil
}

func (r *gormHandler[T]) GetById(id int) (result T, err error) {
	err = r.db.First(&result, id).Error
	if err != nil {
		return
	}
	return result, nil
}

func (r *gormHandler[T]) Update(toUpdate T) (updated T, err error) {
	err = r.db.Save(&toUpdate).Error
	if err != nil {
		return
	}
	return toUpdate, nil
}

func (r *gormHandler[T]) Delete(id int) (deleted T, err error) {
	err = r.db.Delete(&deleted, id).Error
	if err != nil {
		return
	}
	return deleted, nil
}

func (r *gormHandler[T]) GetAll() (result []T, err error) {
	err = r.db.Find(&result).Error
	if err != nil {
		return
	}
	return result, nil
}
