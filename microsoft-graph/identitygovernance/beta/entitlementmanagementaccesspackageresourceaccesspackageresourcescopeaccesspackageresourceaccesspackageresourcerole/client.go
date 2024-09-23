package entitlementmanagementaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient{
		Client: client,
	}, nil
}
