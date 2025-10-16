package akriconnectortemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorTemplateClient struct {
	Client *resourcemanager.Client
}

func NewAkriConnectorTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*AkriConnectorTemplateClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "akriconnectortemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AkriConnectorTemplateClient: %+v", err)
	}

	return &AkriConnectorTemplateClient{
		Client: client,
	}, nil
}
