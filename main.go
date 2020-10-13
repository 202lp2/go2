package main

import (
	"github.com/202lp2/go2/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run("localhost:8085") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
