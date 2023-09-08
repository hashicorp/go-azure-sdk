package memanageddevicedevicecategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceDeviceCategoryClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceDeviceCategoryClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceDeviceCategoryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddevicedevicecategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceDeviceCategoryClient: %+v", err)
	}

	return &MeManagedDeviceDeviceCategoryClient{
		Client: client,
	}, nil
}
