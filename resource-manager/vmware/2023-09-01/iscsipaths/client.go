package iscsipaths

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IscsiPathsClient struct {
	Client *resourcemanager.Client
}

func NewIscsiPathsClientWithBaseURI(sdkApi sdkEnv.Api) (*IscsiPathsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "iscsipaths", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IscsiPathsClient: %+v", err)
	}

	return &IscsiPathsClient{
		Client: client,
	}, nil
}
