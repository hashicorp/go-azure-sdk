package conditionalaccesauthenticationstrengthpolicycombinationconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient struct {
	Client *msgraph.Client
}

func NewConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccesauthenticationstrengthpolicycombinationconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient: %+v", err)
	}

	return &ConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationClient{
		Client: client,
	}, nil
}
