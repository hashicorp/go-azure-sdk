package windowsautopilotdeploymentprofileassigneddeviceintendeddeploymentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClient struct {
	Client *msgraph.Client
}

func NewWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsautopilotdeploymentprofileassigneddeviceintendeddeploymentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClient: %+v", err)
	}

	return &WindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClient{
		Client: client,
	}, nil
}
