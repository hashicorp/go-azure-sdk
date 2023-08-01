package extensionoperationstatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionOperationStatusClient struct {
	Client *resourcemanager.Client
}

func NewExtensionOperationStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*ExtensionOperationStatusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "extensionoperationstatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExtensionOperationStatusClient: %+v", err)
	}

	return &ExtensionOperationStatusClient{
		Client: client,
	}, nil
}
