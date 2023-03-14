// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package test

import "os"

var (
	TenantId                      = os.Getenv("ARM_TENANT_ID")
	ClientId                      = os.Getenv("ARM_CLIENT_ID")
	ClientCertificate             = os.Getenv("ARM_CLIENT_CERTIFICATE")
	ClientCertificatePath         = os.Getenv("ARM_CLIENT_CERTIFICATE_PATH")
	ClientCertPassword            = os.Getenv("ARM_CLIENT_CERTIFICATE_PASSWORD")
	ClientSecret                  = os.Getenv("ARM_CLIENT_SECRET")
	Environment                   = envDefault("ARM_ENVIRONMENT", "global")
	GitHubTokenURL                = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	GitHubToken                   = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	IdToken                       = os.Getenv("ARM_OIDC_TOKEN")
	CustomManagedIdentityEndpoint = os.Getenv("ARM_MSI_ENDPOINT")
	ManagedIdentityToken          = os.Getenv("ARM_MSI_TOKEN")
)

func envDefault(key, def string) (ret string) {
	ret = os.Getenv(key)
	if ret == "" {
		ret = def
	}
	return
}
