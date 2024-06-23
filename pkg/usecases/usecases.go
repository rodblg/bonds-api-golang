package usecases

import (
	"log"

	"github.com/rodblg/bonds-api-golang/pkg/bondApi"
	"github.com/rodblg/bonds-api-golang/pkg/database"
)

type UsecasesController struct {
	Storage *database.MongoController
	//Storage bondApi.Storage
}

func NewUsecasesController(s *database.MongoController) *UsecasesController {
	//func NewUsecasesController(s bondApi.Storage) *UsecasesController {
	return &UsecasesController{
		Storage: s,
	}
}

func (u *UsecasesController) SellNewBonds(data bondApi.Bond) error {

	log.Println("entering usecases")
	err := u.Storage.InsertNewData(data)
	if err != nil {
		return err
	}

	return nil
}
