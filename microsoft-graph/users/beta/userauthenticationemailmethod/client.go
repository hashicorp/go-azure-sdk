package userauthenticationemailmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationEmailMethodClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationEmailMethodClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationEmailMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationemailmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationEmailMethodClient: %+v", err)
	}

	return &UserAuthenticationEmailMethodClient{
		Client: client,
	}, nil
}
