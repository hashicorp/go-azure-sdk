package entitlementmanagementaccesspackageassignmentaccesspackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPackageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPackageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPackageClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPackageClient{
		Client: client,
	}, nil
}
