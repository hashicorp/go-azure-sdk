package entitlementmanagementresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceNamespaceResourceActionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceNamespaceResourceActionClient: %+v", err)
	}

	return &EntitlementManagementResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
