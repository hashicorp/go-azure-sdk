package rolemanagementexchangetransitiveroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementExchangeTransitiveRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementExchangeTransitiveRoleAssignmentAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementExchangeTransitiveRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementexchangetransitiveroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementExchangeTransitiveRoleAssignmentAppScopeClient: %+v", err)
	}

	return &RoleManagementExchangeTransitiveRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
