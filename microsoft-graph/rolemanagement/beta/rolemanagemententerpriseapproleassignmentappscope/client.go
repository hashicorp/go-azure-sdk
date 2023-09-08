package rolemanagemententerpriseapproleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleAssignmentAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleAssignmentAppScopeClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
