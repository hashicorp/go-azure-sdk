package entitlementmanagementroleeligibilityscheduleinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleInstanceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityscheduleinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleInstanceClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleInstanceClient{
		Client: client,
	}, nil
}
