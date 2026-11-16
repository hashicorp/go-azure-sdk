// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func clearManagedIdentityEnvVars(t *testing.T) {
	t.Helper()
	for _, k := range []string{
		"IDENTITY_ENDPOINT",
		"IDENTITY_HEADER",
		"IDENTITY_SERVER_THUMBPRINT",
		"MSI_ENDPOINT",
		"MSI_SECRET",
	} {
		t.Setenv(k, "")
	}
}

func TestManagedIdentityAuthorizer(t *testing.T) {
	clearManagedIdentityEnvVars(t)

	ctx := context.Background()
	env := environments.AzurePublic()

	origClient := auth.MetadataClient
	t.Cleanup(func() {
		auth.MetadataClient = origClient
	})
	auth.MetadataClient = &test.AzureADAccessTokenMockClient{
		Authorization: *env.Authorization,
	}

	opts := auth.ManagedIdentityAuthorizerOptions{
		Api:      env.MicrosoftGraph,
		ClientId: "11111111-0000-0000-0000-0000000000000000",
	}

	authorizer, err := auth.NewManagedIdentityAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewManagedIdentityAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestDetectManagedIdentitySource(t *testing.T) {
	testCases := []struct {
		name               string
		envVars            map[string]string
		expectedSource     auth.ManagedIdentitySource
		expectedEndpoint   string
		expectedAPIVersion string
		expectedHeaders    map[string][]string
	}{
		{
			name:               "default to IMDS when no env vars set",
			envVars:            map[string]string{},
			expectedSource:     auth.ManagedIdentitySourceIMDS,
			expectedEndpoint:   "http://169.254.169.254/metadata/identity/oauth2/token",
			expectedAPIVersion: "2018-02-01",
			expectedHeaders:    nil,
		},
		{
			name: "default to IMDS when only IDENTITY_ENDPOINT is set without IDENTITY_HEADER",
			envVars: map[string]string{
				"IDENTITY_ENDPOINT": "http://localhost:42356/msi/token",
			},
			expectedSource:     auth.ManagedIdentitySourceIMDS,
			expectedEndpoint:   "http://169.254.169.254/metadata/identity/oauth2/token",
			expectedAPIVersion: "2018-02-01",
			expectedHeaders:    nil,
		},
		{
			name: "default to IMDS when only IDENTITY_HEADER is set without IDENTITY_ENDPOINT",
			envVars: map[string]string{
				"IDENTITY_HEADER": "secret-header-value",
			},
			expectedSource:     auth.ManagedIdentitySourceIMDS,
			expectedEndpoint:   "http://169.254.169.254/metadata/identity/oauth2/token",
			expectedAPIVersion: "2018-02-01",
			expectedHeaders:    nil,
		},
		{
			name: "App Service / Container Apps / Functions",
			envVars: map[string]string{
				"IDENTITY_ENDPOINT": "http://localhost:42356/msi/token",
				"IDENTITY_HEADER":   "aca-secret-header",
			},
			expectedSource:     auth.ManagedIdentitySourceAppService,
			expectedEndpoint:   "http://localhost:42356/msi/token",
			expectedAPIVersion: "2019-08-01",
			expectedHeaders: map[string][]string{
				"X-IDENTITY-HEADER": {"aca-secret-header"},
			},
		},
		{
			name: "Service Fabric",
			envVars: map[string]string{
				"IDENTITY_ENDPOINT":          "http://localhost:2377/metadata/identity/oauth2/token",
				"IDENTITY_HEADER":            "sf-secret-header",
				"IDENTITY_SERVER_THUMBPRINT": "ABCDEF1234567890",
			},
			expectedSource:     auth.ManagedIdentitySourceServiceFabric,
			expectedEndpoint:   "http://localhost:2377/metadata/identity/oauth2/token",
			expectedAPIVersion: "2019-07-01-preview",
			expectedHeaders: map[string][]string{
				"Secret": {"sf-secret-header"},
			},
		},
		{
			name: "Azure ML",
			envVars: map[string]string{
				"MSI_ENDPOINT": "http://localhost:50000/msi/token",
				"MSI_SECRET":   "aml-secret-value",
			},
			expectedSource:     auth.ManagedIdentitySourceAzureML,
			expectedEndpoint:   "http://localhost:50000/msi/token",
			expectedAPIVersion: "2017-09-01",
			expectedHeaders: map[string][]string{
				"secret": {"aml-secret-value"},
			},
		},
		{
			name: "Cloud Shell",
			envVars: map[string]string{
				"MSI_ENDPOINT": "http://localhost:50342/oauth2/token",
			},
			expectedSource:     auth.ManagedIdentitySourceCloudShell,
			expectedEndpoint:   "http://localhost:50342/oauth2/token",
			expectedAPIVersion: "2018-02-01",
			expectedHeaders:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clearManagedIdentityEnvVars(t)
			for k, v := range tc.envVars {
				t.Setenv(k, v)
			}

			source, endpoint, apiVersion, headers := auth.DetectManagedIdentitySource()
			if source != tc.expectedSource {
				t.Errorf("expected source %q, got %q", tc.expectedSource, source)
			}
			if endpoint != tc.expectedEndpoint {
				t.Errorf("expected endpoint %q, got %q", tc.expectedEndpoint, endpoint)
			}
			if apiVersion != tc.expectedAPIVersion {
				t.Errorf("expected apiVersion %q, got %q", tc.expectedAPIVersion, apiVersion)
			}
			if !reflect.DeepEqual(headers, tc.expectedHeaders) {
				t.Errorf("expected headers %+v, got %+v", tc.expectedHeaders, headers)
			}
		})
	}
}

