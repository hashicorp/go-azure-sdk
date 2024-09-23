package entitlementmanagementroleeligibilityscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilitySchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilitySchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilitySchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilitySchedulePrincipalClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilitySchedulePrincipalClient{
		Client: client,
	}, nil
}
