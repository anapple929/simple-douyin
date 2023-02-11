package rpc_server

import (
	"fmt"
	"testing"
)

func TestGetUsersInfo(t *testing.T) {
	var userId []int64
	userId = append(userId, 51)
	userId = append(userId, 52)
	userId = append(userId, 53)
	//result, _ := GetUsersInfo(userId, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NTEsImV4cCI6MTY3NjExODIxNSwiaXNzIjoiMTEyMjIzMyJ9.WD4uT4aQciCyxq-3kEpIi8XGtUpEDhrx32H6MqkUn2Q")
	result, _ := GetUsersInfo(userId, "")

	fmt.Println(result)
}
