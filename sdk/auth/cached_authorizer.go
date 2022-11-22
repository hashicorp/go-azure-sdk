package auth

import (
	"context"
	"golang.org/x/oauth2"
	"sync"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Authorizer = &CachedAuthorizer{}

// CachedAuthorizer caches a token until it expires, then acquires a new token from Source
type CachedAuthorizer struct {
	// Source contains the underlying Authorizer for obtaining tokens
	Source Authorizer

	mutex     sync.RWMutex
	token     *oauth2.Token
	auxTokens []*oauth2.Token
}

// Token returns the current token if it's still valid, else will acquire a new token
func (c *CachedAuthorizer) Token(ctx context.Context) (*oauth2.Token, error) {
	c.mutex.RLock()
	valid := c.token != nil && c.token.Valid()
	c.mutex.RUnlock()

	if !valid {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		var err error
		c.token, err = c.Source.Token(ctx)
		if err != nil {
			return nil, err
		}
	}

	return c.token, nil
}

// AuxiliaryTokens returns additional tokens for auxiliary tenant IDs, for use in multi-tenant scenarios
func (c *CachedAuthorizer) AuxiliaryTokens(ctx context.Context) ([]*oauth2.Token, error) {
	c.mutex.RLock()
	var valid bool
	for _, token := range c.auxTokens {
		valid = token != nil && token.Valid()
		if !valid {
			break
		}
	}
	c.mutex.RUnlock()

	if !valid {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		var err error
		c.auxTokens, err = c.Source.AuxiliaryTokens(ctx)
		if err != nil {
			return nil, err
		}
	}

	return c.auxTokens, nil
}

// NewCachedAuthorizer returns an Authorizer that caches an access token for the duration of its validity.
// If the cached token expires, a new one is acquired and cached.
func NewCachedAuthorizer(src Authorizer) Authorizer {
	return &CachedAuthorizer{
		Source: src,
	}
}
