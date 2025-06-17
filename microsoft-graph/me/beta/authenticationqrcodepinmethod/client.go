package authenticationqrcodepinmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationQrCodePinMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationQrCodePinMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationQrCodePinMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationqrcodepinmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationQrCodePinMethodClient: %+v", err)
	}

	return &AuthenticationQrCodePinMethodClient{
		Client: client,
	}, nil
}
