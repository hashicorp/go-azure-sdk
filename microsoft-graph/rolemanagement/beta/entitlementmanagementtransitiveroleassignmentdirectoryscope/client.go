package entitlementmanagementtransitiveroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementTransitiveRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementTransitiveRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementtransitiveroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementTransitiveRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &EntitlementManagementTransitiveRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
