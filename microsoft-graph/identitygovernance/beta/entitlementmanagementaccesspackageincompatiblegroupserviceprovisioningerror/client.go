package entitlementmanagementaccesspackageincompatiblegroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageincompatiblegroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
