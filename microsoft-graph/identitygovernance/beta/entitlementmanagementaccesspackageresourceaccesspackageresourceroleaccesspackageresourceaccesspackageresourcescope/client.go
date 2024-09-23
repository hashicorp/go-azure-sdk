package entitlementmanagementaccesspackageresourceaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourceaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient{
		Client: client,
	}, nil
}
