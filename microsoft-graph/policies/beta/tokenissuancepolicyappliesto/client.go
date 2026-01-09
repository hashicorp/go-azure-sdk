package tokenissuancepolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenIssuancePolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewTokenIssuancePolicyAppliesToClientWithBaseURI(sdkApi sdkEnv.Api) (*TokenIssuancePolicyAppliesToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "tokenissuancepolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TokenIssuancePolicyAppliesToClient: %+v", err)
	}

	return &TokenIssuancePolicyAppliesToClient{
		Client: client,
	}, nil
}
