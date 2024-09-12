package microsofttunnelsiteconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelSiteConfigurationClient struct {
	Client *msgraph.Client
}

func NewMicrosoftTunnelSiteConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*MicrosoftTunnelSiteConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "microsofttunnelsiteconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MicrosoftTunnelSiteConfigurationClient: %+v", err)
	}

	return &MicrosoftTunnelSiteConfigurationClient{
		Client: client,
	}, nil
}
