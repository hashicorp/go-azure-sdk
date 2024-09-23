package entitlementmanagementaccesspackageassignmentaccesspackageassignmentresourceroleaccesspackagesubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageSubjectClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspackageassignmentresourceroleaccesspackagesubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageSubjectClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageSubjectClient{
		Client: client,
	}, nil
}
