package question

import (
	"net/http"

	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/repository/question"
	"github.com/delicioushwan/magickodung/utils/authUtils"
	"github.com/delicioushwan/magickodung/utils/httpUtils"

	"github.com/labstack/echo/v4"
)

type QuestionsController struct {
	Repo question.QuestionInterface
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
		if err := httpUtils.BindAndValidate(c, &req); err != nil {
			return httpUtils.NewBadRequest(err)
		}
		var userId uint64
		if userId = authUtils.CurrentUser(c); userId == 0 {
			userId = authUtils.CurrentVisitor(c)
		}

		newQuestion := entities.Question{
			Title: req.Title,
			CategoryId: req.CategoryId,
			UserId: userId,
			State: "created",
		}

		_, err := ctrl.Repo.Create(newQuestion)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}

		_, err := ctrl.Repo.
	
		return c.JSON(http.StatusOK, "ok")
	}

}
