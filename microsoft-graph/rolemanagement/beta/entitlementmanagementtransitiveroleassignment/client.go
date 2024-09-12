package entitlementmanagementtransitiveroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementTransitiveRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementTransitiveRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementTransitiveRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementtransitiveroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementTransitiveRoleAssignmentClient: %+v", err)
	}

	return &EntitlementManagementTransitiveRoleAssignmentClient{
		Client: client,
	}, nil
}
