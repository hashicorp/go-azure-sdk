package certificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient struct {
	Client *msgraph.Client
}

func NewCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClientWithBaseURI(sdkApi sdkEnv.Api) (*CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "certificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient: %+v", err)
	}

	return &CertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityClient{
		Client: client,
	}, nil
}
