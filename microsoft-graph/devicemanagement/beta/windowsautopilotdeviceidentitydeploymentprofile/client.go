package windowsautopilotdeviceidentitydeploymentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityDeploymentProfileClient struct {
	Client *msgraph.Client
}

func NewWindowsAutopilotDeviceIdentityDeploymentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsAutopilotDeviceIdentityDeploymentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsautopilotdeviceidentitydeploymentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsAutopilotDeviceIdentityDeploymentProfileClient: %+v", err)
	}

	return &WindowsAutopilotDeviceIdentityDeploymentProfileClient{
		Client: client,
	}, nil
}
