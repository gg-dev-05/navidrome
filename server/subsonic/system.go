package subsonic

import (
	"net/http"

	"github.com/navidrome/navidrome/server/subsonic/responses"
)

type SystemController struct{}

func NewSystemController() *SystemController {
	return &SystemController{}
}

func (api *Router) Ping(_ *http.Request) (*responses.Subsonic, error) {
	return newResponse(), nil
}

func (api *Router) GetLicense(_ *http.Request) (*responses.Subsonic, error) {
	response := newResponse()
	response.License = &responses.License{Valid: true}
	return response, nil
}
