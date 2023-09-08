package applicationtokenissuancepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTokenIssuancePolicyClient struct {
	Client *msgraph.Client
}

func NewApplicationTokenIssuancePolicyClientWithBaseURI(api sdkEnv.Api) (*ApplicationTokenIssuancePolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationtokenissuancepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationTokenIssuancePolicyClient: %+v", err)
	}

	return &ApplicationTokenIssuancePolicyClient{
		Client: client,
	}, nil
}
