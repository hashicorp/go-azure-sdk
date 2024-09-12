package tokenlifetimepolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenLifetimePolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewTokenLifetimePolicyAppliesToClientWithBaseURI(sdkApi sdkEnv.Api) (*TokenLifetimePolicyAppliesToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "tokenlifetimepolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TokenLifetimePolicyAppliesToClient: %+v", err)
	}

	return &TokenLifetimePolicyAppliesToClient{
		Client: client,
	}, nil
}
