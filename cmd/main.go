package main

import (
	"main/internal/api"
	"main/internal/database"
)

func main() {
	database.InitDB()
	api.InitServer("127.0.0.1", "8080")
	/*t1 := time.Now()
	time.Sleep(4 * time.Second)
	fmt.Println(time.Since(t1) >= time.Second * 5)
	*/
}
