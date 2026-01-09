package entitlementmanagementroleeligibilityschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
