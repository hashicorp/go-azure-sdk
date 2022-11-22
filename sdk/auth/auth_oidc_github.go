package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GitHubOIDCAuthorizerOptions struct {
	// TODO: document these

	Api                 environments.Api
	AuxiliaryTenantIds  []string
	ClientId            string
	Environment         environments.Environment
	TenantId            string
	IdTokenRequestUrl   string
	IdTokenRequestToken string
}

// NewGitHubOIDCAuthorizer returns an authorizer which acquires a client assertion from a GitHub endpoint, then uses client assertion authentication to obtain an access token.
func NewGitHubOIDCAuthorizer(ctx context.Context, options GitHubOIDCAuthorizerOptions) (Authorizer, error) {
	conf := gitHubOIDCConfig{
		Environment:         options.Environment,
		TenantID:            options.TenantId,
		AuxiliaryTenantIDs:  options.AuxiliaryTenantIds,
		ClientID:            options.ClientId,
		IDTokenRequestURL:   options.IdTokenRequestUrl,
		IDTokenRequestToken: options.IdTokenRequestToken,
		Scopes:              []string{options.Api.DefaultScope()},
	}

	return conf.TokenSource(ctx), nil
}

type GitHubOIDCAuthorizer struct {
	ctx  context.Context
	conf *gitHubOIDCConfig
}

func (a *GitHubOIDCAuthorizer) githubAssertion() (*string, error) {
	req, err := http.NewRequestWithContext(a.ctx, http.MethodGet, a.conf.IDTokenRequestURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("githubAssertion: failed to build request")
	}

	query, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("githubAssertion: cannot parse URL query")
	}

	if query.Get("audience") == "" {
		query.Set("audience", "api://AzureADTokenExchange")
		req.URL.RawQuery = query.Encode()
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.conf.IDTokenRequestToken))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("githubAssertion: cannot request token: %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("githubAssertion: cannot parse response: %v", err)
	}

	if c := resp.StatusCode; c < 200 || c > 299 {
		return nil, fmt.Errorf("githubAssertion: received HTTP status %d with response: %s", resp.StatusCode, body)
	}

	var tokenRes struct {
		Count *int    `json:"count"`
		Value *string `json:"value"`
	}
	if err := json.Unmarshal(body, &tokenRes); err != nil {
		return nil, fmt.Errorf("githubAssertion: cannot unmarshal response: %v", err)
	}

	return tokenRes.Value, nil
}

func (a *GitHubOIDCAuthorizer) tokenSource() (Authorizer, error) {
	assertion, err := a.githubAssertion()
	if err != nil {
		return nil, err
	}
	if assertion == nil {
		return nil, fmt.Errorf("GitHubOIDCAuthorizer: nil JWT assertion received from GitHub")
	}

	conf := clientCredentialsConfig{
		Environment:        a.conf.Environment,
		TenantID:           a.conf.TenantID,
		AuxiliaryTenantIDs: a.conf.AuxiliaryTenantIDs,
		ClientID:           a.conf.ClientID,
		FederatedAssertion: *assertion,
		Scopes:             a.conf.Scopes,
		TokenURL:           a.conf.TokenURL,
		TokenVersion:       TokenVersion2,
		Audience:           a.conf.Audience,
	}

	source := conf.TokenSource(a.ctx, clientCredentialsAssertionType)
	if source == nil {
		return nil, fmt.Errorf("GitHubOIDCAuthorizer: nil Authorizer returned from clientCredentialsConfig")
	}

	return source, nil
}

func (a *GitHubOIDCAuthorizer) Token() (*oauth2.Token, error) {
	source, err := a.tokenSource()
	if err != nil {
		return nil, err
	}
	return source.Token()
}

func (a *GitHubOIDCAuthorizer) AuxiliaryTokens() ([]*oauth2.Token, error) {
	source, err := a.tokenSource()
	if err != nil {
		return nil, err
	}
	return source.AuxiliaryTokens()
}

type gitHubOIDCConfig struct {
	// Environment is the national cloud environment to use
	Environment environments.Environment

	// TenantID is the required tenant ID for the primary token
	TenantID string

	// AuxiliaryTenantIDs is an optional list of tenant IDs for which to obtain additional tokens
	AuxiliaryTenantIDs []string

	// ClientID is the application's ID.
	ClientID string

	// IDTokenRequestURL is the URL for GitHub's OIDC provider.
	IDTokenRequestURL string

	// IDTokenRequestToken is the bearer token for the request to the OIDC provider.
	IDTokenRequestToken string

	// Scopes specifies a list of requested permission scopes (used for v2 tokens)
	Scopes []string

	// TokenURL is the clientCredentialsToken endpoint, which overrides the default endpoint constructed from a tenant ID
	TokenURL string

	// Audience optionally specifies the intended audience of the
	// request.  If empty, the value of TokenURL is used as the
	// intended audience.
	Audience string
}

func (c *gitHubOIDCConfig) TokenSource(ctx context.Context) Authorizer {
	return NewCachedAuthorizer(&GitHubOIDCAuthorizer{ctx, c})
}