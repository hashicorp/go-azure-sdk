package rolemanagemententerpriseapptransitiveroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapptransitiveroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
