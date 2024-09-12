package authentication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationClient struct {
	Client *msgraph.Client
}

func NewAuthenticationClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authentication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationClient: %+v", err)
	}

	return &AuthenticationClient{
		Client: client,
	}, nil
}
