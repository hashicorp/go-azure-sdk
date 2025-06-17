package authenticationqrcodepinmethodpin

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationQrCodePinMethodPinClient struct {
	Client *msgraph.Client
}

func NewAuthenticationQrCodePinMethodPinClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationQrCodePinMethodPinClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationqrcodepinmethodpin", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationQrCodePinMethodPinClient: %+v", err)
	}

	return &AuthenticationQrCodePinMethodPinClient{
		Client: client,
	}, nil
}
