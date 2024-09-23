package entitlementmanagementroleeligibilityschedulerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleRequestClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityschedulerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleRequestClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleRequestClient{
		Client: client,
	}, nil
}
