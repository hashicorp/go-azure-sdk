package certificateauthoritymutualtlsoauthconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthorityMutualTlsOauthConfigurationClient struct {
	Client *msgraph.Client
}

func NewCertificateAuthorityMutualTlsOauthConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*CertificateAuthorityMutualTlsOauthConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "certificateauthoritymutualtlsoauthconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CertificateAuthorityMutualTlsOauthConfigurationClient: %+v", err)
	}

	return &CertificateAuthorityMutualTlsOauthConfigurationClient{
		Client: client,
	}, nil
}
