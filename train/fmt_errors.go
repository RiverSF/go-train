package train

import (
	"fmt"
	"time"
)

//与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error。

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func ErrorOut() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return x, nil
}
