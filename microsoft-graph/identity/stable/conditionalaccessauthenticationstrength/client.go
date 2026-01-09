package conditionalaccessauthenticationstrength

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessAuthenticationStrengthClient struct {
	Client *msgraph.Client
}

func NewConditionalAccessAuthenticationStrengthClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccessAuthenticationStrengthClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccessauthenticationstrength", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccessAuthenticationStrengthClient: %+v", err)
	}

	return &ConditionalAccessAuthenticationStrengthClient{
		Client: client,
	}, nil
}
