package question

import (
	"fmt"
	"net/http"

	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/repository/option"
	"github.com/delicioushwan/magickodung/repository/question"
	"github.com/delicioushwan/magickodung/utils/authUtils"
	"github.com/delicioushwan/magickodung/utils/httpUtils"

	"github.com/labstack/echo/v4"
)

type QuestionsController struct {
	QRepo question.QuestionInterface
	ORepo option.OptionInterface
}

func NewQuestionsControllers(
	qtrep question.QuestionInterface,
	oprep option.OptionInterface,
) *QuestionsController {
	return &QuestionsController{
		QRepo: qtrep,
		ORepo: oprep,
	}
}

func (ctrl QuestionsController) CreateQuestion() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := CreateQuestionRequestFormat{}
		fmt.Println("#@%#$%#ㄲㄸㅆㄸㄲㅆㄸㄲㅆㄸㄲㅆ", c.Request().Body)
		if err := httpUtils.BindAndValidate(c, &req); err != nil {
			return httpUtils.NewBadRequest(err)
		}
		fmt.Println("9843489759384753489753897598", c.Request().Body)

		var userId uint64
		if userId = authUtils.CurrentUser(c); userId == 0 {
			userId, _ = authUtils.CurrentVisitor(c)
		}

		newQuestion := entities.Question{
			Title: req.Title,
			CategoryId: req.CategoryId,
			UserId: userId,
			State: "created",
		}

		question, err := ctrl.QRepo.Create(newQuestion)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}

		var newOptions []entities.Option
		for _, opt := range req.Options {
			newOptions = append(newOptions, entities.Option {
				QuetionId: question.QuestionId,
				Option: opt,
				State: "created",
			})
		}

		err = ctrl.ORepo.Create(newOptions)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}
	
		return c.JSON(http.StatusOK, "ok")
	}

}
