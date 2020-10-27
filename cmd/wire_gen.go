// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	"github.com/deluan/navidrome/core"
	"github.com/deluan/navidrome/core/transcoder"
	"github.com/deluan/navidrome/persistence"
	"github.com/deluan/navidrome/scanner"
	"github.com/deluan/navidrome/server"
	"github.com/deluan/navidrome/server/app"
	"github.com/deluan/navidrome/server/subsonic"
	"github.com/deluan/navidrome/server/subsonic/engine"
	"github.com/google/wire"
)

// Injectors from wire_injectors.go:

func CreateServer(musicFolder string) *server.Server {
	dataStore := persistence.New()
	serverServer := server.New(dataStore)
	return serverServer
}

func CreateScanner(musicFolder string) scanner.Scanner {
	dataStore := persistence.New()
	artworkCache := core.NewImageCache()
	artwork := core.NewArtwork(dataStore, artworkCache)
	cacheWarmer := core.NewCacheWarmer(artwork)
	scannerScanner := scanner.New(dataStore, cacheWarmer)
	return scannerScanner
}

func CreateAppRouter() *app.Router {
	dataStore := persistence.New()
	router := app.New(dataStore)
	return router
}

func CreateSubsonicAPIRouter() *subsonic.Router {
	dataStore := persistence.New()
	artworkCache := core.NewImageCache()
	artwork := core.NewArtwork(dataStore, artworkCache)
	playlists := engine.NewPlaylists(dataStore)
	transcoderTranscoder := transcoder.New()
	transcodingCache := core.NewTranscodingCache()
	mediaStreamer := core.NewMediaStreamer(dataStore, transcoderTranscoder, transcodingCache)
	archiver := core.NewArchiver(dataStore)
	players := engine.NewPlayers(dataStore)
	client := core.LastFMNewClient()
	spotifyClient := core.SpotifyNewClient()
	externalInfo := core.NewExternalInfo(dataStore, client, spotifyClient)
	router := subsonic.New(artwork, playlists, mediaStreamer, archiver, players, externalInfo, dataStore)
	return router
}

// wire_injectors.go:

var allProviders = wire.NewSet(engine.Set, core.Set, scanner.New, subsonic.New, app.New, persistence.New)
