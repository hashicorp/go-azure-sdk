package javacomponents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JavaComponentsClient struct {
	Client *resourcemanager.Client
}

func NewJavaComponentsClientWithBaseURI(sdkApi sdkEnv.Api) (*JavaComponentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "javacomponents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JavaComponentsClient: %+v", err)
	}

	return &JavaComponentsClient{
		Client: client,
	}, nil
}
