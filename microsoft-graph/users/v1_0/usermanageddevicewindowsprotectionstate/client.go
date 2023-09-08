package usermanageddevicewindowsprotectionstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceWindowsProtectionStateClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceWindowsProtectionStateClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceWindowsProtectionStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicewindowsprotectionstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceWindowsProtectionStateClient: %+v", err)
	}

	return &UserManagedDeviceWindowsProtectionStateClient{
		Client: client,
	}, nil
}
