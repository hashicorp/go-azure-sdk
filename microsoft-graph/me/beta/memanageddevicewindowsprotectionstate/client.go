package memanageddevicewindowsprotectionstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceWindowsProtectionStateClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceWindowsProtectionStateClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceWindowsProtectionStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddevicewindowsprotectionstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceWindowsProtectionStateClient: %+v", err)
	}

	return &MeManagedDeviceWindowsProtectionStateClient{
		Client: client,
	}, nil
}
