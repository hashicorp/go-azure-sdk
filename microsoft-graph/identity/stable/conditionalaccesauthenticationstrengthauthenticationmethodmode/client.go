package conditionalaccesauthenticationstrengthauthenticationmethodmode

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient struct {
	Client *msgraph.Client
}

func NewConditionalAccesAuthenticationStrengthAuthenticationMethodModeClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccesauthenticationstrengthauthenticationmethodmode", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient: %+v", err)
	}

	return &ConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient{
		Client: client,
	}, nil
}
