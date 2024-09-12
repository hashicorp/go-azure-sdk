package windowsqualityupdatepolicyassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdatePolicyAssignmentClient struct {
	Client *msgraph.Client
}

func NewWindowsQualityUpdatePolicyAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsQualityUpdatePolicyAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsqualityupdatepolicyassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsQualityUpdatePolicyAssignmentClient: %+v", err)
	}

	return &WindowsQualityUpdatePolicyAssignmentClient{
		Client: client,
	}, nil
}
