package exchangecustomappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeCustomAppScopeClient struct {
	Client *msgraph.Client
}

func NewExchangeCustomAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeCustomAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangecustomappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeCustomAppScopeClient: %+v", err)
	}

	return &ExchangeCustomAppScopeClient{
		Client: client,
	}, nil
}
