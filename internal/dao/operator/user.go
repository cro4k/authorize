package operator

import (
	"context"
	"github.com/go-chocolate/chocolate/pkg/database/repository"
	"github.com/cro4k/authorize/internal/app/dependency"
	"github.com/cro4k/authorize/internal/dao/model"
)

type UserOperator struct {
	*repository.ModelOperator[model.User]
}

func NewUserOperator(ctx context.Context, val *model.User) *UserOperator {
	db := dependency.Get().DB
	rep := repository.NewSimpleRepository[model.User](db)
	op := repository.NewModelOperator[model.User](rep, val)
	return &UserOperator{ModelOperator: op}
}
