package authorizationpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationPolicyClient struct {
	Client *msgraph.Client
}

func NewAuthorizationPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthorizationPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authorizationpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthorizationPolicyClient: %+v", err)
	}

	return &AuthorizationPolicyClient{
		Client: client,
	}, nil
}
