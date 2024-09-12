package manageddevicemobileappconfigurationstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceMobileAppConfigurationStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceMobileAppConfigurationStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicemobileappconfigurationstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceMobileAppConfigurationStateClient: %+v", err)
	}

	return &ManagedDeviceMobileAppConfigurationStateClient{
		Client: client,
	}, nil
}
