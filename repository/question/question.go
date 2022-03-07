package question

import (
	"github.com/delicioushwan/magickodung/entities"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionsRepo(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (qt *QuestionRepository) Create(newQuestion entities.Question) (entities.Question, error) {
	err := qt.db.Create(&newQuestion).Error
	return newQuestion, err
}

func (qt *QuestionRepository) GetRandom(userId uint64, categoryId *uint64) ([]entities.GetCommonQuestionsResponse, error) {
	var result []entities.GetCommonQuestionsResponse

	subQuery := qt.db.
	Table(`questions SQ`).
	Select(`question_id`).
	Joins(`LEFT JOIN answers A ON A.question_id = SQ.question_id AND A.user_id = ?`, userId).
	Where(`A.answer_id IS NULL`).
	Where(`Q.user_id != ?`, userId).
	Order(`RAND()`).
	Limit(3)

	if categoryId == nil {
		subQuery.Where(`category_id = ?`, categoryId)
	}

	query := qt.db.Table("questions Q").
	Select(`Q.question_id, Q.title, O.option, O.option_id, O.quantity`).
	Joins(`INNER JOIN options O ON Q.question_id = O.question_id`).
	Joins(`INNER JOIN (?) AS RQ ON RQ.question_id = Q.question_id`, subQuery).
	Where(`Q.state = ?`, "created")

	err := query.Scan(&result).Error

	return result, err
}

func (qt *QuestionRepository) GetQuestionsByUserId(userId uint64, offset int, limit int) ([]entities.GetCommonQuestionsResponse, error) {
	var result []entities.GetCommonQuestionsResponse

	query := qt.db.Table("questions Q").
	Select(`Q.question_id, Q.title, O.option, O.option_id, O.quantity`).
	Joins(`INNER JOIN options O ON Q.question_id = O.question_id`).
	Where(`Q.state = ?`, "created").
	Where(`Q.user_id = ?`, userId).
	Limit(limit).
	Offset(offset)

	err := query.Scan(&result).Error

	return result, err

}