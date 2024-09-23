package entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequestrequestor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestRequestorClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestRequestorClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestRequestorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequestrequestor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestRequestorClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestRequestorClient{
		Client: client,
	}, nil
}
