package utils

import (
	"fmt"
	"github.com/txgruppi/werr"
	"log"
)

func HandleError(err error) {
	if err != nil {
		err = werr.Wrap(err)
		if wrapped, ok := err.(*werr.Wrapper); ok {
			lg, _ := wrapped.Log()
			fmt.Println(lg)
		}

		log.Fatal(fmt.Sprintf("Error : %s", err.Error()))
	}
}
