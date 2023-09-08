package applicationhomerealmdiscoverypolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationHomeRealmDiscoveryPolicyClient struct {
	Client *msgraph.Client
}

func NewApplicationHomeRealmDiscoveryPolicyClientWithBaseURI(api sdkEnv.Api) (*ApplicationHomeRealmDiscoveryPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationhomerealmdiscoverypolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationHomeRealmDiscoveryPolicyClient: %+v", err)
	}

	return &ApplicationHomeRealmDiscoveryPolicyClient{
		Client: client,
	}, nil
}
