package conditionalaccessauthenticationstrengthpolicycombinationconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient struct {
	Client *msgraph.Client
}

func NewConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccessauthenticationstrengthpolicycombinationconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient: %+v", err)
	}

	return &ConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient{
		Client: client,
	}, nil
}
