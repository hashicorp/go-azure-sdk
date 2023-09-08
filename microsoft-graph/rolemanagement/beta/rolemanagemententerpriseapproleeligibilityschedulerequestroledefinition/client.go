package rolemanagemententerpriseapproleeligibilityschedulerequestroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapproleeligibilityschedulerequestroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient{
		Client: client,
	}, nil
}
