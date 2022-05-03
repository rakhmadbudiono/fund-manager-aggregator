package pasardana

import "net/http"

var (
	Client  *http.Client = http.DefaultClient
	BaseURL string       = "https://www.pasardana.id/api"
)
