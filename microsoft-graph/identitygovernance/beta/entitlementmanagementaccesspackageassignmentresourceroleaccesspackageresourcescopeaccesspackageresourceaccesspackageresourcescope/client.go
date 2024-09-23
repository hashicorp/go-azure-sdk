package entitlementmanagementaccesspackageassignmentresourceroleaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentresourceroleaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient{
		Client: client,
	}, nil
}
