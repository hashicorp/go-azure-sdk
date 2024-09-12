package entitlementmanagementaccesspackageresourcerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceRequestClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourcerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceRequestClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceRequestClient{
		Client: client,
	}, nil
}
