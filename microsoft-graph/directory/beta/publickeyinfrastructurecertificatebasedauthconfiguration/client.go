package publickeyinfrastructurecertificatebasedauthconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublicKeyInfrastructureCertificateBasedAuthConfigurationClient struct {
	Client *msgraph.Client
}

func NewPublicKeyInfrastructureCertificateBasedAuthConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*PublicKeyInfrastructureCertificateBasedAuthConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "publickeyinfrastructurecertificatebasedauthconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PublicKeyInfrastructureCertificateBasedAuthConfigurationClient: %+v", err)
	}

	return &PublicKeyInfrastructureCertificateBasedAuthConfigurationClient{
		Client: client,
	}, nil
}
