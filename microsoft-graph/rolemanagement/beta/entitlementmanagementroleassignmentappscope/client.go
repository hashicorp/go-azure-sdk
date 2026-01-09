package entitlementmanagementroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentAppScopeClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
