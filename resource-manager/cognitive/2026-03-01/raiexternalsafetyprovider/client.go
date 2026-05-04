package raiexternalsafetyprovider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiExternalSafetyProviderClient struct {
	Client *resourcemanager.Client
}

func NewRaiExternalSafetyProviderClientWithBaseURI(sdkApi sdkEnv.Api) (*RaiExternalSafetyProviderClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "raiexternalsafetyprovider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RaiExternalSafetyProviderClient: %+v", err)
	}

	return &RaiExternalSafetyProviderClient{
		Client: client,
	}, nil
}
