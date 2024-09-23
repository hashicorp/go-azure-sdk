package exchange

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeClient struct {
	Client *msgraph.Client
}

func NewExchangeClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchange", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeClient: %+v", err)
	}

	return &ExchangeClient{
		Client: client,
	}, nil
}
