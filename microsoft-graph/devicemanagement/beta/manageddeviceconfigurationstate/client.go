package manageddeviceconfigurationstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceConfigurationStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceConfigurationStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceConfigurationStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddeviceconfigurationstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceConfigurationStateClient: %+v", err)
	}

	return &ManagedDeviceConfigurationStateClient{
		Client: client,
	}, nil
}
