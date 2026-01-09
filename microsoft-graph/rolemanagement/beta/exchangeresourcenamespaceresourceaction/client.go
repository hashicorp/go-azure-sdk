package exchangeresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewExchangeResourceNamespaceResourceActionClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeResourceNamespaceResourceActionClient: %+v", err)
	}

	return &ExchangeResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
