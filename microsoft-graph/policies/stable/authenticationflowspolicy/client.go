package authenticationflowspolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationFlowsPolicyClient struct {
	Client *msgraph.Client
}

func NewAuthenticationFlowsPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationFlowsPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationflowspolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationFlowsPolicyClient: %+v", err)
	}

	return &AuthenticationFlowsPolicyClient{
		Client: client,
	}, nil
}
