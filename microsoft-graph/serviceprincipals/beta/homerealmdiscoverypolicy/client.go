package homerealmdiscoverypolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HomeRealmDiscoveryPolicyClient struct {
	Client *msgraph.Client
}

func NewHomeRealmDiscoveryPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*HomeRealmDiscoveryPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "homerealmdiscoverypolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HomeRealmDiscoveryPolicyClient: %+v", err)
	}

	return &HomeRealmDiscoveryPolicyClient{
		Client: client,
	}, nil
}
