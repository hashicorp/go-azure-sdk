package authenticationplatformcredentialmethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationPlatformCredentialMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewAuthenticationPlatformCredentialMethodDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationPlatformCredentialMethodDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationplatformcredentialmethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationPlatformCredentialMethodDeviceClient: %+v", err)
	}

	return &AuthenticationPlatformCredentialMethodDeviceClient{
		Client: client,
	}, nil
}
