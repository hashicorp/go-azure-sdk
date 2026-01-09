package grouppolicyconfigurationassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationAssignmentClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyConfigurationAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyConfigurationAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyconfigurationassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyConfigurationAssignmentClient: %+v", err)
	}

	return &GroupPolicyConfigurationAssignmentClient{
		Client: client,
	}, nil
}
