package entitlementmanagementaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient{
		Client: client,
	}, nil
}
