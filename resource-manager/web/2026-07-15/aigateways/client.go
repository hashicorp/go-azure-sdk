package aigateways

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AiGatewaysClient struct {
	Client *resourcemanager.Client
}

func NewAiGatewaysClientWithBaseURI(sdkApi sdkEnv.Api) (*AiGatewaysClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "aigateways", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AiGatewaysClient: %+v", err)
	}

	return &AiGatewaysClient{
		Client: client,
	}, nil
}
