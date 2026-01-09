package windowsautopilotdeviceidentity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityClient struct {
	Client *msgraph.Client
}

func NewWindowsAutopilotDeviceIdentityClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsAutopilotDeviceIdentityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsautopilotdeviceidentity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsAutopilotDeviceIdentityClient: %+v", err)
	}

	return &WindowsAutopilotDeviceIdentityClient{
		Client: client,
	}, nil
}
