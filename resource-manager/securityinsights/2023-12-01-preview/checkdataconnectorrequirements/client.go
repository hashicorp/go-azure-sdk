package checkdataconnectorrequirements

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckDataConnectorRequirementsClient struct {
	Client *resourcemanager.Client
}

func NewCheckDataConnectorRequirementsClientWithBaseURI(sdkApi sdkEnv.Api) (*CheckDataConnectorRequirementsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "checkdataconnectorrequirements", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CheckDataConnectorRequirementsClient: %+v", err)
	}

	return &CheckDataConnectorRequirementsClient{
		Client: client,
	}, nil
}
