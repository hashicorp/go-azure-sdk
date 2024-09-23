package exchangeroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewExchangeRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeRoleDefinitionClient: %+v", err)
	}

	return &ExchangeRoleDefinitionClient{
		Client: client,
	}, nil
}
