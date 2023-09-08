package usermanageddevicedeviceconfigurationstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceDeviceConfigurationStateClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceDeviceConfigurationStateClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceDeviceConfigurationStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicedeviceconfigurationstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceDeviceConfigurationStateClient: %+v", err)
	}

	return &UserManagedDeviceDeviceConfigurationStateClient{
		Client: client,
	}, nil
}
