package exchangeconnector

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeConnectorClient struct {
	Client *msgraph.Client
}

func NewExchangeConnectorClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeConnectorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeconnector", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeConnectorClient: %+v", err)
	}

	return &ExchangeConnectorClient{
		Client: client,
	}, nil
}
