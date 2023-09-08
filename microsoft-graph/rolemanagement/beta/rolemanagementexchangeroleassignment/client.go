package rolemanagementexchangeroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementExchangeRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementExchangeRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*RoleManagementExchangeRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementexchangeroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementExchangeRoleAssignmentClient: %+v", err)
	}

	return &RoleManagementExchangeRoleAssignmentClient{
		Client: client,
	}, nil
}
