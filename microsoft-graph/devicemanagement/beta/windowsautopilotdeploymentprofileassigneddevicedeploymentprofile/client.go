package windowsautopilotdeploymentprofileassigneddevicedeploymentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClient struct {
	Client *msgraph.Client
}

func NewWindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsautopilotdeploymentprofileassigneddevicedeploymentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClient: %+v", err)
	}

	return &WindowsAutopilotDeploymentProfileAssignedDeviceDeploymentProfileClient{
		Client: client,
	}, nil
}
