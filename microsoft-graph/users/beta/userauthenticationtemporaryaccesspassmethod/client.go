package userauthenticationtemporaryaccesspassmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationTemporaryAccessPassMethodClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationTemporaryAccessPassMethodClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationTemporaryAccessPassMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationtemporaryaccesspassmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationTemporaryAccessPassMethodClient: %+v", err)
	}

	return &UserAuthenticationTemporaryAccessPassMethodClient{
		Client: client,
	}, nil
}
