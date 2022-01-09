package interfaces_test

import (
	"engineering-task/interfaces"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCarRepository_Instantiated_NotNil(t *testing.T) {
	carRepository := interfaces.NewCarRepository()
	assert.NotNil(t, carRepository)
}

func TestGetAllCars_ParsesReturnsCars_NonEmptyResultSet(t *testing.T) {
	carRepo := interfaces.NewCarRepository()
	cars := carRepo.GetAllCars()
	fmt.Println(cars)
}
