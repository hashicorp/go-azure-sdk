package defenderresourcenamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderResourceNamespaceClient struct {
	Client *msgraph.Client
}

func NewDefenderResourceNamespaceClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderResourceNamespaceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderresourcenamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderResourceNamespaceClient: %+v", err)
	}

	return &DefenderResourceNamespaceClient{
		Client: client,
	}, nil
}
