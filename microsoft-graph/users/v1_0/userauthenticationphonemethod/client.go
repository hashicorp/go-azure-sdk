package userauthenticationphonemethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationPhoneMethodClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationPhoneMethodClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationPhoneMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationphonemethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationPhoneMethodClient: %+v", err)
	}

	return &UserAuthenticationPhoneMethodClient{
		Client: client,
	}, nil
}
