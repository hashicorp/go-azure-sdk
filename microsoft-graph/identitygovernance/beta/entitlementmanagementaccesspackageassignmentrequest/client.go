package entitlementmanagementaccesspackageassignmentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentRequestClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentRequestClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentRequestClient{
		Client: client,
	}, nil
}
