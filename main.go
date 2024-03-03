package main

import (
	//"github.com/gin-gonic/gin"
	"fmt"
	"helloworld/routers/api"
)

func main() {
	data := []api.Data_structure{}
	fmt.Print("start")
	api.GET_testdata(data)
}

/*





 */
