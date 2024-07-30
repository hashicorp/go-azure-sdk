package entitlementmanagementaccesspackageincompatiblegroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageIncompatibleGroupClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageIncompatibleGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageIncompatibleGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageincompatiblegroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageIncompatibleGroupClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageIncompatibleGroupClient{
		Client: client,
	}, nil
}
