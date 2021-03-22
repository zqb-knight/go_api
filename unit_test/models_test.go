package unit_test

import (
	"go_api/common/handlers/config"
	"testing"
	//"github.com/stretchr/testify/assert"
	"go_api/model"
)

func TestSomething(t *testing.T) {
	config.InitConfig("")
	model.Init()
	model.CloseDB()

}
