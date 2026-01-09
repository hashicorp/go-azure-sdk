package manageddevicewindowsosimage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceWindowsOSImageClient struct {
	Client *msgraph.Client
}

func NewManagedDeviceWindowsOSImageClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDeviceWindowsOSImageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "manageddevicewindowsosimage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDeviceWindowsOSImageClient: %+v", err)
	}

	return &ManagedDeviceWindowsOSImageClient{
		Client: client,
	}, nil
}
