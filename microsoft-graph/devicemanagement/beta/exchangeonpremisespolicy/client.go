package exchangeonpremisespolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeOnPremisesPolicyClient struct {
	Client *msgraph.Client
}

func NewExchangeOnPremisesPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeOnPremisesPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeonpremisespolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeOnPremisesPolicyClient: %+v", err)
	}

	return &ExchangeOnPremisesPolicyClient{
		Client: client,
	}, nil
}
