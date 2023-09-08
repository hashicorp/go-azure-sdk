package rolemanagementcloudpcroleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementCloudPCRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewRoleManagementCloudPCRoleAssignmentPrincipalClientWithBaseURI(api sdkEnv.Api) (*RoleManagementCloudPCRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementcloudpcroleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementCloudPCRoleAssignmentPrincipalClient: %+v", err)
	}

	return &RoleManagementCloudPCRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
