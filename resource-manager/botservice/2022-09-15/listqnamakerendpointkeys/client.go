package listqnamakerendpointkeys

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListQnAMakerEndpointKeysClient struct {
	Client *resourcemanager.Client
}

func NewListQnAMakerEndpointKeysClientWithBaseURI(sdkApi sdkEnv.Api) (*ListQnAMakerEndpointKeysClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "listqnamakerendpointkeys", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ListQnAMakerEndpointKeysClient: %+v", err)
	}

	return &ListQnAMakerEndpointKeysClient{
		Client: client,
	}, nil
}
