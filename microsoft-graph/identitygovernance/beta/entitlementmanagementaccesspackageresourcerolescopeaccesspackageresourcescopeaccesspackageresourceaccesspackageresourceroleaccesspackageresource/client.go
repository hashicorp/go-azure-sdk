package entitlementmanagementaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient{
		Client: client,
	}, nil
}
