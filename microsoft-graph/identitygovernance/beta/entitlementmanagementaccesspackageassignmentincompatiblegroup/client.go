package entitlementmanagementaccesspackageassignmentincompatiblegroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentIncompatibleGroupClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentIncompatibleGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentIncompatibleGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentincompatiblegroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentIncompatibleGroupClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentIncompatibleGroupClient{
		Client: client,
	}, nil
}
