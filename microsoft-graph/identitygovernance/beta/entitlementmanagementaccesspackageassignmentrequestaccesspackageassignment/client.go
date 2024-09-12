package entitlementmanagementaccesspackageassignmentrequestaccesspackageassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentRequestAccessPackageAssignmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentRequestAccessPackageAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentRequestAccessPackageAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentrequestaccesspackageassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentRequestAccessPackageAssignmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentRequestAccessPackageAssignmentClient{
		Client: client,
	}, nil
}
