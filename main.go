package main

import (
	"fmt"
	call_api "go_test/rest_methods"
)

func main() {
	//call_api_go.get_method()

	call_api.GetMethod()
	fmt.Println("======================================")
	call_api.GetMethodById(5)
	fmt.Println("======================================")
	call_api.PostMethod("Apple MacBook Pro 16", 2024, 2500.55, "Intel Core i9", "1 TB")
	fmt.Println("======================================")
	call_api.PutMethod("Apple MacBook Pro 16", 2024, 2500.55, "Intel Core i9", "1 TB", "silver")
}
