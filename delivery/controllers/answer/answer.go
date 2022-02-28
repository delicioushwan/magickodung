package answer

import (
	"net/http"

	"github.com/delicioushwan/magickodung/entities"
	"github.com/delicioushwan/magickodung/repository/answer"
	"github.com/delicioushwan/magickodung/utils/authUtils"
	"github.com/delicioushwan/magickodung/utils/httpUtils"
	"github.com/labstack/echo/v4"
)



type AnswersController struct {
	ARepo answer.AnswerInterface
}

func NewAnswersControllers(atrep answer.AnswerInterface) *AnswersController {
	return &AnswersController{
		ARepo: atrep,
	}
}


func (ctrl AnswersController) CreateAnswer() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := CreateAnswerRequestFormat{}
		
		if err := httpUtils.BindAndValidate(c, &req); err != nil {
			return httpUtils.NewBadRequest(err)
		}

		var userId uint64
		if userId = authUtils.CurrentUser(c); userId == 0 {
			userId, _ = authUtils.CurrentVisitor(c)
		}

		newAnswer := entities.Answer{
			UserId: userId,
			QuestionId: req.QuestionId,
			OptionId: req.OptionId,
		}

		_, err := ctrl.ARepo.Create(newAnswer)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}
		return c.JSON(http.StatusOK, "ok")
	}
}