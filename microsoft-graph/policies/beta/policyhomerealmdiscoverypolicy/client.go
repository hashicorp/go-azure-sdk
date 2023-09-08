package policyhomerealmdiscoverypolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyHomeRealmDiscoveryPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyHomeRealmDiscoveryPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyHomeRealmDiscoveryPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyhomerealmdiscoverypolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyHomeRealmDiscoveryPolicyClient: %+v", err)
	}

	return &PolicyHomeRealmDiscoveryPolicyClient{
		Client: client,
	}, nil
}
