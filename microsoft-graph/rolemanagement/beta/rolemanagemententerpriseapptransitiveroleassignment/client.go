package rolemanagemententerpriseapptransitiveroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppTransitiveRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppTransitiveRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppTransitiveRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapptransitiveroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppTransitiveRoleAssignmentClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppTransitiveRoleAssignmentClient{
		Client: client,
	}, nil
}
