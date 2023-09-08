package identityconditionalaccesauthenticationstrength

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityConditionalAccesAuthenticationStrengthClient struct {
	Client *msgraph.Client
}

func NewIdentityConditionalAccesAuthenticationStrengthClientWithBaseURI(api sdkEnv.Api) (*IdentityConditionalAccesAuthenticationStrengthClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityconditionalaccesauthenticationstrength", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityConditionalAccesAuthenticationStrengthClient: %+v", err)
	}

	return &IdentityConditionalAccesAuthenticationStrengthClient{
		Client: client,
	}, nil
}