func TestNewManagedIdentityConfig(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()

	testCases := []struct {
		name               string
		envVars            map[string]string
		customEndpoint     string
		customAPIVersion   string
		customHeaders      map[string][]string
		expectedEndpoint   string
		expectedAPIVersion string
		expectedHeaders    map[string][]string
	}{
		{
			name:               "IMDS defaults when no env vars and no custom options",
			envVars:            map[string]string{},
			expectedEndpoint:   "http://169.254.169.254/metadata/identity/oauth2/token",
			expectedAPIVersion: "2018-02-01",
			expectedHeaders:    nil,
		},
		{
			name: "App Service / Container Apps defaults from environment",
			envVars: map[string]string{
				"IDENTITY_ENDPOINT": "http://localhost:42356/msi/token",
				"IDENTITY_HEADER":   "aca-secret-header",
			},
			expectedEndpoint:   "http://localhost:42356/msi/token",
			expectedAPIVersion: "2019-08-01",
			expectedHeaders: map[string][]string{
				"X-IDENTITY-HEADER": {"aca-secret-header"},
			},
		},
		{
			name: "CustomManagedIdentityEndpoint overrides IDENTITY_ENDPOINT without dropping IDENTITY_HEADER",
			envVars: map[string]string{
				"IDENTITY_ENDPOINT": "http://localhost:42356/msi/token",
				"IDENTITY_HEADER":   "aca-secret-header",
			},
			customEndpoint:     "http://custom-endpoint:8080/msi/token",
			expectedEndpoint:   "http://custom-endpoint:8080/msi/token",
			expectedAPIVersion: "2019-08-01",
			expectedHeaders: map[string][]string{
				"X-IDENTITY-HEADER": {"aca-secret-header"},
			},
		},
		{
			name: "Explicit CustomManagedIdentity* options override and merge with environment defaults",
			envVars: map[string]string{
				"IDENTITY_ENDPOINT": "http://localhost:42356/msi/token",
				"IDENTITY_HEADER":   "aca-secret-header",
			},
			customEndpoint:   "http://override-endpoint/token",
			customAPIVersion: "2021-01-01",
			customHeaders: map[string][]string{
				"X-IDENTITY-HEADER": {"overridden-header"},
				"X-Custom-Extra":    {"extra-val"},
			},
			expectedEndpoint:   "http://override-endpoint/token",
			expectedAPIVersion: "2021-01-01",
			expectedHeaders: map[string][]string{
				"X-IDENTITY-HEADER": {"overridden-header"},
				"X-Custom-Extra":    {"extra-val"},
			},
		},
		{
			name: "Service Fabric defaults from environment",
			envVars: map[string]string{
				"IDENTITY_ENDPOINT":          "http://localhost:2377/metadata/identity/oauth2/token",
				"IDENTITY_HEADER":            "sf-secret",
				"IDENTITY_SERVER_THUMBPRINT": "THUMBPRINT",
			},
			expectedEndpoint:   "http://localhost:2377/metadata/identity/oauth2/token",
			expectedAPIVersion: "2019-07-01-preview",
			expectedHeaders: map[string][]string{
				"Secret": {"sf-secret"},
			},
		},
		{
			name: "Azure ML defaults from environment",
			envVars: map[string]string{
				"MSI_ENDPOINT": "http://localhost:50000/msi/token",
				"MSI_SECRET":   "aml-secret",
			},
			expectedEndpoint:   "http://localhost:50000/msi/token",
			expectedAPIVersion: "2017-09-01",
			expectedHeaders: map[string][]string{
				"secret": {"aml-secret"},
			},
		},
		{
			name: "Cloud Shell defaults from environment",
			envVars: map[string]string{
				"MSI_ENDPOINT": "http://localhost:50342/oauth2/token",
			},
			expectedEndpoint:   "http://localhost:50342/oauth2/token",
			expectedAPIVersion: "2018-02-01",
			expectedHeaders:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clearManagedIdentityEnvVars(t)
			for k, v := range tc.envVars {
				t.Setenv(k, v)
			}

			authorizer, err := auth.NewManagedIdentityAuthorizer(ctx, auth.ManagedIdentityAuthorizerOptions{
				Api:                             env.ResourceManager,
				ClientId:                        "00000000-0000-0000-0000-000000000000",
				CustomManagedIdentityEndpoint:   tc.customEndpoint,
				CustomManagedIdentityAPIVersion: tc.customAPIVersion,
				CustomManagedIdentityHeaders:    tc.customHeaders,
			})
			if err != nil {
				t.Fatalf("NewManagedIdentityAuthorizer() unexpected error: %v", err)
			}
			cached, ok := authorizer.(*auth.CachedAuthorizer)
			if !ok {
				t.Fatalf("expected *auth.CachedAuthorizer, got %T", authorizer)
			}
			msiAuth, ok := cached.Source.(*auth.ManagedIdentityAuthorizer)
			if !ok {
				t.Fatalf("expected *auth.ManagedIdentityAuthorizer, got %T", cached.Source)
			}
			conf := msiAuth.Config()
			if conf.MsiEndpoint != tc.expectedEndpoint {
				t.Errorf("expected MsiEndpoint %q, got %q", tc.expectedEndpoint, conf.MsiEndpoint)
			}
			if conf.MsiApiVersion != tc.expectedAPIVersion {
				t.Errorf("expected MsiApiVersion %q, got %q", tc.expectedAPIVersion, conf.MsiApiVersion)
			}
			if !reflect.DeepEqual(conf.MsiHeaders, tc.expectedHeaders) {
				t.Errorf("expected MsiHeaders %+v, got %+v", tc.expectedHeaders, conf.MsiHeaders)
			}
		})
	}
}

