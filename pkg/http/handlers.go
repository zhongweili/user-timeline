package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zhongweili/vue-go-spa/pkg/core"
	"github.com/zhongweili/vue-go-spa/pkg/kudo"
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
	router.GET("/kudos", service.Index)
	router.POST("/kudos", service.Create)
	router.DELETE("/kudos/:id", service.Delete)
	router.PUT("/kudos/:id", service.Update)

	service.Router = UseMiddlewares(router)

	return service
}

func (s Service) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := kudo.NewService(s.repo, r.Context().Value("userId").(string))
	kudos, err := service.GetKudos()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kudos)
}

func (s Service) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := kudo.NewService(s.repo, r.Context().Value("userId").(string))
	payload, _ := ioutil.ReadAll(r.Body)

	githubUser := kudo.GitHubUser{}
	json.Unmarshal(payload, &githubUser)

	kudo, err := service.CreateKudoFor(githubUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kudo)
}

func (s Service) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := kudo.NewService(s.repo, r.Context().Value("userId").(string))

	ID, _ := strconv.Atoi(params.ByName("id"))
	githubUser := kudo.GitHubUser{ID: int64(ID)}

	_, err := service.RemoveKudo(githubUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s Service) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := kudo.NewService(s.repo, r.Context().Value("userId").(string))
	payload, _ := ioutil.ReadAll(r.Body)

	githubUser := kudo.GitHubUser{}
	json.Unmarshal(payload, &githubUser)

	kudo, err := service.UpdateKudoWith(githubUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kudo)
}
