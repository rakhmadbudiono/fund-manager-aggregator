package pasardana

import "net/http"

type PasardanaClient struct {
	Client  *http.Client
	BaseURL string
}
