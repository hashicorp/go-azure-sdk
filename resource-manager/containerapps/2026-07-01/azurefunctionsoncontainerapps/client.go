package azurefunctionsoncontainerapps

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureFunctionsOnContainerAppsClient struct {
	Client *resourcemanager.Client
}

func NewAzureFunctionsOnContainerAppsClientWithBaseURI(sdkApi sdkEnv.Api) (*AzureFunctionsOnContainerAppsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "azurefunctionsoncontainerapps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AzureFunctionsOnContainerAppsClient: %+v", err)
	}

	return &AzureFunctionsOnContainerAppsClient{
		Client: client,
	}, nil
}
