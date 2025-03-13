package policytokens

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyTokensClient struct {
	Client *resourcemanager.Client
}

func NewPolicyTokensClientWithBaseURI(sdkApi sdkEnv.Api) (*PolicyTokensClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "policytokens", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyTokensClient: %+v", err)
	}

	return &PolicyTokensClient{
		Client: client,
	}, nil
}
