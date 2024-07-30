package tokenlifetimepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenLifetimePolicyClient struct {
	Client *msgraph.Client
}

func NewTokenLifetimePolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*TokenLifetimePolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "tokenlifetimepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TokenLifetimePolicyClient: %+v", err)
	}

	return &TokenLifetimePolicyClient{
		Client: client,
	}, nil
}
