package entitlementmanagementaccesspackageassignmentresourceroleaccesspackagesubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageSubjectClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentresourceroleaccesspackagesubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageSubjectClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageSubjectClient{
		Client: client,
	}, nil
}
