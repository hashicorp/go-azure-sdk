package comanageddevicemanageddevicemobileappconfigurationstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceManagedDeviceMobileAppConfigurationStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicemanageddevicemobileappconfigurationstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient: %+v", err)
	}

	return &ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient{
		Client: client,
	}, nil
}
