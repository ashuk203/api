package services

import (
	"github.com/HackIllinois/api-gateway/config"
	"github.com/arbor-dev/arbor"
	"net/http"
)

const AuthURL = config.AuthURL

const AuthFormat string = "JSON"

var AuthRoutes = arbor.RouteCollection{
	arbor.Route{
		"GithubAuth",
		"POST",
		"/auth/github/",
		GithubAuth,
	},
}

func GithubAuth(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}
