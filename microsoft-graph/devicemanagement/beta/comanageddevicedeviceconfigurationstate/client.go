package comanageddevicedeviceconfigurationstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceDeviceConfigurationStateClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceDeviceConfigurationStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceDeviceConfigurationStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicedeviceconfigurationstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceDeviceConfigurationStateClient: %+v", err)
	}

	return &ComanagedDeviceDeviceConfigurationStateClient{
		Client: client,
	}, nil
}
