package conditionalaccesauthenticationstrengthpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccesAuthenticationStrengthPolicyClient struct {
	Client *msgraph.Client
}

func NewConditionalAccesAuthenticationStrengthPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccesAuthenticationStrengthPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccesauthenticationstrengthpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccesAuthenticationStrengthPolicyClient: %+v", err)
	}

	return &ConditionalAccesAuthenticationStrengthPolicyClient{
		Client: client,
	}, nil
}
