package cloudpcroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewCloudPCRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &CloudPCRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
