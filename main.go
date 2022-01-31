package main

import (
	"fmt"
	"log"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/delivery/controllers/user"
	"github.com/delicioushwan/magickodung/delivery/routes"
	userRepo "github.com/delicioushwan/magickodung/repository/user"
	"github.com/delicioushwan/magickodung/utils"
	"github.com/labstack/echo/v4"
)


func main() {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	e := echo.New()

	userRepo := userRepo.NewUsersRepo(db)
	userCtrl := user.NewUsersControllers(userRepo)

	routes.RegisterPath(e,  userCtrl,)

	address := fmt.Sprintf(":%d", config.Port)
	log.Fatal(e.Start(address))

}
