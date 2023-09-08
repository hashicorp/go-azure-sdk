package serviceprincipalhomerealmdiscoverypolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalHomeRealmDiscoveryPolicyClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalHomeRealmDiscoveryPolicyClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalHomeRealmDiscoveryPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalhomerealmdiscoverypolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalHomeRealmDiscoveryPolicyClient: %+v", err)
	}

	return &ServicePrincipalHomeRealmDiscoveryPolicyClient{
		Client: client,
	}, nil
}
