package entitlementmanagementaccesspackageaccesspackagesincompatiblewith

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackagesIncompatibleWithClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackagesincompatiblewith", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient{
		Client: client,
	}, nil
}
