package conditionalaccessauthenticationstrengthpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessAuthenticationStrengthPolicyClient struct {
	Client *msgraph.Client
}

func NewConditionalAccessAuthenticationStrengthPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccessAuthenticationStrengthPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccessauthenticationstrengthpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccessAuthenticationStrengthPolicyClient: %+v", err)
	}

	return &ConditionalAccessAuthenticationStrengthPolicyClient{
		Client: client,
	}, nil
}
