package entitlementmanagementroleassignmentapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentApprovalClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentApprovalClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentApprovalClient{
		Client: client,
	}, nil
}
