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
		check       func(authorizer auth.Authorizer) error
	}{
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticatingUsingClientCertificate = true
				return
			},
			check: func(a auth.Authorizer) error {
				b, ok := a.(*auth.CachedAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer was not an *auth.CachedAuthorizer")
				}

				_, ok = b.Source.(*auth.ClientAssertionAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer source was not an *auth.ClientAssertionAuthorizer")
				}

				return nil
			},
		},
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticatingUsingClientSecret = true
				return
			},
			check: func(a auth.Authorizer) error {
				b, ok := a.(*auth.CachedAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer was not an *auth.CachedAuthorizer")
				}

				_, ok = b.Source.(*auth.ClientSecretAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer source was not an *auth.ClientSecretAuthorizer")
				}

				return nil
			},
		},
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticatingUsingManagedIdentity = true
				return
			},
			check: func(a auth.Authorizer) error {
				b, ok := a.(*auth.CachedAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer was not an *auth.CachedAuthorizer")
				}

				_, ok = b.Source.(*auth.ManagedIdentityAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer source was not an *auth.ManagedIdentityAuthorizer")
				}

				return nil
			},
		},
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticationUsingGitHubOIDC = true
				return
			},
			check: func(a auth.Authorizer) error {
				b, ok := a.(*auth.CachedAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer was not an *auth.CachedAuthorizer")
				}

				_, ok = b.Source.(*auth.GitHubOIDCAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer source was not an *auth.GitHubOIDCAuthorizer")
				}

				return nil
			},
		},
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticationUsingOIDC = true
				return
			},
			check: func(a auth.Authorizer) error {
				b, ok := a.(*auth.CachedAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer was not an *auth.CachedAuthorizer")
				}

				_, ok = b.Source.(*auth.ClientAssertionAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer source was not an *auth.ClientAssertionAuthorizer")
				}

				return nil
			},
		},
		{
			credentials: func() auth.Credentials {
				return credentials
			},
			shouldError: true,
		},
	}

	for i, testCase := range testCases {
		authorizer, err := auth.NewAuthorizerFromCredentials(ctx, testCase.credentials(), env.MicrosoftGraph)
		if testCase.shouldError {
			if err == nil {
				t.Errorf("Test Case #%d: NewAuthorizerFromCredentials() should have errored but no error was returned", i)
			}
		} else {
			if err != nil {
				t.Errorf("Test Case #%d: NewAuthorizerFromCredentials() returned an error: %+v", i, err)
			}
			if authorizer == nil {
				t.Errorf("Test Case #%d: NewAuthorizerFromCredentials() returned a nil Authorizer", i)
			}
			if testCase.check != nil {
				if err = testCase.check(authorizer); err != nil {
					t.Error(err)
				}
			}
		}
	}
}

func TestAccNewAuthorizerFromCredentials(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	credentials := auth.Credentials{
		AuxiliaryTenantIDs:            test.AuxiliaryTenantIds,
		ClientCertificateData:         test.Base64DecodeCertificate(t, test.ClientCertificate),
		ClientCertificatePassword:     test.ClientCertPassword,
		ClientCertificatePath:         test.ClientCertificatePath,
		ClientID:                      test.ClientId,
		ClientSecret:                  test.ClientSecret,
		CustomManagedIdentityEndpoint: test.CustomManagedIdentityEndpoint,
		Environment:                   *env,
		GitHubOIDCTokenRequestToken:   test.GitHubToken,
		GitHubOIDCTokenRequestURL:     test.GitHubTokenURL,
		OIDCAssertionToken:            test.IdToken,
		TenantID:                      test.TenantId,
	}

	testCases := []struct {
		credentials func() auth.Credentials
		shouldError bool
		check       func(authorizer auth.Authorizer) error
	}{
		{
			credentials: func() (ret auth.Credentials) {
				ret = credentials
				ret.EnableAuthenticatingUsingAzureCLI = true
				ret.AzureCliSubscriptionIDHint = test.SubscriptionId
				return
			},
			check: func(a auth.Authorizer) error {
				b, ok := a.(*auth.CachedAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer was not an *auth.CachedAuthorizer")
				}

				c, ok := b.Source.(*auth.AzureCliAuthorizer)
				if !ok {
					return fmt.Errorf("authorizer source was not an *auth.AzureCliAuthorizer")
				}

				if c.TenantID != test.TenantId {
					return fmt.Errorf("unexpected value for authorizer TenantID, expected: %q, saw: %q", test.TenantId, c.TenantID)
				}

				return nil
			},
		},
	}

	for i, testCase := range testCases {
		authorizer, err := auth.NewAuthorizerFromCredentials(ctx, testCase.credentials(), env.MicrosoftGraph)
		if testCase.shouldError {
			if err == nil {
				t.Errorf("Test Case #%d: NewAuthorizerFromCredentials() should have errored but no error was returned", i)
			}
		} else {
			if err != nil {
				t.Errorf("Test Case #%d: NewAuthorizerFromCredentials() returned an error: %+v", i, err)
			}
			if authorizer == nil {
				t.Errorf("Test Case #%d: NewAuthorizerFromCredentials() returned a nil Authorizer", i)
			}
			if testCase.check != nil {
				if err = testCase.check(authorizer); err != nil {
					t.Error(err)
				}
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
