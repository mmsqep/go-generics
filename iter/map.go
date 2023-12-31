package iter

// Map produces an iterator of values transformed from an input iterator by a simple mapping function.
func Map[F ~func(T) U, T, U any](inp Of[T], f F) Of[U] {
	return Mapx(inp, func(val T) (U, error) {
		return f(val), nil
	})
}

// Mapx is the extended form of [Map].
// It produces an iterator of values transformed from an input iterator by a mapping function.
// If the mapping function returns an error,
// iteration stops and the error is available via the output iterator's Err method.
func Mapx[F ~func(T) (U, error), T, U any](inp Of[T], f F) Of[U] {
	return &mapIter[T, U]{inp: inp, f: f}
}

type mapIter[T, U any] struct {
	inp Of[T]
	f   func(T) (U, error)
	val U
	err error
}

func (m *mapIter[T, U]) Next() bool {
	if !m.inp.Next() {
		m.err = m.inp.Err()
		return false
	}
	m.val, m.err = m.f(m.inp.Val())
	return m.err == nil
}

func (m *mapIter[T, U]) Val() U {
	return m.val
}

func (m *mapIter[T, U]) Err() error {
	return m.err
}
