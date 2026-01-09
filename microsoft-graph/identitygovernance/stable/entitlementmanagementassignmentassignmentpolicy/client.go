package entitlementmanagementassignmentassignmentpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentAssignmentPolicyClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentAssignmentPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentAssignmentPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentassignmentpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentAssignmentPolicyClient: %+v", err)
	}

	return &EntitlementManagementAssignmentAssignmentPolicyClient{
		Client: client,
	}, nil
}
