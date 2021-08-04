package main

import (
	"fmt"

	"github.com/juju/errors"
)

func main() {
	err := errors.New("first error")
	err = errors.Trace(err)
	// err = errors.Trace(err)
	// err = errors.Trace(err)
	err = errors.Annotate(err, "some context")
	// fmt.Println(err)

	// err2 := errors.New("second error")
	// err = errors.Wrap(err, err2)
	// fmt.Println(err)

	err = errors.Maskf(err, "masked")
	// fmt.Println(err)
	// err = errors.Annotate(err, "more context")
	// fmt.Println(err)
	// err = errors.Trace(err)
	// fmt.Println(err)
	fmt.Println(err.Error())
	fmt.Print(errors.ErrorStack(err))
}
