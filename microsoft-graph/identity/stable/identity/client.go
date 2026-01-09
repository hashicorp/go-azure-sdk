package identity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityClient struct {
	Client *msgraph.Client
}

func NewIdentityClientWithBaseURI(sdkApi sdkEnv.Api) (*IdentityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "identity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityClient: %+v", err)
	}

	return &IdentityClient{
		Client: client,
	}, nil
}
