package roledefinitionroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleDefinitionRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleDefinitionRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleDefinitionRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "roledefinitionroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleDefinitionRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &RoleDefinitionRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
