package main

import (
	"fmt"
	call_api "go_test/rest_methods"
)

func main() {
	//call_api_go.get_method()

	call_api.Get_method()
	fmt.Println("======================================")
	call_api.Get_method_by_id(5)
	fmt.Println("======================================")
	call_api.Post_method("Apple MacBook Pro 16", 2024, 2500.55, "Intel Core i9", "1 TB")
	fmt.Println("======================================")
	call_api.Put_method("Apple MacBook Pro 16", 2024, 2500.55, "Intel Core i9", "1 TB", "silver")
}
