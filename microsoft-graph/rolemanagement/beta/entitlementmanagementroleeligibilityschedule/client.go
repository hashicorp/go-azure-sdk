package entitlementmanagementroleeligibilityschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleClient{
		Client: client,
	}, nil
}
