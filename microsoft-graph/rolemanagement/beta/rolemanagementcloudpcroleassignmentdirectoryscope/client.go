package rolemanagementcloudpcroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementCloudPCRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewRoleManagementCloudPCRoleAssignmentDirectoryScopeClientWithBaseURI(api sdkEnv.Api) (*RoleManagementCloudPCRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementcloudpcroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementCloudPCRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &RoleManagementCloudPCRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
