package cognitiveservicesaccounts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CognitiveServicesAccountsClient struct {
	Client *resourcemanager.Client
}

func NewCognitiveServicesAccountsClientWithBaseURI(api sdkEnv.Api) (*CognitiveServicesAccountsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "cognitiveservicesaccounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CognitiveServicesAccountsClient: %+v", err)
	}

	return &CognitiveServicesAccountsClient{
		Client: client,
	}, nil
}
