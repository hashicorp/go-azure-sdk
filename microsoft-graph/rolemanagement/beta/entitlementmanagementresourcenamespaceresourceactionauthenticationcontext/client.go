package entitlementmanagementresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &EntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
