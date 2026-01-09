package microsofttunnelsitemicrosofttunnelconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelSiteMicrosoftTunnelConfigurationClient struct {
	Client *msgraph.Client
}

func NewMicrosoftTunnelSiteMicrosoftTunnelConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*MicrosoftTunnelSiteMicrosoftTunnelConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "microsofttunnelsitemicrosofttunnelconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MicrosoftTunnelSiteMicrosoftTunnelConfigurationClient: %+v", err)
	}

	return &MicrosoftTunnelSiteMicrosoftTunnelConfigurationClient{
		Client: client,
	}, nil
}
