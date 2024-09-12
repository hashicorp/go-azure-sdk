package entitlementmanagementaccesspackageassignmentresourceroleresourceaccesspackageresourcescoperesourceresourceenvironment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentresourceroleresourceaccesspackageresourcescoperesourceresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceResourceEnvironmentClient{
		Client: client,
	}, nil
}
