package entitlementmanagementresourcerequestresourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestResourceRoleClient{
		Client: client,
	}, nil
}
