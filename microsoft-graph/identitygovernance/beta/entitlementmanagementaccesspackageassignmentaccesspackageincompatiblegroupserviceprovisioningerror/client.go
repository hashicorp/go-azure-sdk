package entitlementmanagementaccesspackageassignmentaccesspackageincompatiblegroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPackageIncompatibleGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPackageIncompatibleGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPackageIncompatibleGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspackageincompatiblegroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPackageIncompatibleGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPackageIncompatibleGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
