package ocim

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type OpenStackServer struct {
	Name        string    `json:"name"`
	OpenstackId uuid.UUID `json:"openstack_id"`
}

type AppServer struct {
	ID     int             `json:"id"`
	Server OpenStackServer `json:"server"`
}

// GetAppServer calls Ocim to fetch app server information.
func GetAppServer(c *http.Client, baseURL *url.URL, id int) (*AppServer, error) {
	var appServer AppServer

	res, err := c.Get(fmt.Sprintf("%s/api/v1/openedx_appserver/%d/", baseURL.String(), id))
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&appServer); err != nil {
		return nil, err
	}

	return &appServer, nil
}
