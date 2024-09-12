package entitlementmanagementroleeligibilityscheduleroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityscheduleroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleRoleDefinitionClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleRoleDefinitionClient{
		Client: client,
	}, nil
}
