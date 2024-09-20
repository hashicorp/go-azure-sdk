package entitlementmanagementresourcerequestresourcescoperesourcerole

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestResourceScopeResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestResourceScopeResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestResourceScopeResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestresourcescoperesourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestResourceScopeResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestResourceScopeResourceRoleClient{
		Client: client,
	}, nil
}
