package comanageddevicewindowsprotectionstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceWindowsProtectionStateClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceWindowsProtectionStateClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceWindowsProtectionStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicewindowsprotectionstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceWindowsProtectionStateClient: %+v", err)
	}

	return &ComanagedDeviceWindowsProtectionStateClient{
		Client: client,
	}, nil
}
