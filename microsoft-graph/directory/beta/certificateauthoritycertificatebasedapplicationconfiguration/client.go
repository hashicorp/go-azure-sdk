package certificateauthoritycertificatebasedapplicationconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthorityCertificateBasedApplicationConfigurationClient struct {
	Client *msgraph.Client
}

func NewCertificateAuthorityCertificateBasedApplicationConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*CertificateAuthorityCertificateBasedApplicationConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "certificateauthoritycertificatebasedapplicationconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CertificateAuthorityCertificateBasedApplicationConfigurationClient: %+v", err)
	}

	return &CertificateAuthorityCertificateBasedApplicationConfigurationClient{
		Client: client,
	}, nil
}
