package identityconditionalaccesauthenticationstrengthauthenticationmethodmode

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient struct {
	Client *msgraph.Client
}

func NewIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClientWithBaseURI(api sdkEnv.Api) (*IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityconditionalaccesauthenticationstrengthauthenticationmethodmode", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient: %+v", err)
	}

	return &IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient{
		Client: client,
	}, nil
}
