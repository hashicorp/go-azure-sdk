package asev3networkingconfigurations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AseV3NetworkingConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewAseV3NetworkingConfigurationsClientWithBaseURI(sdkApi sdkEnv.Api) (*AseV3NetworkingConfigurationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "asev3networkingconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AseV3NetworkingConfigurationsClient: %+v", err)
	}

	return &AseV3NetworkingConfigurationsClient{
		Client: client,
	}, nil
}
