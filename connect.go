package mollie

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type oauthCore struct {
	*http.Client
}

type oauthConfig struct {
	*oauth2.Config
}

func (c *oauthConfig) NewClient(token *oauth2.Token) *oauthCore {
	return &oauthCore{c.Client(context.Background(), token)}
}

func dingen() {
	config := oauthConfig{
		&oauth2.Config{
			ClientID:     "app_gDtM3HAduUvUGWNdPPJUVBCF",
			ClientSecret: "hT3pbhRTTSEJn9Mftp5d4JjwkV6tbebgQQvkMwHq",
			RedirectURL:  "https://api.maex.nl/mollie/authorization",
			Scopes:       []string{PaymentsRead, PaymentsWrite},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://www.mollie.com/oauth2/authorize",
				TokenURL: "https://api.mollie.com/oauth2/tokens",
			},
		},
	}

	url := config.AuthCodeURL("csrfToken", oauth2.ApprovalForce)

	token, err := config.Exchange(oauth2.NoContext, "code")
	if err != nil {
		panic(err)
	}

	_ = config.NewClient(token)
}
