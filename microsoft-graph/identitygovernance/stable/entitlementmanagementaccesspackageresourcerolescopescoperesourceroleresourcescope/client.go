package entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClient{
		Client: client,
	}, nil
}
