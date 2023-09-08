package rolemanagementcloudpcroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementCloudPCRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementCloudPCRoleAssignmentAppScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementCloudPCRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementcloudpcroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementCloudPCRoleAssignmentAppScopeClient: %+v", err)
	}

	return &RoleManagementCloudPCRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
