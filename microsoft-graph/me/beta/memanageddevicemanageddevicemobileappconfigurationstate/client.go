package memanageddevicemanageddevicemobileappconfigurationstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceManagedDeviceMobileAppConfigurationStateClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceManagedDeviceMobileAppConfigurationStateClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceManagedDeviceMobileAppConfigurationStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddevicemanageddevicemobileappconfigurationstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceManagedDeviceMobileAppConfigurationStateClient: %+v", err)
	}

	return &MeManagedDeviceManagedDeviceMobileAppConfigurationStateClient{
		Client: client,
	}, nil
}
