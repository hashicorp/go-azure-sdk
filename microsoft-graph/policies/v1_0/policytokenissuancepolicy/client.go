package policytokenissuancepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyTokenIssuancePolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyTokenIssuancePolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyTokenIssuancePolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policytokenissuancepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyTokenIssuancePolicyClient: %+v", err)
	}

	return &PolicyTokenIssuancePolicyClient{
		Client: client,
	}, nil
}
