package defenderroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDefenderRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &DefenderRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
