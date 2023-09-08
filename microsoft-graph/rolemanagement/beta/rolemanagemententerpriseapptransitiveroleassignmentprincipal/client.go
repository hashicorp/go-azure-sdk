package rolemanagemententerpriseapptransitiveroleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagemententerpriseapptransitiveroleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClient: %+v", err)
	}

	return &RoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
