package microsofttunnelconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelConfigurationClient struct {
	Client *msgraph.Client
}

func NewMicrosoftTunnelConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*MicrosoftTunnelConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "microsofttunnelconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MicrosoftTunnelConfigurationClient: %+v", err)
	}

	return &MicrosoftTunnelConfigurationClient{
		Client: client,
	}, nil
}
