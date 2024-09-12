package entitlementmanagementaccesspackageaccesspackageassignmentpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackageAssignmentPolicyClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackageAssignmentPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackageAssignmentPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackageassignmentpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackageAssignmentPolicyClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackageAssignmentPolicyClient{
		Client: client,
	}, nil
}
