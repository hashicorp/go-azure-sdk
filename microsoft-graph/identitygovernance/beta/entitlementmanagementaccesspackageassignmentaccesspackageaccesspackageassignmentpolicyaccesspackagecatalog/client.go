package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageassignmentpolicyaccesspackagecatalog

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageassignmentpolicyaccesspackagecatalog", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClient{
		Client: client,
	}, nil
}
