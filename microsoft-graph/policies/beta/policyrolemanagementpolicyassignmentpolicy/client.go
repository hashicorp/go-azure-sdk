package policyrolemanagementpolicyassignmentpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyRoleManagementPolicyAssignmentPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyRoleManagementPolicyAssignmentPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyRoleManagementPolicyAssignmentPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyrolemanagementpolicyassignmentpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyRoleManagementPolicyAssignmentPolicyClient: %+v", err)
	}

	return &PolicyRoleManagementPolicyAssignmentPolicyClient{
		Client: client,
	}, nil
}
