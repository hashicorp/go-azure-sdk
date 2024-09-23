package entitlementmanagementaccesspackageassignmenttargetconnectedorganization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentTargetConnectedOrganizationClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentTargetConnectedOrganizationClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentTargetConnectedOrganizationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmenttargetconnectedorganization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentTargetConnectedOrganizationClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentTargetConnectedOrganizationClient{
		Client: client,
	}, nil
}
