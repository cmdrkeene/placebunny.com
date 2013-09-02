package conf

import (
	"os"
)

var Port string

func init() {
	Port = os.Getenv("PORT")
	if len(Port) == 0 {
		Port = "3000"
	}
}
