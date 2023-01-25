package auth_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAccClientCertificateAuthorizer(t *testing.T) {
	test.AccTest(t)

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	pfx := test.Base64DecodeCertificate(t, test.ClientCertificate)

	ctx := context.Background()
	opts := auth.ClientCertificateAuthorizerOptions{
		Environment:  *env,
		Api:          env.MicrosoftGraph,
		TenantId:     test.TenantId,
		AuxTenantIds: []string{},
		ClientId:     test.ClientId,
		Pkcs12Data:   pfx,
		Pkcs12Path:   test.ClientCertificatePath,
		Pkcs12Pass:   test.ClientCertPassword,
	}
	authorizer, err := auth.NewClientCertificateAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewClientCertificateAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err := authorizer.Token(ctx, &http.Request{})
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
	}
	if token == nil {
		t.Fatalf("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatal("token.AccessToken was empty")
	}
}
