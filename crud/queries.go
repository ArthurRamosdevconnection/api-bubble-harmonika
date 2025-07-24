package crud

import "strconv"

func (r *GormHandler[T]) Where(column string, value any) *GormHandler[T] {
	returnHandler := r
	returnHandler.db = returnHandler.db.Where(column, value)
	return returnHandler
}

func (r *GormHandler[T]) Limit(value int) *GormHandler[T] {
	returnHandler := r
	returnHandler.db = returnHandler.db.Limit(value)
	return returnHandler
}

func (r *GormHandler[T]) Offset(value int) *GormHandler[T] {
	returnHandler := r
	returnHandler.db = returnHandler.db.Offset(value)
	return returnHandler
}

func (r *GormHandler[T]) OrderBy(column string, order string) *GormHandler[T] {
	returnHandler := r
	returnHandler.db = returnHandler.db.Order(column + " " + order)
	return returnHandler
}

func (r *GormHandler[T]) WhereId(id any) *GormHandler[T] {
	switch id.(type) {
	case int, int64, uint, uint64:
	case string:
		_, err := strconv.Atoi(id.(string))
		if err != nil {
			panic("erro ao converter string para int")
		}
	default:
		panic("tipo de ID não suportado")
	}
	returnHandler := r
	returnHandler.db = returnHandler.db.Where("id", id)
	return returnHandler
}
