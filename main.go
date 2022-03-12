package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/delivery/controllers/answer"
	"github.com/delicioushwan/magickodung/delivery/controllers/question"
	"github.com/delicioushwan/magickodung/delivery/controllers/user"
	"github.com/delicioushwan/magickodung/delivery/routes"
	answerRepo "github.com/delicioushwan/magickodung/repository/answer"
	optionRepo "github.com/delicioushwan/magickodung/repository/option"
	questionRepo "github.com/delicioushwan/magickodung/repository/question"
	userRepo "github.com/delicioushwan/magickodung/repository/user"
	"github.com/delicioushwan/magickodung/utils"
	"github.com/delicioushwan/magickodung/utils/authUtils"
	"github.com/delicioushwan/magickodung/utils/httpUtils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.Timeout())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderCookie, echo.HeaderSetCookie},
		AllowCredentials: true,
	}))
	e.Use(middleware.RemoveTrailingSlash())

	tokenMiddleware := authUtils.NewTokenMiddleware(config.Secret)
	e.Use(tokenMiddleware)
	e.Validator = httpUtils.NewValidator()

	userRepo := userRepo.NewUsersRepo(db)
	questionRepo := questionRepo.NewQuestionsRepo(db)
	optionRepo := optionRepo.NewOptionsRepo(db)
	answerRepo := answerRepo.NewAnswersRepo(db)

	userCtrl := user.NewUsersControllers(userRepo)
	questionCtrl := question.NewQuestionsControllers(questionRepo, optionRepo)
	answerCtrl := answer.NewAnswersControllers(answerRepo)


	routes.RegisterPath(e,  userCtrl, questionCtrl, answerCtrl)

	go func() {
		address := fmt.Sprintf(":%d", config.Port)	
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
