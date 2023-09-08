package usermanageddevicemanageddevicemobileappconfigurationstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceManagedDeviceMobileAppConfigurationStateClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceManagedDeviceMobileAppConfigurationStateClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceManagedDeviceMobileAppConfigurationStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicemanageddevicemobileappconfigurationstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceManagedDeviceMobileAppConfigurationStateClient: %+v", err)
	}

	return &UserManagedDeviceManagedDeviceMobileAppConfigurationStateClient{
		Client: client,
	}, nil
}
