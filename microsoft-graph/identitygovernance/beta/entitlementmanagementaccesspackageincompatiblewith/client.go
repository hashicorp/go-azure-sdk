package entitlementmanagementaccesspackageincompatiblewith

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageIncompatibleWithClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageIncompatibleWithClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageIncompatibleWithClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageincompatiblewith", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageIncompatibleWithClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageIncompatibleWithClient{
		Client: client,
	}, nil
}
