package entitlementmanagementaccesspackageaccesspackageassignmentpolicyaccesspackagecatalog

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackageassignmentpolicyaccesspackagecatalog", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClient{
		Client: client,
	}, nil
}
