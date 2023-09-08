package usermanageddevicedevicecategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceDeviceCategoryClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceDeviceCategoryClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceDeviceCategoryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicedevicecategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceDeviceCategoryClient: %+v", err)
	}

	return &UserManagedDeviceDeviceCategoryClient{
		Client: client,
	}, nil
}
