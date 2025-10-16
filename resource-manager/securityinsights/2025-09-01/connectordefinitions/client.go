package connectordefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorDefinitionsClient struct {
	Client *resourcemanager.Client
}

func NewConnectorDefinitionsClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectorDefinitionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "connectordefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectorDefinitionsClient: %+v", err)
	}

	return &ConnectorDefinitionsClient{
		Client: client,
	}, nil
}
