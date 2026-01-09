package entitlementmanagementroleeligibilityscheduleappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleAppScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityscheduleappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleAppScopeClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleAppScopeClient{
		Client: client,
	}, nil
}
