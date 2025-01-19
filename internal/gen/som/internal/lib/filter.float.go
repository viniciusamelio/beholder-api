// Code generated by github.com/go-surreal/som, DO NOT EDIT.

package lib

type float_ interface {
	~float32 | ~float64 | ~*float32 | ~*float64
}

type Float[M any, T float_] struct {
	*Numeric[M, T]
}

func NewFloat[M any, T float_](key Key[M]) *Float[M, T] {
	return &Float[M, T]{
		Numeric: NewNumeric[M, T](key),
	}
}

type FloatPtr[M any, T float_] struct {
	*Float[M, T]
	*Nillable[M]
}

func NewFloatPtr[M any, T float_](key Key[M]) *FloatPtr[M, T] {
	return &FloatPtr[M, T]{
		Float:    NewFloat[M, T](key),
		Nillable: NewNillable(key),
	}
}

func (f *Float[M, T]) key() Key[M] {
	return f.Numeric.key()
}

func (f *Float[M, T]) Ceil() *Int[M, int] {
	return NewInt[M, int](f.Base.fn("math::ceil"))
}

func (f *Float[M, T]) Fixed(places int) *Float[M, T] {
	return NewFloat[M, T](f.Base.fn("math::fixed", places))
}

func (f *Float[M, T]) Floor() *Int[M, int] {
	return NewInt[M, int](f.Base.fn("math::floor"))
}

func (f *Float[M, T]) Round() *Int[M, int] {
	return NewInt[M, int](f.Base.fn("math::round"))
}
