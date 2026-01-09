package manageddevicewindowsprotectionstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceWindowsProtectionStateClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceWindowsProtectionStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceWindowsProtectionStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicewindowsprotectionstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceWindowsProtectionStateClient: %+v", err)
	}

	return &ManagedDeviceWindowsProtectionStateClient{
		Client: client,
	}, nil
}
