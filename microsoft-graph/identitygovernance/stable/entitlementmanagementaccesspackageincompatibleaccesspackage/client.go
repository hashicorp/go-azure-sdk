package entitlementmanagementaccesspackageincompatibleaccesspackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageIncompatibleAccessPackageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageIncompatibleAccessPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageIncompatibleAccessPackageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageincompatibleaccesspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageIncompatibleAccessPackageClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageIncompatibleAccessPackageClient{
		Client: client,
	}, nil
}
