package rolemanagementexchangeroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementExchangeRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementExchangeRoleAssignmentAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementExchangeRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementexchangeroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementExchangeRoleAssignmentAppScopeClient: %+v", err)
	}

	return &RoleManagementExchangeRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
