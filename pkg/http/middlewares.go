package http

import (
	"context"
	"log"
	"net/http"
	"strings"

	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	"github.com/rs/cors"
)

func OktaAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header["Authorization"]
		jwt, err := validateAccessToken(accessToken)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}
		ctx := context.WithValue(r.Context(), "userId", jwt.Claims["sub"].(string))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateAccessToken(accessToken []string) (*jwtverifier.Jwt, error) {
	parts := strings.Split(accessToken[0], " ")
	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           "https://dev-179387.okta.com/oauth2/default",
		ClaimsToValidate: map[string]string{"aud": "api://default", "cid": "0oa9x42szoVX3cGOu356"},
	}
	verifier := jwtVerifierSetup.New()
	return verifier.VerifyIdToken(parts[1])
}

func JSONApi(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func AccsessLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s", r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

func Cors(h http.Handler) http.Handler {
	corsConfig := cors.New(cors.Options{
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowedMethods: []string{"POST", "PUT", "GET", "PATCH", "OPTIONS", "HEAD", "DELETE"},
		Debug:          true,
	})
	return corsConfig.Handler(h)
}

func UseMiddlewares(h http.Handler) http.Handler {
	h = JSONApi(h)
	h = OktaAuth(h)
	h = Cors(h)
	return AccsessLog(h)
}
