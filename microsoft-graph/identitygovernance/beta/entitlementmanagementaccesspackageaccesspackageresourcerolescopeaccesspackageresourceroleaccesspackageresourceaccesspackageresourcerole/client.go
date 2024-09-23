package entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourceroleaccesspackageresourceaccesspackageresourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourceroleaccesspackageresourceaccesspackageresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleClient{
		Client: client,
	}, nil
}
