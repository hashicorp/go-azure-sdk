package hypervsites

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVSitesClient struct {
	Client *resourcemanager.Client
}

func NewHyperVSitesClientWithBaseURI(sdkApi sdkEnv.Api) (*HyperVSitesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hypervsites", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HyperVSitesClient: %+v", err)
	}

	return &HyperVSitesClient{
		Client: client,
	}, nil
}
