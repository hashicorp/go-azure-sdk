package federatedtokenvalidationpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FederatedTokenValidationPolicyClient struct {
	Client *msgraph.Client
}

func NewFederatedTokenValidationPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*FederatedTokenValidationPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "federatedtokenvalidationpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FederatedTokenValidationPolicyClient: %+v", err)
	}

	return &FederatedTokenValidationPolicyClient{
		Client: client,
	}, nil
}
