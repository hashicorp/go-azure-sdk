package raitoollabels

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiToolLabelsClient struct {
	Client *resourcemanager.Client
}

func NewRaiToolLabelsClientWithBaseURI(sdkApi sdkEnv.Api) (*RaiToolLabelsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "raitoollabels", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RaiToolLabelsClient: %+v", err)
	}

	return &RaiToolLabelsClient{
		Client: client,
	}, nil
}
