package defenderresourcenamespaceresourceactionresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderResourceNamespaceResourceActionResourceScopeClient struct {
	Client *msgraph.Client
}

func NewDefenderResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderResourceNamespaceResourceActionResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderresourcenamespaceresourceactionresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderResourceNamespaceResourceActionResourceScopeClient: %+v", err)
	}

	return &DefenderResourceNamespaceResourceActionResourceScopeClient{
		Client: client,
	}, nil
}
