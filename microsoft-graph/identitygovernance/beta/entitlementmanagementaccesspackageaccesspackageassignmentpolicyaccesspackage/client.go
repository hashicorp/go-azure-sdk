package entitlementmanagementaccesspackageaccesspackageassignmentpolicyaccesspackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageaccesspackageassignmentpolicyaccesspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageClient{
		Client: client,
	}, nil
}
