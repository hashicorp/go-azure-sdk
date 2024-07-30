package conditionalaccesauthenticationstrength

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccesAuthenticationStrengthClient struct {
	Client *msgraph.Client
}

func NewConditionalAccesAuthenticationStrengthClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccesAuthenticationStrengthClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccesauthenticationstrength", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccesAuthenticationStrengthClient: %+v", err)
	}

	return &ConditionalAccesAuthenticationStrengthClient{
		Client: client,
	}, nil
}
