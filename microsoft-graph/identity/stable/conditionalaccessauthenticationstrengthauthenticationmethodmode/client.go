package conditionalaccessauthenticationstrengthauthenticationmethodmode

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient struct {
	Client *msgraph.Client
}

func NewConditionalAccessAuthenticationStrengthAuthenticationMethodModeClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccessauthenticationstrengthauthenticationmethodmode", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient: %+v", err)
	}

	return &ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient{
		Client: client,
	}, nil
}
