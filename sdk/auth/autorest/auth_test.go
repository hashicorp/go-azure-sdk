package autorest_test

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	authWrapper "github.com/hashicorp/go-azure-sdk/sdk/auth/autorest"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAuthorizerWithAuthorization(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()
	conn := test.NewConnection(t)
	api := conn.AuthConfig.Environment.MicrosoftGraph
	conn.Authorize(ctx, t, api)

	endpoint, exists := api.Endpoint()
	if !exists {
		t.Fatalf("could not find endpoint for API %q", api.Name())
	}

	wrapper := &authWrapper.Authorizer{Authorizer: conn.Authorizer}
	if err := testWithAuthorization(wrapper, *endpoint); err != nil {
		t.Fatal(err)
	}
}

func TestAuthorizerBearerAuthorizerCallback(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()
	conn := test.NewConnection(t)
	api := conn.AuthConfig.Environment.KeyVault
	conn.Authorize(ctx, t, api)

	wrapper := &authWrapper.Authorizer{Authorizer: conn.Authorizer}

	callback := wrapper.BearerAuthorizerCallback()
	if err := testWithAuthorization(callback, "https://contoso.vault.azure.net/secrets"); err != nil {
		t.Fatal(err)
	}
}

type preparer struct{}

func (preparer) Prepare(r *http.Request) (*http.Request, error) {
	return r, nil
}

func testWithAuthorization(authorizer autorest.Authorizer, resource string) error {
	u, err := url.Parse(resource)
	if err != nil {
		return err
	}

	r := &http.Request{
		URL: u,
	}

	r, err = authorizer.WithAuthorization()(preparer{}).Prepare(r)
	if err != nil {
		return err
	}

	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		return errors.New("WithAuthorization(): Authorization header has no bearer token")
	}

	return nil
}
