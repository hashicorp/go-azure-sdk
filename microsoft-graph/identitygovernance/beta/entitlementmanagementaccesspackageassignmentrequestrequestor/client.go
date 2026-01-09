package entitlementmanagementaccesspackageassignmentrequestrequestor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentRequestRequestorClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentRequestRequestorClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentRequestRequestorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentrequestrequestor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentRequestRequestorClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentRequestRequestorClient{
		Client: client,
	}, nil
}
