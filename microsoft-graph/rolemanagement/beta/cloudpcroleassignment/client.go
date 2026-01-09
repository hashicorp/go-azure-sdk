package cloudpcroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewCloudPCRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCRoleAssignmentClient: %+v", err)
	}

	return &CloudPCRoleAssignmentClient{
		Client: client,
	}, nil
}
