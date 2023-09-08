package userauthenticationmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationMethodClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationMethodClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationMethodClient: %+v", err)
	}

	return &UserAuthenticationMethodClient{
		Client: client,
	}, nil
}
