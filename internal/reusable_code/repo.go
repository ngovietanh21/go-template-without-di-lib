package reusablecode

import (
	"promotion/pkg/databases"
	"promotion/pkg/failure"

	"github.com/gin-gonic/gin"
)

type Repo struct {
	db databases.MySQLDB
}

func newRepo(db databases.MySQLDB) *Repo {
	return &Repo{db}
}

func (r *Repo) GetByCode(ctx *gin.Context, code string) (*ReusableCode, error) {
	var rc ReusableCode
	if err := r.db.Where(&ReusableCode{Code: code}).Take(&rc).
		Error; err != nil {
		return nil, failure.ErrorWithTrace(err)
	}
	return &rc, nil
}
