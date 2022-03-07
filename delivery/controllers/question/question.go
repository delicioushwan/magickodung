package question

import (
	"fmt"
	"net/http"
	"strconv"

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
		
		if err := httpUtils.BindAndValidate(c, &req); err != nil {
			return httpUtils.NewBadRequest(err)
		}

		userId := authUtils.CurrentUserId(c)

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
				QuestionId: question.QuestionId,
				Option: opt,
				State: "created",
			})
		}

		_, err = ctrl.ORepo.Create(newOptions)

		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}
	
		return c.JSON(http.StatusOK, "ok")
	}
}

func (ctrl QuestionsController) FindRandomeQuestions() echo.HandlerFunc {
	return func(c echo.Context) error {
		var categoryId *uint64

		id, err := strconv.ParseUint(c.Param("category"), 10, 64)
		*categoryId = id
		if err != nil {
			categoryId = nil
		}
		userId := authUtils.CurrentUserId(c)
		
		res, err := ctrl.QRepo.GetRandom(userId, categoryId)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}

		result := makeQuestionsWithOptions(res)

		return c.JSON(http.StatusOK, result)
	}
}

func (ctrl QuestionsController) FindQuestionsByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		offset, _ := strconv.Atoi(c.Param("offset"))
		limit, _ := strconv.Atoi( c.Param("limit"))

		userId, err := authUtils.CurrentAuthUserId(c)
		if err != nil {
			return err
		}
		fmt.Println("!@#@$RWEFEG#TGEFGET#%$%#$%$%#$T#T#$T#$T#x")
		fmt.Println(offset)
		fmt.Println(limit)

		res, err := ctrl.QRepo.GetQuestionsByUserId(userId, limit, offset)
		if err != nil {
			return httpUtils.NewInternalServerError(err)
		}

		result := makeQuestionsWithOptions(res)
		
		return c.JSON(http.StatusOK, result)
	}
}

func makeQuestionsWithOptions (questions []entities.GetCommonQuestionsResponse) (result []CommonQuestionsResponse) {
	tempMap := map[uint64]CommonQuestionsResponse{}

	for _, val := range(questions) {
		newVal := CommonQuestionsResponse{
			QuestionId: val.QuestionId,
			Title: val.Title,
			Options: []Option{{
					OptionId: val.OptionId,
					Option: val.Option,
					Quantity: val.Quantity,
			},},
		}
		
		if mapVal, ok := tempMap[val.QuestionId]; ok {
			mapVal.Options = append(mapVal.Options, newVal.Options[0])
			tempMap[val.QuestionId] = mapVal
		} else {
			tempMap[val.QuestionId] = newVal
		}
	}

	for  _, value := range tempMap {
		result = append(result, value)
	}

	return result
}