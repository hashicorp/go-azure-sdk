package windowsautopilotdeviceidentityintendeddeploymentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityIntendedDeploymentProfileClient struct {
	Client *msgraph.Client
}

func NewWindowsAutopilotDeviceIdentityIntendedDeploymentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsAutopilotDeviceIdentityIntendedDeploymentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsautopilotdeviceidentityintendeddeploymentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsAutopilotDeviceIdentityIntendedDeploymentProfileClient: %+v", err)
	}

	return &WindowsAutopilotDeviceIdentityIntendedDeploymentProfileClient{
		Client: client,
	}, nil
}
