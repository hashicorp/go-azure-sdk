package entitlementmanagementassignmentaccesspackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentAccessPackageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentAccessPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentAccessPackageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentaccesspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentAccessPackageClient: %+v", err)
	}

	return &EntitlementManagementAssignmentAccessPackageClient{
		Client: client,
	}, nil
}
