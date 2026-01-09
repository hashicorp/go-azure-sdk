package entitlementmanagementaccesspackageresourceaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescopeaccesspackageresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourceaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescopeaccesspackageresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceClient{
		Client: client,
	}, nil
}
