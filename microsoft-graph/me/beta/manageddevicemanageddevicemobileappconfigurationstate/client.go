package manageddevicemanageddevicemobileappconfigurationstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceManagedDeviceMobileAppConfigurationStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceManagedDeviceMobileAppConfigurationStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceManagedDeviceMobileAppConfigurationStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicemanageddevicemobileappconfigurationstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceManagedDeviceMobileAppConfigurationStateClient: %+v", err)
	}

	return &ManagedDeviceManagedDeviceMobileAppConfigurationStateClient{
		Client: client,
	}, nil
}
