package policyhomerealmdiscoverypolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyHomeRealmDiscoveryPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewPolicyHomeRealmDiscoveryPolicyAppliesToClientWithBaseURI(api sdkEnv.Api) (*PolicyHomeRealmDiscoveryPolicyAppliesToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyhomerealmdiscoverypolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyHomeRealmDiscoveryPolicyAppliesToClient: %+v", err)
	}

	return &PolicyHomeRealmDiscoveryPolicyAppliesToClient{
		Client: client,
	}, nil
}
