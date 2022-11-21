package auth_test

import (
	"os"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var (
	tenantId              = os.Getenv("ARM_TENANT_ID")
	clientId              = os.Getenv("ARM_CLIENT_ID")
	clientCertificate     = os.Getenv("ARM_CLIENT_CERTIFICATE")
	clientCertificatePath = os.Getenv("ARM_CLIENT_CERTIFICATE_PATH")
	clientCertPassword    = os.Getenv("ARM_CLIENT_CERTIFICATE_PASSWORD")
	clientSecret          = os.Getenv("ARM_CLIENT_SECRET")
	environment           = envDefault("ARM_ENVIRONMENT", "global")
	gitHubTokenURL        = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	gitHubToken           = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	idToken               = os.Getenv("ARM_OIDC_TOKEN")
	msiEndpoint           = os.Getenv("ARM_MSI_ENDPOINT")
	msiToken              = os.Getenv("ARM_MSI_TOKEN")
)

func envDefault(key, def string) (ret string) {
	ret = os.Getenv(key)
	if ret == "" {
		ret = def
	}
	return
}
