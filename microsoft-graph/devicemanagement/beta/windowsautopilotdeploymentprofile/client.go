package windowsautopilotdeploymentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfileClient struct {
	Client *msgraph.Client
}

func NewWindowsAutopilotDeploymentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsAutopilotDeploymentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsautopilotdeploymentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsAutopilotDeploymentProfileClient: %+v", err)
	}

	return &WindowsAutopilotDeploymentProfileClient{
		Client: client,
	}, nil
}
