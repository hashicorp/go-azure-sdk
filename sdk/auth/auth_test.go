// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
	"golang.org/x/oauth2"
)

func TestNewAuthorizerFromCredentials(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()

	credentials := auth.Credentials{
		AuxiliaryTenantIDs: []string{
			"00000000-2222-0000-0000-000000000000",
			"00000000-3333-0000-0000-000000000000",
		},
		ClientCertificateData:         test.Base64DecodeCertificate(t, dummyClientCertificate),
		ClientCertificatePassword:     "certpassword",
		ClientCertificatePath:         "/path/to/cert",
		ClientID:                      "11111111-0000-0000-0000-000000000000",
		ClientSecret:                  "supersecret",
		CustomManagedIdentityEndpoint: "https://endpoint",
		Environment:                   *env,
		GitHubOIDCTokenRequestToken:   "githubtoken",
		GitHubOIDCTokenRequestURL:     "https://githubtokenendpoint",
		OIDCAssertionToken:            "idtokenblurh",
		TenantID:                      "00000000-1111-0000-0000-000000000000",
	}

	testCases := []struct {
		credentials func() auth.Credentials
		shouldError bool
	}{
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticatingUsingAzureCLI = true
				return
			},
		},
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticatingUsingClientCertificate = true
				return
			},
		},
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticatingUsingClientSecret = true
				return
			},
		},
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticatingUsingManagedIdentity = true
				return
			},
		},
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticationUsingGitHubOIDC = true
				return
			},
		},
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticationUsingOIDC = true
				return
			},
		},
	}

	for i, testCase := range testCases {
		authorizer, err := auth.NewAuthorizerFromCredentials(ctx, testCase.credentials(), env.MicrosoftGraph)
		if testCase.shouldError {
			if err == nil {
				t.Fatalf("Test Case #%d: NewAuthorizerFromCredentials() should have errored but no error was returned", i)
			}
		} else {
			if err != nil {
				t.Fatalf("Test Case #%d: NewAuthorizerFromCredentials() returned an error: %+v", i, err)
			}
			if authorizer == nil {
				t.Fatalf("Test Case #%d: NewAuthorizerFromCredentials() returned a nil Authorizer", i)
			}
		}
	}
}

func testObtainAccessToken(ctx context.Context, authorizer auth.Authorizer) (*oauth2.Token, error) {
	token, err := authorizer.Token(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("authorizer.Token(): %v", err)
	}

	if token == nil {
		return nil, fmt.Errorf("token was nil")
	}

	if token.AccessToken == "" {
		return token, fmt.Errorf("token.AccessToken was empty")
	}

	return token, nil
}
