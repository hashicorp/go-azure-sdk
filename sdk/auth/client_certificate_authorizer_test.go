package auth_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAccClientCertificateAuthorizerV2(t *testing.T) {
	test.AccTest(t)

	env, err := environments.FromNamed(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	pfx := test.Base64DecodeCertificate(t, test.ClientCertificate)

	opts := auth.ClientCertificateAuthorizerOptions{
		Environment:  *env,
		Api:          env.MSGraph,
		TenantId:     test.TenantId,
		AuxTenantIds: []string{},
		ClientId:     test.ClientId,
		Pkcs12Data:   pfx,
		Pkcs12Path:   test.ClientCertPassword,
		Pkcs12Pass:   test.ClientCertificatePath,
	}
	authorizer, err := auth.NewClientCertificateAuthorizer(context.Background(), opts)
	if err != nil {
		t.Fatalf("NewClientCertificateAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err := authorizer.Token()
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
