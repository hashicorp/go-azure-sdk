package authentications

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationsClient struct {
	Client *resourcemanager.Client
}

func NewAuthenticationsClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "authentications", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationsClient: %+v", err)
	}

	return &AuthenticationsClient{
		Client: client,
	}, nil
}
