package rolemanagementexchangetransitiveroleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementExchangeTransitiveRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementExchangeTransitiveRoleAssignmentPrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementExchangeTransitiveRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementexchangetransitiveroleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementExchangeTransitiveRoleAssignmentPrincipalClient: %+v", err)
	}

	return &RoleManagementExchangeTransitiveRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
