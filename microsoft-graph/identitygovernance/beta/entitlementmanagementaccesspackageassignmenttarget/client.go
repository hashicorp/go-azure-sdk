package entitlementmanagementaccesspackageassignmenttarget

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentTargetClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentTargetClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentTargetClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmenttarget", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentTargetClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentTargetClient{
		Client: client,
	}, nil
}
