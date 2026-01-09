package entitlementmanagementroleeligibilityscheduleinstanceappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityscheduleinstanceappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient{
		Client: client,
	}, nil
}
