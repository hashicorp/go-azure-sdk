package confluents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfluentsClient struct {
	Client *resourcemanager.Client
}

func NewConfluentsClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfluentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "confluents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfluentsClient: %+v", err)
	}

	return &ConfluentsClient{
		Client: client,
	}, nil
}
