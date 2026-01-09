package defenderroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewDefenderRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &DefenderRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
