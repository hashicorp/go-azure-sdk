package defenderroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewDefenderRoleAssignmentAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderRoleAssignmentAppScopeClient: %+v", err)
	}

	return &DefenderRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
