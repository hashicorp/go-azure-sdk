package embeddedsimactivationcodepooldevicestate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMActivationCodePoolDeviceStateClient struct {
	Client *msgraph.Client
}

func NewEmbeddedSIMActivationCodePoolDeviceStateClientWithBaseURI(sdkApi sdkEnv.Api) (*EmbeddedSIMActivationCodePoolDeviceStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "embeddedsimactivationcodepooldevicestate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmbeddedSIMActivationCodePoolDeviceStateClient: %+v", err)
	}

	return &EmbeddedSIMActivationCodePoolDeviceStateClient{
		Client: client,
	}, nil
}
