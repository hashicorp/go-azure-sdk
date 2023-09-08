package userauthenticationfido2method

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationFido2MethodClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationFido2MethodClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationFido2MethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationfido2method", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationFido2MethodClient: %+v", err)
	}

	return &UserAuthenticationFido2MethodClient{
		Client: client,
	}, nil
}
