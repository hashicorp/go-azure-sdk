package identityprovider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProviderClient struct {
	Client *msgraph.Client
}

func NewIdentityProviderClientWithBaseURI(sdkApi sdkEnv.Api) (*IdentityProviderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "identityprovider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityProviderClient: %+v", err)
	}

	return &IdentityProviderClient{
		Client: client,
	}, nil
}
