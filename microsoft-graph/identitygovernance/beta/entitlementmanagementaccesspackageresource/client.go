package entitlementmanagementaccesspackageresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceClient{
		Client: client,
	}, nil
}
