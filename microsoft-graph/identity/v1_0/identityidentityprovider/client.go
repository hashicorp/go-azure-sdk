package identityidentityprovider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityIdentityProviderClient struct {
	Client *msgraph.Client
}

func NewIdentityIdentityProviderClientWithBaseURI(api sdkEnv.Api) (*IdentityIdentityProviderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityidentityprovider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityIdentityProviderClient: %+v", err)
	}

	return &IdentityIdentityProviderClient{
		Client: client,
	}, nil
}
