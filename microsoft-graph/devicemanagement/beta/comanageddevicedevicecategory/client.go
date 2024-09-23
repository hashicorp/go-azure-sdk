package comanageddevicedevicecategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceDeviceCategoryClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceDeviceCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceDeviceCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicedevicecategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceDeviceCategoryClient: %+v", err)
	}

	return &ComanagedDeviceDeviceCategoryClient{
		Client: client,
	}, nil
}
