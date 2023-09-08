package rolemanagemententerpriseapproleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleAssignmentClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleAssignmentClient{
		Client: client,
	}, nil
}
