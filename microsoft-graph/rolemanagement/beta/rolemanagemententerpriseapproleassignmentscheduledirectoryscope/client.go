package rolemanagemententerpriseapproleassignmentscheduledirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleassignmentscheduledirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClient{
		Client: client,
	}, nil
}
