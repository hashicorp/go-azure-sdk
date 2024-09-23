package entitlementmanagementaccesspackageresourceaccesspackageresourceroleaccesspackageresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourceaccesspackageresourceroleaccesspackageresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient{
		Client: client,
	}, nil
}
