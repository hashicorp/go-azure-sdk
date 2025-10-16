package integrationfabrics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationFabricsClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationFabricsClientWithBaseURI(sdkApi sdkEnv.Api) (*IntegrationFabricsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "integrationfabrics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationFabricsClient: %+v", err)
	}

	return &IntegrationFabricsClient{
		Client: client,
	}, nil
}
