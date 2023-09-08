package rolemanagementcloudpcroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementCloudPCRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleManagementCloudPCRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*RoleManagementCloudPCRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "rolemanagementcloudpcroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementCloudPCRoleAssignmentClient: %+v", err)
	}

	return &RoleManagementCloudPCRoleAssignmentClient{
		Client: client,
	}, nil
}
