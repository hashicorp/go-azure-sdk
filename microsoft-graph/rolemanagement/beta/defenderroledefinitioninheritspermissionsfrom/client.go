package defenderroledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewDefenderRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderroledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &DefenderRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}
