// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	Root *mux.Router

	Players    *mux.Router
	NeedPlayer *mux.Router

	Games    *mux.Router
	NeedGame *mux.Router

	Ai *mux.Router
}

var BaseRoutes *Routes

func InitApi() {
	BaseRoutes = &Routes{}
	BaseRoutes.Root = Srv.Router

	BaseRoutes.Players = Srv.Router.PathPrefix("/players").SubRouter()
	BaseRoutes.NeedPlayer = BaseRoutes.Players.PathPrefix("/{player_id:[A-Za-z0-9]+}").SubRouter()

	BaseRoutes.Games = Srv.Router.PathPrefix("/games").SubRouter()
	BaseRoutes.NeedGame = BaseRoutes.Games.PathPrefix("/{game_id:[A-Za-z0-9]+}").SubRouter()

	InitPlayer()
	InitGame()

	Srv.Router.Handle("/", http.HandlerFunc(Handle404))
}