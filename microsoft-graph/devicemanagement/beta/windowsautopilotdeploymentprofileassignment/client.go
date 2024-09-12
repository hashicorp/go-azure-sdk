package windowsautopilotdeploymentprofileassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfileAssignmentClient struct {
	Client *msgraph.Client
}

func NewWindowsAutopilotDeploymentProfileAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsAutopilotDeploymentProfileAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsautopilotdeploymentprofileassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsAutopilotDeploymentProfileAssignmentClient: %+v", err)
	}

	return &WindowsAutopilotDeploymentProfileAssignmentClient{
		Client: client,
	}, nil
}
