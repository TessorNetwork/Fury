//go:generate go run github.com/swaggo/swag/cmd/swag init --parseDependency --parseInternal -g swagger.go -o ../swagger

package x

import (
	_ "github.com/cosmos/cosmos-sdk/types"
)

// @title TessorNetwork API
// @description Swagger API for TessorNetwork
// @contact.email developers@tessor.network

// @BasePath /

type JSONResult struct {
	Height string      `json:"height" example:"1234"`
	Result interface{} `json:"result"`
}
