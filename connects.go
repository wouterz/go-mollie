package mollie

import (
	"golang.org/x/oauth2"
)

type ConnectAPI struct {
	*oauth2.Config
}

func newConnects(clientID, clientSecret, redirectURL string, scopes ...string) *ConnectAPI {
	return &ConnectAPI{
		&oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://www.mollie.com/oauth2/authorize",
				TokenURL: "https://api.mollie.com/oauth2/tokens",
			},
		},
	}

}

func (c *ConnectAPI) Authorize(opts ...oauth2.AuthCodeOption) string {
	csrfString := "toBeGenerated"
	return c.AuthCodeURL(csrfString, opts...)
}

func (c *ConnectAPI) Tokens(code string) (*oauth2.Token, error) {
	return c.Exchange(oauth2.NoContext, code)
}
