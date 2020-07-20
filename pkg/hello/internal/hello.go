package internal

import (
	"fmt"
	"github.com/pgmot/awesomelib"
)

func InternalHello(name string) string {
	data := awesomelib.NewHoge(name)
	fmt.Println(data)

	return fmt.Sprintf("Hello %s!", name)
}
