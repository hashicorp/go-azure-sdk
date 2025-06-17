package authenticationrequirement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationRequirementClient struct {
	Client *msgraph.Client
}

func NewAuthenticationRequirementClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationRequirementClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationrequirement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationRequirementClient: %+v", err)
	}

	return &AuthenticationRequirementClient{
		Client: client,
	}, nil
}
