package main

import (
	"context"
	"fmt"

	"go-graphql-access/infra/mysql"
)

// 脚本可以直接调用domain、infra 层
func main() {
	// if err := mysql.InitDB(); err != nil {
	// 	panic(err)
	// }
	user, err := mysql.NewBasicUser().Get(context.Background(), "liwentao")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)
}
