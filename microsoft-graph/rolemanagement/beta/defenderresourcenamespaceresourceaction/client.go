package defenderresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewDefenderResourceNamespaceResourceActionClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderResourceNamespaceResourceActionClient: %+v", err)
	}

	return &DefenderResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
