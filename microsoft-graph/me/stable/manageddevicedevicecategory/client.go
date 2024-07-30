package manageddevicedevicecategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceCategoryClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceDeviceCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceDeviceCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicedevicecategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceDeviceCategoryClient: %+v", err)
	}

	return &ManagedDeviceDeviceCategoryClient{
		Client: client,
	}, nil
}
