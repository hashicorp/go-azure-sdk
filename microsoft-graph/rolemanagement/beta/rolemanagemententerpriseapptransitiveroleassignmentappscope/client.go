package rolemanagemententerpriseapptransitiveroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapptransitiveroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
