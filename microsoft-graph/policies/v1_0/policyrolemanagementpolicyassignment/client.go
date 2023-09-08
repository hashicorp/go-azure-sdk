package policyrolemanagementpolicyassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyRoleManagementPolicyAssignmentClient struct {
	Client *msgraph.Client
}

func NewPolicyRoleManagementPolicyAssignmentClientWithBaseURI(api sdkEnv.Api) (*PolicyRoleManagementPolicyAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyrolemanagementpolicyassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyRoleManagementPolicyAssignmentClient: %+v", err)
	}

	return &PolicyRoleManagementPolicyAssignmentClient{
		Client: client,
	}, nil
}
