package datatypes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataTypesClient struct {
	Client *resourcemanager.Client
}

func NewDataTypesClientWithBaseURI(sdkApi sdkEnv.Api) (*DataTypesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "datatypes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataTypesClient: %+v", err)
	}

	return &DataTypesClient{
		Client: client,
	}, nil
}
