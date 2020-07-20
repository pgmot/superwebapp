package internal

import "fmt"

func InternalHello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
