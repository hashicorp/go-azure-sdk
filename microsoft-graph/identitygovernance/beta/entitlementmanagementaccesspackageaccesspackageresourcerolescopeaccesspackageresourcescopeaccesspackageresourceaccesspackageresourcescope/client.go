package entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient{
		Client: client,
	}, nil
}
