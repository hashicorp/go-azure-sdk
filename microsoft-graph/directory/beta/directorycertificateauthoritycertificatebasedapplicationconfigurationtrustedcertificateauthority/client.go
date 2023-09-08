package directorycertificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient struct {
	Client *msgraph.Client
}

func NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClientWithBaseURI(api sdkEnv.Api) (*DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directorycertificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient: %+v", err)
	}

	return &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient{
		Client: client,
	}, nil
}
