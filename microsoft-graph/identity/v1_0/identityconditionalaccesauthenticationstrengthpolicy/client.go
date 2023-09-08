package identityconditionalaccesauthenticationstrengthpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityConditionalAccesAuthenticationStrengthPolicyClient struct {
	Client *msgraph.Client
}

func NewIdentityConditionalAccesAuthenticationStrengthPolicyClientWithBaseURI(api sdkEnv.Api) (*IdentityConditionalAccesAuthenticationStrengthPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityconditionalaccesauthenticationstrengthpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityConditionalAccesAuthenticationStrengthPolicyClient: %+v", err)
	}

	return &IdentityConditionalAccesAuthenticationStrengthPolicyClient{
		Client: client,
	}, nil
}
