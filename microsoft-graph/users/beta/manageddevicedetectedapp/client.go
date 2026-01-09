package manageddevicedetectedapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDetectedAppClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceDetectedAppClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceDetectedAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicedetectedapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceDetectedAppClient: %+v", err)
	}

	return &ManagedDeviceDetectedAppClient{
		Client: client,
	}, nil
}
