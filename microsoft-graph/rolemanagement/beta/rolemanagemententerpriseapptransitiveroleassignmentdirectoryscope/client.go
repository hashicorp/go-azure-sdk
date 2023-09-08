package rolemanagemententerpriseapptransitiveroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapptransitiveroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