type mockHTTPServerClient struct {
	handler http.HandlerFunc
}

func (m *mockHTTPServerClient) Do(req *http.Request) (*http.Response, error) {
	rec := httptest.NewRecorder()
	m.handler(rec, req)
	resp := rec.Result()
	resp.Request = req
	return resp, nil
}

func TestManagedIdentityAuthorizer_EnvironmentDetection(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()
	clientID := "11111111-2222-3333-4444-555555555555"

	origClient := auth.MetadataClient
	t.Cleanup(func() {
		auth.MetadataClient = origClient
	})

	serverURL := "http://127.0.0.1:42356"

	testCases := []struct {
		name               string
		envVars            map[string]string
		expectedAPIVersion string
		expectedClientKey  string
		expectedHeaderKey  string
		expectedHeaderVal  string
	}{
		{
			name: "App Service / Container Apps",
			envVars: map[string]string{
				"IDENTITY_ENDPOINT": serverURL + "/msi/token",
				"IDENTITY_HEADER":   "aca-test-secret",
			},
			expectedAPIVersion: "2019-08-01",
			expectedClientKey:  "client_id",
			expectedHeaderKey:  "X-Identity-Header",
			expectedHeaderVal:  "aca-test-secret",
		},
		{
			name: "Service Fabric",
			envVars: map[string]string{
				"IDENTITY_ENDPOINT":          serverURL + "/msi/token",
				"IDENTITY_HEADER":            "sf-test-secret",
				"IDENTITY_SERVER_THUMBPRINT": "THUMB123",
			},
			expectedAPIVersion: "2019-07-01-preview",
			expectedClientKey:  "client_id",
			expectedHeaderKey:  "Secret",
			expectedHeaderVal:  "sf-test-secret",
		},
		{
			name: "Azure ML",
			envVars: map[string]string{
				"MSI_ENDPOINT": serverURL + "/msi/token",
				"MSI_SECRET":   "aml-test-secret",
			},
			expectedAPIVersion: "2017-09-01",
			expectedClientKey:  "clientid",
			expectedHeaderKey:  "Secret",
			expectedHeaderVal:  "aml-test-secret",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var capturedReq *http.Request
			auth.MetadataClient = &mockHTTPServerClient{
				handler: func(w http.ResponseWriter, r *http.Request) {
					capturedReq = r
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprintf(w, `{"access_token":%q,"expires_in":"3600","token_type":"Bearer"}`, test.DummyAccessToken)
				},
			}

			clearManagedIdentityEnvVars(t)
			for k, v := range tc.envVars {
				t.Setenv(k, v)
			}

			authorizer, err := auth.NewManagedIdentityAuthorizer(ctx, auth.ManagedIdentityAuthorizerOptions{
				Api:      env.MicrosoftGraph,
				ClientId: clientID,
			})
			if err != nil {
				t.Fatalf("NewManagedIdentityAuthorizer(): %v", err)
			}

			token, err := testObtainAccessToken(ctx, authorizer)
			if err != nil {
				t.Fatalf("testObtainAccessToken(): %v", err)
			}
			if token.AccessToken != test.DummyAccessToken {
				t.Fatalf("unexpected token: %+v", token)
			}

			if capturedReq == nil {
				t.Fatal("expected HTTP request to hit mock server, but none was received")
			}
			if capturedReq.URL.Host != "127.0.0.1:42356" {
				t.Errorf("expected host 127.0.0.1:42356, got %q", capturedReq.URL.Host)
			}
			if capturedReq.URL.Path != "/msi/token" {
				t.Errorf("expected path /msi/token, got %q", capturedReq.URL.Path)
			}
			if got := capturedReq.URL.Query().Get("api-version"); got != tc.expectedAPIVersion {
				t.Errorf("expected api-version %q, got %q", tc.expectedAPIVersion, got)
			}
			if got := capturedReq.URL.Query().Get(tc.expectedClientKey); got != clientID {
				t.Errorf("expected %s=%q, got %q", tc.expectedClientKey, clientID, got)
			}
			if got := capturedReq.Header.Get("Metadata"); got != "true" {
				t.Errorf("expected Metadata header 'true', got %q", got)
			}
			if got := capturedReq.Header.Get(tc.expectedHeaderKey); got != tc.expectedHeaderVal {
				t.Errorf("expected header %s=%q, got %q", tc.expectedHeaderKey, tc.expectedHeaderVal, got)
			}
		})
	}
}

func TestAccManagedIdentityAuthorizer(t *testing.T) {
	test.AccTest(t)

	if test.CustomManagedIdentityEndpoint == "" {
		t.Skip("test.CustomManagedIdentityEndpoint was empty")
	}

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.ManagedIdentityAuthorizerOptions{
		Api:                           env.MicrosoftGraph,
		ClientId:                      test.ClientId,
		CustomManagedIdentityEndpoint: test.CustomManagedIdentityEndpoint,
	}

	authorizer, err := auth.NewManagedIdentityAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewManagedIdentityAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}
