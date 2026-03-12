package identifiers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentifiersClient struct {
	Client *resourcemanager.Client
}

func NewIdentifiersClientWithBaseURI(sdkApi sdkEnv.Api) (*IdentifiersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "identifiers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentifiersClient: %+v", err)
	}

	return &IdentifiersClient{
		Client: client,
	}, nil
}
