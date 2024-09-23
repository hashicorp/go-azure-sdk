package homerealmdiscoverypolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HomeRealmDiscoveryPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewHomeRealmDiscoveryPolicyAppliesToClientWithBaseURI(sdkApi sdkEnv.Api) (*HomeRealmDiscoveryPolicyAppliesToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "homerealmdiscoverypolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HomeRealmDiscoveryPolicyAppliesToClient: %+v", err)
	}

	return &HomeRealmDiscoveryPolicyAppliesToClient{
		Client: client,
	}, nil
}
