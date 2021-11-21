package gee

import (
	"fmt"
	"log"
	"net/http"
)

func Recovery() HandleFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", message)
				c.String(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		//c.Next()
	}
}
