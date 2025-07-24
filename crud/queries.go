package crud

import "strconv"

func (r *GormHandler[T]) Where(column string, value any) *GormHandler[T] {
	r.db = r.db.Where(column, value)
	return r
}

func (r *GormHandler[T]) Limit(value int) *GormHandler[T] {
	r.db = r.db.Limit(value)
	return r
}

func (r *GormHandler[T]) Offset(value int) *GormHandler[T] {
	r.db = r.db.Offset(value)
	return r
}

func (r *GormHandler[T]) OrderBy(column string, order string) *GormHandler[T] {
	r.db = r.db.Order(column + " " + order)
	return r
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
	r.db.Where("id", id)
	return r
}
