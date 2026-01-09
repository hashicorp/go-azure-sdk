package manageddevicedeviceconfigurationstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceConfigurationStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceDeviceConfigurationStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceDeviceConfigurationStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicedeviceconfigurationstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceDeviceConfigurationStateClient: %+v", err)
	}

	return &ManagedDeviceDeviceConfigurationStateClient{
		Client: client,
	}, nil
}
