// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package test

import (
	"os"
	"strings"
)

var (
	AuxiliaryTenantIds            = strings.Split(os.Getenv("ARM_AUXILIARY_TENANT_IDS"), ";")
	ClientCertPassword            = os.Getenv("ARM_CLIENT_CERTIFICATE_PASSWORD")
	ClientCertificate             = os.Getenv("ARM_CLIENT_CERTIFICATE")
	ClientCertificatePath         = os.Getenv("ARM_CLIENT_CERTIFICATE_PATH")
	ClientId                      = os.Getenv("ARM_CLIENT_ID")
	ClientSecret                  = os.Getenv("ARM_CLIENT_SECRET")
	CustomManagedIdentityEndpoint = os.Getenv("ARM_MSI_ENDPOINT")
	Environment                   = envDefault("ARM_ENVIRONMENT", "global")
	OIDCRequestToken              = anyOfEnvVar("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "SYSTEM_ACCESSTOKEN")
	OIDCRequestURL                = anyOfEnvVar("ACTIONS_ID_TOKEN_REQUEST_URL", "SYSTEM_OIDCREQUESTURI")
	ADOServiceConnectionId        = os.Getenv("ARM_ADO_PIPELINE_SERVICE_CONNECTION_ID")
	IdToken                       = os.Getenv("ARM_OIDC_TOKEN")
	SubscriptionId                = os.Getenv("ARM_SUBSCRIPTION_ID")
	TenantId                      = os.Getenv("ARM_TENANT_ID")
)

func anyOfEnvVar(vs ...string) string {
	for _, v := range vs {
		if vv := os.Getenv(v); vv != "" {
			return vv
		}
	}
	return ""
}

func envDefault(key, def string) (ret string) {
	ret = os.Getenv(key)
	if ret == "" {
		ret = def
	}
	return
}
