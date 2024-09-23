package entitlementmanagementaccesspackageassignmentresourceroleaccesspackageresourceroleaccesspackageresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceRoleAccessPackageResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceRoleAccessPackageResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceRoleAccessPackageResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentresourceroleaccesspackageresourceroleaccesspackageresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceRoleAccessPackageResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceRoleAccessPackageResourceClient{
		Client: client,
	}, nil
}
