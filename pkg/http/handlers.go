package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zhongweili/vue-go-spa/pkg/core"
	"github.com/zhongweili/vue-go-spa/pkg/following"
)

type Service struct {
	repo   core.Repository
	Router http.Handler
}

func New(repo core.Repository) Service {
	service := Service{
		repo: repo,
	}

	router := httprouter.New()
	router.GET("/followings", service.Index)
	router.POST("/followings", service.Create)
	router.DELETE("/followings/:id", service.Delete)
	router.PUT("/followings/:id", service.Update)

	service.Router = UseMiddlewares(router)

	return service
}

func (s Service) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := following.NewService(s.repo, r.Context().Value("userId").(string))
	followings, err := service.GetFollowings()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(followings)
}

func (s Service) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := following.NewService(s.repo, r.Context().Value("userId").(string))
	payload, _ := ioutil.ReadAll(r.Body)

	githubUser := following.GitHubUser{}
	json.Unmarshal(payload, &githubUser)

	following, err := service.CreateFollowingFor(githubUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(following)
}

func (s Service) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := following.NewService(s.repo, r.Context().Value("userId").(string))

	ID, _ := strconv.Atoi(params.ByName("id"))
	githubUser := following.GitHubUser{ID: int64(ID)}

	_, err := service.RemoveFollowing(githubUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s Service) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := following.NewService(s.repo, r.Context().Value("userId").(string))
	payload, _ := ioutil.ReadAll(r.Body)

	githubUser := following.GitHubUser{}
	json.Unmarshal(payload, &githubUser)

	following, err := service.UpdateFollowingWith(githubUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(following)
}
