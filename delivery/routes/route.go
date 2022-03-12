package routes

import (
	"net/http"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/delivery/controllers/answer"
	"github.com/delicioushwan/magickodung/delivery/controllers/category"
	"github.com/delicioushwan/magickodung/delivery/controllers/question"
	"github.com/delicioushwan/magickodung/delivery/controllers/user"
	"github.com/delicioushwan/magickodung/utils/authUtils"

	"github.com/labstack/echo/v4"
)

var config = configs.GetConfig()


func Session (c echo.Context) error {
		return c.JSON(http.StatusOK, "nice")
}

func RegisterPath(
	e *echo.Echo,
	uctrl *user.UsersController,
	qctrl *question.QuestionsController,
	actrl *answer.AnswersController,
) {

	jwtMiddleware := authUtils.NewJWTMiddleware(config.Secret)
	// authorizeMiddleware := authUtils.NewAuthorizeMiddleware(config.Secret)

	e.GET("/session", Session)
	e.GET("/category", category.GetCategory())

	usersGroup := e.Group("/users")
	usersGroup.POST("/signup", uctrl.Signup())
	usersGroup.POST("/login", uctrl.Login())

	questionGroup := e.Group("/questions")
	questionGroup.GET("/", qctrl.FindRandomeQuestions())
	questionGroup.POST("/", qctrl.CreateQuestion())

	answerGroup := e.Group("/answers")
	answerGroup.POST("/", actrl.CreateAnswer())

	userGroup := e.Group("/user")
	userGroup.Use(jwtMiddleware)
	userGroup.GET("/questions",qctrl.FindQuestionsByUserId())
}
