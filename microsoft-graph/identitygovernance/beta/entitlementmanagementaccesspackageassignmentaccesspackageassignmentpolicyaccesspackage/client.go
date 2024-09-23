package entitlementmanagementaccesspackageassignmentaccesspackageassignmentpolicyaccesspackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyAccessPackageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyAccessPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyAccessPackageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspackageassignmentpolicyaccesspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyAccessPackageClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentPolicyAccessPackageClient{
		Client: client,
	}, nil
}
