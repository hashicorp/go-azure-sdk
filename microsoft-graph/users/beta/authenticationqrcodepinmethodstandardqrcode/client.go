package authenticationqrcodepinmethodstandardqrcode

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationQrCodePinMethodStandardQRCodeClient struct {
	Client *msgraph.Client
}

func NewAuthenticationQrCodePinMethodStandardQRCodeClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationQrCodePinMethodStandardQRCodeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationqrcodepinmethodstandardqrcode", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationQrCodePinMethodStandardQRCodeClient: %+v", err)
	}

	return &AuthenticationQrCodePinMethodStandardQRCodeClient{
		Client: client,
	}, nil
}
