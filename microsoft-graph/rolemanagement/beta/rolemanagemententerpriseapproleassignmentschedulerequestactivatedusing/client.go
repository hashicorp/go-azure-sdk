package rolemanagemententerpriseapproleassignmentschedulerequestactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleassignmentschedulerequestactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient{
		Client: client,
	}, nil
}
