package entitlementmanagementassignmentrequestaccesspackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentRequestAccessPackageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentRequestAccessPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentRequestAccessPackageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentrequestaccesspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentRequestAccessPackageClient: %+v", err)
	}

	return &EntitlementManagementAssignmentRequestAccessPackageClient{
		Client: client,
	}, nil
}
