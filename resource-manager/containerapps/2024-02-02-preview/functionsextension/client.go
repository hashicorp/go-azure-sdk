package functionsextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FunctionsExtensionClient struct {
	Client *resourcemanager.Client
}

func NewFunctionsExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*FunctionsExtensionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "functionsextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FunctionsExtensionClient: %+v", err)
	}

	return &FunctionsExtensionClient{
		Client: client,
	}, nil
}
