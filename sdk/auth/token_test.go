package auth_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"golang.org/x/oauth2"
)

var (
	testTokenIssued = time.Now()
	testTokenExpiry = time.Now().Add(3599 * time.Second)
)

var _ auth.Authorizer = &testAuthorizer{}

type testAuthorizer struct{}

func (*testAuthorizer) Token(_ context.Context, _ *http.Request) (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: testAccessToken(),
		TokenType:   "Bearer",
		Expiry:      testTokenExpiry,
	}, nil
}

func (*testAuthorizer) AuxiliaryTokens(_ context.Context, _ *http.Request) ([]*oauth2.Token, error) {
	return []*oauth2.Token{
		{
			AccessToken: testAccessToken(),
			TokenType:   "Bearer",
			Expiry:      testTokenExpiry,
		},
		{
			AccessToken: testAccessToken(),
			TokenType:   "Bearer",
			Expiry:      testTokenExpiry,
		},
		{
			AccessToken: testAccessToken(),
			TokenType:   "Bearer",
			Expiry:      testTokenExpiry,
		},
	}, nil
}

func (*testAuthorizer) ExpireTokens() error {
	return nil
}

func testAccessToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"appid": "10000000-2000-3000-4000-500000000000",
		"aud":   "https://not.azure.net",
		"exp":   testTokenExpiry.Unix(),
		"iat":   testTokenIssued.Unix(),
		"iss":   "Not Azure",
		"nbf":   testTokenIssued.Unix(),
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

func TestSetAuthHeaders(t *testing.T) {
	req := &http.Request{}
	authorizer := &testAuthorizer{}

	err := auth.SetAuthHeaders(context.Background(), req, authorizer)
	if err != nil {
		t.Fatalf("received error: %+v", err)
	}

	expected := regexp.MustCompile("^Bearer [a-zA-Z0-9-]+[.][a-zA-Z0-9-]+[.][a-zA-Z0-9-]+")
	if val := req.Header.Get("Authorization"); !expected.MatchString(val) {
		t.Fatalf("Authorization header mismatch, received: %q", val)
	}
	auxVals := strings.Split(req.Header.Get("X-Ms-Authorization-Auxiliary"), ",")
	if len(auxVals) != 3 {
		t.Fatalf("X-Ms-Authorization-Auxiliary mismatch, should contain 3 tokens, received: %q", strings.Join(auxVals, ","))
	}
	for i, auxVal := range auxVals {
		if !expected.MatchString(strings.TrimSpace(auxVal)) {
			t.Fatalf("Authorization header mismatch for token %d, received: %q", i, auxVal)
		}
	}
}
