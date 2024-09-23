package entitlementmanagementroleeligibilityscheduledirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityscheduledirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleDirectoryScopeClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleDirectoryScopeClient{
		Client: client,
	}, nil
}
