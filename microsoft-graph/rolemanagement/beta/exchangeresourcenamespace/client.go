package exchangeresourcenamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeResourceNamespaceClient struct {
	Client *msgraph.Client
}

func NewExchangeResourceNamespaceClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeResourceNamespaceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeresourcenamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeResourceNamespaceClient: %+v", err)
	}

	return &ExchangeResourceNamespaceClient{
		Client: client,
	}, nil
}
