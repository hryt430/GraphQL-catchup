package services

import (
	"context"
	"gql_server/graph/model"

	"github.com/volatiletech/sqlboiler/boil"
)

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepoService interface {
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int32) (*model.Issue, error)
}

type Services interface {
	UserService
	RepoService
	IssueService
}

type services struct {
	*userService
	*repoService
	*issueService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:  &userService{exec: exec},
		repoService:  &repoService{exec: exec},
		issueService: &issueService{exec: exec},
	}
}
