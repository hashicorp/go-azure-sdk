package enrichment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrichmentClient struct {
	Client *resourcemanager.Client
}

func NewEnrichmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EnrichmentClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "enrichment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnrichmentClient: %+v", err)
	}

	return &EnrichmentClient{
		Client: client,
	}, nil
}
