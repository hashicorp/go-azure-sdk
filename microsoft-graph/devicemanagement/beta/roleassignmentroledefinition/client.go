package roleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "roleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &RoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
