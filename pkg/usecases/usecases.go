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

func (u *UsecasesController) GetBond(bondId string) (*bondApi.Bond, error) {
	log.Println("getting bond from db...")
	bond, err := u.Storage.GetBond(bondId)
	if err != nil {
		return nil, err
	}
	return bond, nil
}

func (u *UsecasesController) GetAllBonds() ([]bondApi.Bond, error) {
	log.Println("getting all bonds from db...")
	allBonds, err := u.Storage.GetAllBonds()
	if err != nil {
		return nil, err
	}
	return allBonds, nil
}

func (u *UsecasesController) PublishNewBond(data bondApi.Bond) error {

	log.Println("entering usecases")
	err := u.Storage.InsertNewBond(data)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsecasesController) CreateUser(data bondApi.User) error {

	log.Println("entering usecases")
	err := u.Storage.CreateUser(data)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsecasesController) GetUser(username string) (*bondApi.User, error) {

	log.Println("entering usecases")
	user, err := u.Storage.GetUser(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
