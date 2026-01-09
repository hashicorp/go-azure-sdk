package customauthenticationextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomAuthenticationExtensionClient struct {
	Client *msgraph.Client
}

func NewCustomAuthenticationExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomAuthenticationExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "customauthenticationextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomAuthenticationExtensionClient: %+v", err)
	}

	return &CustomAuthenticationExtensionClient{
		Client: client,
	}, nil
}
