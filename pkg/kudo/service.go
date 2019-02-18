package kudo

import (
	"strconv"

	"github.com/zhongweili/vue-go-spa/pkg/core"
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

func (s Service) GetKudos() ([]*core.Kudo, error) {
	return s.repo.FindAll(map[string]interface{}{"userId": s.userId})
}

func (s Service) CreateKudoFor(githubUser GitHubUser) (*core.Kudo, error) {
	kudo := s.githubUserToKudo(githubUser)
	err := s.repo.Create(kudo)
	if err != nil {
		return nil, err
	}
	return kudo, nil
}

func (s Service) UpdateKudoWith(githubUser GitHubUser) (*core.Kudo, error) {
	kudo := s.githubUserToKudo(githubUser)
	err := s.repo.Update(kudo)
	if err != nil {
		return nil, err
	}
	return kudo, nil
}

func (s Service) RemoveKudo(githubUser GitHubUser) (*core.Kudo, error) {
	kudo := s.githubUserToKudo(githubUser)
	err := s.repo.Delete(kudo)
	if err != nil {
		return nil, err
	}
	return kudo, nil
}

func (s Service) githubUserToKudo(githubUser GitHubUser) *core.Kudo {
	return &core.Kudo{
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
