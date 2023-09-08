package rolemanagemententerpriseapproleassignmentscheduleinstanceappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleassignmentscheduleinstanceappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient{
		Client: client,
	}, nil
}
