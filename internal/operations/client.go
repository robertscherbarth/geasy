package operations

import (
	"context"
	"net/http"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2/google"
)

func GetClient(scopes ...string) *http.Client {
	ctx := context.Background()

	client, err := google.DefaultClient(ctx, scopes)
	cobra.CheckErr(err)

	return client
}
