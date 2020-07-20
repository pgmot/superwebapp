package hello

import (
	"github.com/pgmot/superwebapp/pkg/hello/internal"
)

func Hello(name string) string {
	return internal.InternalHello(name)
}
