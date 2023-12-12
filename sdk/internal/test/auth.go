// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"golang.org/x/oauth2"
)

var (
	TestTokenIssued = time.Now()
	TestTokenExpiry = time.Now().Add(3599 * time.Second)
)

var _ auth.Authorizer = &TestAuthorizer{}

type TestAuthorizer struct{}

func (*TestAuthorizer) Token(_ context.Context, _ *http.Request) (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: testAccessToken(),
		TokenType:   "Bearer",
		Expiry:      TestTokenExpiry,
	}, nil
}

func (*TestAuthorizer) AuxiliaryTokens(_ context.Context, _ *http.Request) ([]*oauth2.Token, error) {
	return []*oauth2.Token{
		{
			AccessToken: testAccessToken(),
			TokenType:   "Bearer",
			Expiry:      TestTokenExpiry,
		},
		{
			AccessToken: testAccessToken(),
			TokenType:   "Bearer",
			Expiry:      TestTokenExpiry,
		},
		{
			AccessToken: testAccessToken(),
			TokenType:   "Bearer",
			Expiry:      TestTokenExpiry,
		},
	}, nil
}

func (*TestAuthorizer) ExpireTokens() error {
	return nil
}

func testAccessToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"appid": "10000000-2000-3000-4000-500000000000",
		"aud":   "https://not.azure.net",
		"exp":   TestTokenExpiry.Unix(),
		"iat":   TestTokenIssued.Unix(),
		"iss":   "Not Azure",
		"nbf":   TestTokenIssued.Unix(),
		"oid":   "aaaaaa48-bb50-cc40-dd17-beeeeeeeeeef",
		"sub":   "12345678-1234-5678-90ab-1234567890ab",
		"tid":   "99999999-8888-7777-6666-555555543210",
	})

	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	ret, err := token.SignedString(key)
	if err != nil {
		panic(fmt.Sprintf("failed to sign test token: %+v", err))
	}

	return ret
}
