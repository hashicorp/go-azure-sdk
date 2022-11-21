package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"golang.org/x/oauth2"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

// Authorizer is anything that can return an access token for authorizing API connections
type Authorizer interface {
	Token() (*oauth2.Token, error)
	AuxiliaryTokens() ([]*oauth2.Token, error)
}

// NewAuthorizer returns a suitable Authorizer depending on what is defined in the Config
// Authorizers are selected for authentication methods in the following preferential order:
// - Client certificate authentication
// - Client secret authentication
// - OIDC authentication
// - GitHub OIDC authentication
// - MSI authentication
// - Azure CLI authentication
//
// Whether one of these is returned depends on whether it is enabled in the Config, and whether sufficient
// configuration fields are set to enable that authentication method.
//
// For client certificate authentication, specify TenantID, ClientID and ClientCertData / ClientCertPath.
// For client secret authentication, specify TenantID, ClientID and ClientSecret.
// For OIDC authentication, specify TenantID, ClientID and IDToken.
// For GitHub OIDC authentication, specify TenantID, ClientID, IDTokenRequestURL and IDTokenRequestToken.
// MSI authentication (if enabled) using the Azure Metadata Service is then attempted
// Azure CLI authentication (if enabled) is attempted last
//
// It's recommended to only enable the mechanisms you have configured and are known to work in the execution
// environment. If any authentication mechanism fails due to misconfiguration or some other error, the function
// will return (nil, error) and later mechanisms will not be attempted.
func (c *Config) NewAuthorizer(ctx context.Context, api environments.Api) (Authorizer, error) {
	// Default token version should be v2
	if c.Version == 0 {
		c.Version = TokenVersion2
	}

	if c.EnableClientCertAuth && strings.TrimSpace(c.TenantID) != "" && strings.TrimSpace(c.ClientID) != "" && (len(c.ClientCertData) > 0 || strings.TrimSpace(c.ClientCertPath) != "") {
		opts := ClientCertificateAuthorizerOptions{
			Environment:  c.Environment,
			Api:          api,
			TokenVersion: c.Version,
			TenantId:     c.TenantID,
			AuxTenantIds: c.AuxiliaryTenantIDs,
			ClientId:     c.ClientID,
			Pkcs12Data:   c.ClientCertData,
			Pkcs12Path:   c.ClientCertPath,
			Pkcs12Pass:   c.ClientCertPassword,
		}
		a, err := NewClientCertificateAuthorizer(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("could not configure ClientCertificate Authorizer: %s", err)
		}
		if a != nil {
			return a, nil
		}
	}

	if c.EnableClientSecretAuth && strings.TrimSpace(c.TenantID) != "" && strings.TrimSpace(c.ClientID) != "" && strings.TrimSpace(c.ClientSecret) != "" {
		opts := ClientSecretAuthorizerOptions{
			Environment:  c.Environment,
			Api:          api,
			TokenVersion: c.Version,
			TenantId:     c.TenantID,
			AuxTenantIds: c.AuxiliaryTenantIDs,
			ClientId:     c.ClientID,
			ClientSecret: c.ClientSecret,
		}
		a, err := NewClientSecretAuthorizer(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("could not configure ClientSecret Authorizer: %s", err)
		}
		if a != nil {
			return a, nil
		}
	}

	if c.EnableOIDCAuth && strings.TrimSpace(c.TenantID) != "" && strings.TrimSpace(c.ClientID) != "" && strings.TrimSpace(c.IDToken) != "" {
		opts := OIDCAuthorizerOptions{
			Environment:        c.Environment,
			Api:                api,
			TenantId:           c.TenantID,
			AuxiliaryTenantIds: c.AuxiliaryTenantIDs,
			ClientId:           c.ClientID,
			FederatedAssertion: c.IDToken,
		}
		a, err := NewOIDCAuthorizer(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("could not configure OIDC Authorizer: %s", err)
		}
		if a != nil {
			return a, nil
		}
	}

	if c.EnableGitHubOIDCAuth && strings.TrimSpace(c.TenantID) != "" && strings.TrimSpace(c.ClientID) != "" && strings.TrimSpace(c.IDTokenRequestURL) != "" && strings.TrimSpace(c.IDTokenRequestToken) != "" {
		opts := GitHubOIDCAuthorizerOptions{
			Api:                 api,
			AuxiliaryTenantIds:  c.AuxiliaryTenantIDs,
			ClientId:            c.TenantID,
			Environment:         c.Environment,
			IdTokenRequestUrl:   c.IDTokenRequestURL,
			IdTokenRequestToken: c.IDTokenRequestToken,
			TenantId:            c.TenantID,
		}
		a, err := NewGitHubOIDCAuthorizer(context.Background(), opts)
		if err != nil {
			return nil, fmt.Errorf("could not configure GitHubOIDC Authorizer: %s", err)
		}
		if a != nil {
			return a, nil
		}
	}

	if c.EnableMsiAuth {
		opts := ManagedIdentityAuthorizerOptions{
			Api:                           api,
			ClientId:                      c.ClientID,
			CustomManagedIdentityEndpoint: c.MsiEndpoint,
		}
		a, err := NewManagedIdentityAuthorizer(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("could not configure MSI Authorizer: %s", err)
		}
		if a != nil {
			return a, nil
		}
	}

	if c.EnableAzureCliToken {
		opts := AzureCliAuthorizerOptions{
			Api:      api,
			TenantId: c.TenantID,
		}
		a, err := NewAzureCliAuthorizer(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("could not configure AzureCli Authorizer: %s", err)
		}
		if a != nil {
			return a, nil
		}
	}

	return nil, fmt.Errorf("no Authorizer could be configured, please check your configuration")
}
