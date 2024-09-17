package keyvalues

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyValuesClient struct {
	Client *resourcemanager.Client
}

func NewKeyValuesClientWithBaseURI(sdkApi sdkEnv.Api) (*KeyValuesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "keyvalues", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KeyValuesClient: %+v", err)
	}

	return &KeyValuesClient{
		Client: client,
	}, nil
}
