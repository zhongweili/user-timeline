package following

import (
	"strconv"

	"github.com/zhongweili/user-timeline/pkg/core"
)

type GitHubUser struct {
	ID          int64  `json:"id"`
	UserURL     string `json:"html_url"`
	UserName    string `json:"login"`
	Avatar      string `json:"avatar_url"`
	Description string `json:"subscription_url"`
}

type Service struct {
	userId string
	repo   core.Repository
}

func (s Service) GetFollowings() ([]*core.Following, error) {
	return s.repo.FindAll(map[string]interface{}{"userId": s.userId})
}

func (s Service) CreateFollowingFor(githubUser GitHubUser) (*core.Following, error) {
	following := s.githubUserToFollowing(githubUser)
	err := s.repo.Create(following)
	if err != nil {
		return nil, err
	}
	return following, nil
}

func (s Service) UpdateFollowingWith(githubUser GitHubUser) (*core.Following, error) {
	following := s.githubUserToFollowing(githubUser)
	err := s.repo.Update(following)
	if err != nil {
		return nil, err
	}
	return following, nil
}

func (s Service) RemoveFollowing(githubUser GitHubUser) (*core.Following, error) {
	following := s.githubUserToFollowing(githubUser)
	err := s.repo.Delete(following)
	if err != nil {
		return nil, err
	}
	return following, nil
}

func (s Service) githubUserToFollowing(githubUser GitHubUser) *core.Following {
	return &core.Following{
		UserID:      s.userId,
		UID:         strconv.Itoa(int(githubUser.ID)),
		UserName:    githubUser.UserName,
		UserURL:     githubUser.UserURL,
		Avatar:      githubUser.Avatar,
		Description: githubUser.Description,
	}
}

func NewService(repo core.Repository, userId string) Service {
	return Service{
		repo:   repo,
		userId: userId,
	}
}
