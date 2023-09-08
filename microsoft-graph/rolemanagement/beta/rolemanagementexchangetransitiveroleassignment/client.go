package rolemanagementexchangetransitiveroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementExchangeTransitiveRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementExchangeTransitiveRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*RoleManagementExchangeTransitiveRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementexchangetransitiveroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementExchangeTransitiveRoleAssignmentClient: %+v", err)
	}

	return &RoleManagementExchangeTransitiveRoleAssignmentClient{
		Client: client,
	}, nil
}
