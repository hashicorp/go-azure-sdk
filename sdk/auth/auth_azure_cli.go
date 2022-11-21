package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/azurecli"
	"golang.org/x/oauth2"
	"os"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureCliAuthorizerOptions struct {
	// TODO: document these

	Api      environments.Api
	TenantId string
}

// NewAzureCliAuthorizer returns an Authorizer which authenticates using the Azure CLI.
func NewAzureCliAuthorizer(ctx context.Context, options AzureCliAuthorizerOptions) (Authorizer, error) {
	conf, err := newAzureCliConfig(options.Api, options.TenantId)
	if err != nil {
		return nil, err
	}
	return conf.TokenSource(ctx), nil
}

// AzureCliAuthorizer is an Authorizer which supports the Azure CLI.
type AzureCliAuthorizer struct {
	ctx  context.Context
	conf *azureCliConfig
}

// Token returns an access token using the Azure CLI as an authentication mechanism.
func (a *AzureCliAuthorizer) Token() (*oauth2.Token, error) {
	if a.conf == nil {
		return nil, fmt.Errorf("could not request token: conf is nil")
	}

	azArgs := []string{"account", "get-access-token"}

	// check for MSAL support
	if err := azurecli.CheckAzVersion(azureCliMsalVersion, nil); err == nil {
		azArgs = append(azArgs, "--scope", a.conf.Api.DefaultScope())
	} else {
		azArgs = append(azArgs, "--resource", a.conf.Api.ResourceUrl())
	}

	// Try to detect if we're running in Cloud Shell
	if cloudShell := os.Getenv("AZUREPS_HOST_ENVIRONMENT"); !strings.HasPrefix(cloudShell, "cloud-shell/") {
		// Seemingly not, so we'll append the tenant ID to the az args
		azArgs = append(azArgs, "--tenant", a.conf.TenantID)
	}

	var token azureCliToken
	if err := azurecli.JSONUnmarshalAzCmd(&token, azArgs...); err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
	}, nil
}

// AuxiliaryTokens returns additional tokens for auxiliary tenant IDs, for use in multi-tenant scenarios
func (a *AzureCliAuthorizer) AuxiliaryTokens() ([]*oauth2.Token, error) {
	if a.conf == nil {
		return nil, fmt.Errorf("could not request token: conf is nil")
	}

	// Try to detect if we're running in Cloud Shell
	if cloudShell := os.Getenv("AZUREPS_HOST_ENVIRONMENT"); strings.HasPrefix(cloudShell, "cloud-shell/") {
		return nil, fmt.Errorf("auxiliary tokens not supported in Cloud Shell")
	}

	azArgs := []string{"account", "get-access-token"}

	// check for MSAL support
	if err := azurecli.CheckAzVersion(azureCliMsalVersion, nil); err == nil {
		azArgs = append(azArgs, "--scope", a.conf.Api.DefaultScope())
	} else {
		azArgs = append(azArgs, "--resource", a.conf.Api.ResourceUrl())
	}

	tokens := make([]*oauth2.Token, 0)
	for _, tenantId := range a.conf.AuxiliaryTenantIDs {
		argsWithTenant := append(azArgs, "--tenant", tenantId)

		var token azureCliToken
		if err := azurecli.JSONUnmarshalAzCmd(&token, argsWithTenant...); err != nil {
			return nil, err
		}

		tokens = append(tokens, &oauth2.Token{
			AccessToken: token.AccessToken,
			TokenType:   token.TokenType,
		})
	}

	return tokens, nil
}

const (
	azureCliMinimumVersion   = "2.0.81"
	azureCliMsalVersion      = "2.30.0"
	azureCliNextMajorVersion = "3.0.0"
)

// azureCliConfig configures an AzureCliAuthorizer.
type azureCliConfig struct {
	Api environments.Api

	// TenantID is the required tenant ID for the primary token
	TenantID string

	// AuxiliaryTenantIDs is an optional list of tenant IDs for which to obtain additional tokens
	AuxiliaryTenantIDs []string
}

// newAzureCliConfig validates the supplied tenant ID and returns a new azureCliConfig.
func newAzureCliConfig(api environments.Api, tenantId string) (*azureCliConfig, error) {
	var err error

	// check az-cli version
	nextMajor := azureCliNextMajorVersion
	if err = azurecli.CheckAzVersion(azureCliMinimumVersion, &nextMajor); err != nil {
		return nil, err
	}

	// check tenant id
	tenantId, err = azurecli.CheckTenantID(tenantId)
	if err != nil {
		return nil, err
	}
	if tenantId == "" {
		return nil, errors.New("invalid tenantId or unable to determine tenantId")
	}

	return &azureCliConfig{Api: api, TenantID: tenantId}, nil
}

// TokenSource provides a source for obtaining access tokens using AzureCliAuthorizer.
func (c *azureCliConfig) TokenSource(ctx context.Context) Authorizer {
	// Cache access tokens internally to avoid unnecessary `az` invocations
	return NewCachedAuthorizer(&AzureCliAuthorizer{
		ctx:  ctx,
		conf: c,
	})
}

type azureCliToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresOn   string `json:"expiresOn"`
	Tenant      string `json:"tenant"`
	TokenType   string `json:"tokenType"`
}
